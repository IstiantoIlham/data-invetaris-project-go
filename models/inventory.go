package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Stock     int     `json:"stock" validate:"gte=0"`
}
