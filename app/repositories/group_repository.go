package repositories

import (
	"errors"
	"fmt"
	"stockingapi/app/models"

	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func CreateGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (r *GroupRepository) FindAll() (*[]models.GroupModel, error) {
	var groupModels []models.GroupModel

	result := r.db.Find(&groupModels)
	if result.Error != nil {
		fmt.Println("ERR: (GroupRepository)(FindAll) ", result.Error)
		return nil, result.Error
	}

	return &groupModels, nil
}

func (r *GroupRepository) FindOneByCode(groupCode string) (*models.GroupModel, error) {
	var groupModel models.GroupModel

	result := r.db.Where("group_code = ?", groupCode).First(&groupModel)
	if result.Error != nil {
		fmt.Println("ERR: (GroupRepository)(FindOneByCode) ", result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &groupModel, nil
}

func (r *GroupRepository) Create(groupModel *models.GroupModel) *gorm.DB {
	return r.db.Create(groupModel)
}

func (r *GroupRepository) Update(groupCode string, updateModel *models.EditGroupModel) error {
	var groupModel models.GroupModel

	result := r.db.Where("group_code = ?", groupCode).First(&groupModel)
	if result.Error != nil {
		return result.Error
	}

	groupModel.GroupName = updateModel.GroupName

	r.db.Save(&groupModel)
	return nil
}

func (r *GroupRepository) Delete(groupCode string) error {
	result := r.db.Where("group_code = ?", groupCode).Delete(&models.GroupModel{})
	if result.Error != nil {
		fmt.Println("ERR: (GroupRepository)(Delete) ", result.Error.Error())
		return result.Error
	}

	return nil
}
