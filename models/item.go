package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Item struct {
	gorm.Model
	Name         string
	Description  string
	Price        float64
	Images       []Image
	Availability uint
	Begin        time.Time
	End          time.Time
	//Listed       bool
}

type OrderedItem struct {
	gorm.Model
	ParentID uint  // The ID of its parent Group
	Parent   Group `gorm:"AssociationForeignKey:ParentID;unique_index:idx_item_order"`
	RefID    uint  // The ID of its representing Group
	Ref      Item  `gorm:"AssociationForeignKey:RefID"`
	Order    uint  `gorm:"unique_index:idx_item_order"`
}

func (i OrderedItem) Place() uint {
	return i.Order
}
