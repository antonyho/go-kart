package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	items []Item
}
