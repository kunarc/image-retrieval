package model

import (
	"image-retrieval/internal/resource/database"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Name         string `gorm:"type:varchar(30);not null;" json:"name"`
	Artist       string `gorm:"type:varchar(30);" json:"artist"`
	Picture      string `gorm:"type:varchar(255);not null;" json:"picture"`
	Year         string `gorm:"type:varchar(10);" json:"year"`
	Type         string `gorm:"type:varchar(10);not null;" json:"type"`
	SpecificType string `gorm:"type:varchar(20);not null;" json:"specificType"`
	Size         string `gorm:"type:varchar(20);" json:"size"`
	Museum       string `gorm:"type:varchar(50);" json:"museum"`
	Details      string `gorm:"type:text;" json:"details"`
	Rubbing      string `gorm:"type:varchar(255);" json:"rubbing"`
	Picture3D    string `gorm:"type:varchar(255);" json:"picture3D"`
	View         uint32 `gorm:"type:int(11); not null;" json:"view"`
	Download     uint32 `gorm:"type:int(11); not null;" json:"download"`
}

func (*Image) TableName() string {
	return "images"
}

func RecordImage(count int) (images []Image, total int64, err error) {
	res := database.DB.Model(&Image{}).Find(&images).Order("view desc").Limit(count)
	total = res.RowsAffected
	err = res.Error
	return
}

func GetImageByCount(count int) (images []Image, total int64, err error) {
	res := database.DB.Model(&Image{}).Find(&images).Order("created_at desc").Limit(count)
	total = res.RowsAffected
	err = res.Error
	return
}

func (*Image) Index() string {
	return "images"

}
func (*Image) Mapping() string {
	return `
{
  "settings": {
    "index.refresh_interval": "5s",
    "number_of_shards": 1,
    "analysis": {
      "analyzer": {
        "ik_smart": {
          "tokenizer": "ik_smart"
        },
        "ik_max_word": {
          "tokenizer": "ik_max_word"
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "name": {
        "type": "text",
        "analyzer": "ik_max_word",
        "search_analyzer": "ik_smart"
      },
      "artist": {
        "type": "text",
        "analyzer": "ik_max_word",
        "search_analyzer": "ik_smart"
      },
      "details": {
        "type": "text",
        "analyzer": "ik_max_word",
        "search_analyzer": "ik_smart"
      },
      "year": {
        "type": "keyword"
      },
      "type": {
        "type": "keyword"
      },
      "specificType": {
        "type": "keyword"
      },
      "size": {
        "type": "keyword"
      },
      "museum": {
        "type": "keyword"
      },
      "view": {
        "type": "integer"
      },
      "download": {
        "type": "integer"
      }
    }
  }
}
	`
}
