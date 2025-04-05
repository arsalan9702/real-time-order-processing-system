package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    int    `json:"user_id" gorm:"not null"`
	ProductID int    `json:"product_id" gorm:"not null"`
	Quantity  int    `json:"quantity" gorm:"not null"`
	Status    string `json:"status" gorm:"type:varchar(20);default:'pending'"`
}
