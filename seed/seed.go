package main

import (
	"fmt"
	"math/rand"
	"real-time-order-processing-system/config"
	"real-time-order-processing-system/models"
	"time"
)

func main() {
	// Initialize DB
	config.InitDB()

	// Optionally migrate
	config.MigrateDB()

	// Seed
	rand.Seed(time.Now().UnixNano())

	statuses := []string{"Pending", "Shipped", "Delivered", "Cancelled"}

	for i := 1; i <= 10; i++ {
		order := models.Order{
			UserID:    rand.Intn(100) + 1,
			ProductID: rand.Intn(50) + 1,
			Quantity:  rand.Intn(5) + 1,
			Status:    statuses[rand.Intn(len(statuses))],
		}
		config.DB.Create(&order)
		fmt.Printf("Inserted Order #%d: %+v\n", i, order)
	}

	fmt.Println("🌱 Seeding complete.")
}
