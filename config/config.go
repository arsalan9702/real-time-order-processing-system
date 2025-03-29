package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"real-time-order-processing-system/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=user password=password dbname=orders port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database")
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.Order{})
	if err != nil {
		return
	}
}
