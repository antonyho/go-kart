package models

import (
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	ItemID uint   `gorm:"unique_index:idx_image_order"`
	Order  int    `gorm:"unique_index:idx_image_order"`
	Path   string // Relative path to the image file
}
