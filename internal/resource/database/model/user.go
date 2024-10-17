package model

import (
	"image-retrieval/internal/resource/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string  `gorm:"type:varchar(20);not null;unique"`
	Password      string  `gorm:"type:varchar(12);not null;"`
	Email         string  `gorm:"type:varchar(50);not null;unique"`
	ImagesLove    []Image `gorm:"many2many:user_images;"`
	ImagesHistory []Image `gorm:"many2many:history_images;"`
}

func (*User) TableName() string {
	return "users"
}

func GetUser(username string) (user *User, err error) {
	user = &User{}
	res := database.DB.Model(&User{}).Where("username = ?", username).Find(user)
	if res.Error != nil {
		err = res.Error
		return
	}
	if res.RowsAffected == 0 {
		user = nil
	}
	return
}
func CreateUser(user *User) (err error) {
	err = database.DB.Model(&User{}).Create(user).Error
	return
}
