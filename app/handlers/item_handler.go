package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"stockingapi/app/dtos"
	"stockingapi/app/models"
	"stockingapi/app/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ItemHandler struct {
	validate            *validator.Validate
	db                  *gorm.DB
	groupRepository     *repositories.GroupRepository
	warehouseRepository *repositories.WarehouseRepository
	shelfRepository     *repositories.ShelfRepository
	supplierRepository  *repositories.SupplierRepository
	itemRepository      *repositories.ItemRepository
}

func CreateItemHandler(db *gorm.DB) *ItemHandler {
	return &ItemHandler{
		validate:            validator.New(),
		db:                  db,
		groupRepository:     repositories.CreateGroupRepository(db),
		warehouseRepository: repositories.CreateWarehouseRepository(db),
		shelfRepository:     repositories.CreateShelfRepository(db),
		supplierRepository:  repositories.CreateSupplierRepository(db),
		itemRepository:      repositories.CreateItemRepository(db),
	}
}

type ValidateItemReturn struct {
	groupModel    *models.GroupModel
	shelfModel    *models.ShelfModel
	supplierModel *models.SupplierModel
}

func (h *ItemHandler) validateItemRelation(reqBody *dtos.CreateItemRequestDTO) (*ValidateItemReturn, error) {
	groupExists, _ := h.groupRepository.FindOneByCode(reqBody.GroupCode)
	if groupExists == nil {
		return nil, errors.New("Group not found.")
	}

	shelfExists, _ := h.shelfRepository.FindOneShelfByCode(reqBody.ShelfCode)
	if shelfExists == nil {
		return nil, errors.New("Shelf not found.")
	}

	supplierExists, _ := h.supplierRepository.FindOneSupplierByCode(reqBody.SupplierCode)
	if supplierExists == nil {
		return nil, errors.New("Supplier not found.")
	}

	return &ValidateItemReturn{
		groupExists,
		shelfExists,
		supplierExists,
	}, nil
}

func (h *ItemHandler) Create(c *gin.Context) {
	var reqBody dtos.CreateItemRequestDTO
	userProps, _ := c.Get("user")
	user, _ := userProps.(string)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validateResult, errValidate := h.validateItemRelation(&reqBody)
	if errValidate != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": errValidate.Error()})
		return
	}

	error := h.db.Transaction(func(tx *gorm.DB) error {
		itemRepository := repositories.CreateItemRepository(tx)
		itemEventRepository := repositories.CreateItemEventRepository(tx)

		// TODO: (ITEM TRANSACTION SCRIPT TMR. - SUN, 4 JUN 2023 23:52)
		itemExists, _ := itemRepository.FindOneByCode(reqBody.ItemCode)
		if itemExists != nil {
			return errors.New("Item code already taken.")
		}

		itemModel := models.ItemModel{
			ItemCode:  reqBody.ItemCode,
			ItemName:  reqBody.ItemName,
			ItemInfo:  reqBody.ItemInfo,
			Size:      reqBody.Size,
			Stock:     reqBody.Stock,
			Weight:    reqBody.Weight,
			Price:     reqBody.Price,
			CreatedBy: user,
			Group:     *validateResult.groupModel,
			Shelf:     *validateResult.shelfModel,
			Supplier:  *validateResult.supplierModel,
		}

		if err := itemRepository.Create(&itemModel).Error; err != nil {
			return errors.New(err.Error())
		}

		itemEventModel := models.ItemEventModel{
			StockBefore: 0,
			StockChange: int(itemModel.Stock),
			StockAfter:  itemModel.Stock,
			Type:        "CREATE",
			CreatedBy:   user,
			Item:        itemModel,
			Shelf:       *validateResult.shelfModel,
		}

		if err := itemEventRepository.Create(&itemEventModel).Error; err != nil {
			return errors.New(err.Error())
		}
		return nil
	})

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item successfully created."})
}

func (h *ItemHandler) Find(c *gin.Context) {
	itemModels, err := h.itemRepository.FindBy()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, itemModels)
	return
}

