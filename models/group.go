package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Name     string
	Exhibits []Exhibit `gorm:"-"` // Ignored by GORM and constructed by query
}

type Exhibit interface {
	Place() uint
}

type OrderedGroup struct {
	gorm.Model
	ParentID uint  // The ID of its parent Group
	Parent   Group `gorm:"AssociationForeignKey:ParentID;unique_index:idx_group_order"`
	RefID    uint  // The ID of its representing Group
	Ref      Group `gorm:"AssociationForeignKey:RefID"`
	Order    uint  `gorm:"unique_index:idx_group_order"`
}

func (g OrderedGroup) Place() uint {
	return g.Order
}
