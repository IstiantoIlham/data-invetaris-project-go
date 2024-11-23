package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity" validate:"gt=0"`
	Status    string  `json:"status" validate:"required,oneof=Pending Completed Cancelled"`
}
