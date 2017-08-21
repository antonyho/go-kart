package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
}

type OrderedGroup struct {
	gorm.Model
	ParentID uint  // The ID of its parent Group
	Parent   Group `gorm:"AssociationForeignKey:ParentID;unique_index:idx_group_order"`
	RefID    uint  // The ID of its representing Group
	Ref      Group `gorm:"AssociationForeignKey:RefID"`
	Order    uint  `gorm:"unique_index:idx_group_order"`
}
