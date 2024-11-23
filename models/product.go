package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required,min=3"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Category    string  `json:"category" validate:"required,alpha"`
	ImagePath   string  `json:"image_path"`
}
