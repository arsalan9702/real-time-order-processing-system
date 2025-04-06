package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"real-time-order-processing-system/internal/kafka"
	"real-time-order-processing-system/models"
	"time"

)

func main() {
	// Start Kafka consumer in a goroutine
	kafka.StartKafkaConsumer()


	// Initialize Kafka producer
	kafka.InitKafkaProducer()

	// Seed orders
	rand.Seed(time.Now().UnixNano())
	statuses := []string{"Pending", "Shipped", "Delivered", "Cancelled"}

	for i := 1; i <= 10; i++ {
		order := models.Order{
			UserID:    rand.Intn(100) + 1,
			ProductID: rand.Intn(50) + 1,
			Quantity:  rand.Intn(5) + 1,
			Status:    statuses[rand.Intn(len(statuses))],
		}
		orderJSON, err := json.Marshal(order)
		if err != nil {
			log.Println("Failed to serialize order:", err)
			continue
		}
		kafka.SendMessage("orders", orderJSON)
		fmt.Printf("🚚 Sent order #%d to Kafka\n", i)
	}

	fmt.Println("🌱 Seeding complete.")

}
