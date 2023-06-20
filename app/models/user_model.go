package models

type UserModel struct {
	BaseModel
	UserId    string `gorm:"primaryKey"`
	UserName  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Level     string `gorm:"not null"`
	IsDeleted bool   `gorm:"default:false"`
}

func (UserModel) TableName() string {
	return "tm_user"
}
