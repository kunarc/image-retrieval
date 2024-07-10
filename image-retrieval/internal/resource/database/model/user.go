package model

import (
	"image-retrieval/internal/resource/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
}

func (*User) TableName() string {
	return "users"
}

func GetUser() (user *User, err error) {
	err = database.DB.Model(&User{}).Find(user).Error
	return
}
