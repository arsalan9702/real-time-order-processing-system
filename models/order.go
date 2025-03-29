package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	Status    int `json:"status"`
}
