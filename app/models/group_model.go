package models

type GroupModel struct {
	BaseModel
	GroupCode string `gorm:"primaryKey"`
	GroupName string `gorm:"not null"`
}

func (GroupModel) TableName() string {
	return "tm_group"
}

type EditGroupModel struct {
	GroupName string
}
