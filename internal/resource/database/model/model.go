package model

import "image-retrieval/internal/resource/database"

func InitModel() {
	database.DB.AutoMigrate(&User{})
	database.DB.AutoMigrate(&Image{})
}