func (h *ItemHandler) Transaction(c *gin.Context) {
	var reqBody dtos.TransactionItemRequestDTO
	var user = c.MustGet("user").(string)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if isValidType := reqBody.ValidateType(); isValidType == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Type is not valid, please refer to a valid type."})
		return
	}

	error := h.db.Transaction(func(tx *gorm.DB) error {
		itemRepository := repositories.CreateItemRepository(tx)
		itemEventRepository := repositories.CreateItemEventRepository(tx)

		itemExist, _ := itemRepository.FindOneByCode(reqBody.ItemCode)

		if itemExist == nil {
			return errors.New("Item not found.")
		}

		stockUpdate := int(reqBody.Stock)

		if stockUpdate > int(itemExist.Stock) {
			return errors.New("Stock cannot more than Item Stock.")
		}

		if reqBody.Type == "OUT" {
			stockUpdate = int(reqBody.Stock) * -1
		}

		newStock := int(itemExist.Stock) + stockUpdate
		if err := itemRepository.UpdateStock(reqBody.ItemCode, uint(newStock)); err != nil {
			return errors.New(err.Error())
		}

		itemEvent := models.ItemEventModel{
			Type:        "TRANSACTION",
			StockBefore: itemExist.Stock,
			StockChange: stockUpdate,
			StockAfter:  uint(newStock),
			CreatedBy:   user,
			RShelfCode:  itemExist.RShelfCode,
			Item:        *itemExist,
		}

		if err := itemEventRepository.Create(&itemEvent).Error; err != nil {
			return errors.New(err.Error())
		}

		return nil
	})
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successfully created."})
}

func (h *ItemHandler) Movement(c *gin.Context) {
	var reqBody dtos.MovementItemRequestDTO
	var user = c.MustGet("user").(string)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	itemExist, _ := h.itemRepository.FindOneByCode(reqBody.ItemCode)
	if itemExist == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Item code not found."})
		return
	}
	fmt.Println(itemExist.Shelf)

	shelfBefore, ShelfAfter, errShelf := h.validateMovementShelf(reqBody.ShelfBefore, reqBody.ShelfAfter)
	if errShelf != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": errShelf.Error()})
		return
	}

	errTrx := h.db.Transaction(func(tx *gorm.DB) error {
		itemRepository := repositories.CreateItemRepository(tx)
		itemEventRepository := repositories.CreateItemEventRepository(tx)

		itemExist, _ := itemRepository.FindOneByCode(reqBody.ItemCode)
		if itemExist == nil {
			return errors.New("Item Code not found.")
		}

		if itemExist.RShelfCode != reqBody.ShelfBefore {
			return errors.New("Item Shelf not found.")
		}

		if err := itemRepository.UpdateShelf(reqBody.ItemCode, reqBody.ShelfAfter); err != nil {
			return err
		}
		fmt.Println(itemExist.Shelf)

		itemEventModelOut := models.ItemEventModel{
			Item:        *itemExist,
			Shelf:       *shelfBefore,
			Type:        "MOVEMENT",
			StockBefore: itemExist.Stock,
			StockChange: int(itemExist.Stock) * -1,
			StockAfter:  0,
			CreatedBy:   user,
		}

		errItemOut := itemEventRepository.Create(&itemEventModelOut).Error
		if errItemOut != nil {
			return errors.New("ERR_ITEM_OUT: " + errItemOut.Error())
		}

		itemEventModelIn := models.ItemEventModel{
			Item:        *itemExist,
			Shelf:       *ShelfAfter,
			Type:        "MOVEMENT",
			StockBefore: 0,
			StockChange: int(itemExist.Stock),
			StockAfter:  itemExist.Stock,
			CreatedBy:   user,
		}

		errItemIn := itemEventRepository.Create(&itemEventModelIn).Error
		if errItemIn != nil {
			return errors.New("ERR_ITEM_IN: " + errItemIn.Error())
		}

		return nil
	})

	if errTrx != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errTrx.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item successfully moved."})
	return
}

func (h *ItemHandler) validateMovementShelf(shelf_code_before string, shelf_code_after string) (*models.ShelfModel, *models.ShelfModel, error) {
	shelfCodeBefore, _ := h.shelfRepository.FindOneShelfByCode(shelf_code_before)
	if shelfCodeBefore == nil {
		return nil, nil, errors.New("Origin Shelf Code is not found.")
	}

	shelfCodeAfter, _ := h.shelfRepository.FindOneShelfByCode(shelf_code_after)
	if shelfCodeAfter == nil {
		return nil, nil, errors.New("Target Shelf Code is not found.")
	}

	return shelfCodeBefore, shelfCodeAfter, nil
}
