package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Name    string
	Path    string
	Year    string
	Type    string
	Museum  string
	Details string
}
