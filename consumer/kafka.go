package main

import (
	"encoding/json"
	"fmt"
	"log"
	"real-time-order-processing-system/config"
	"real-time-order-processing-system/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartKafkaConsumer() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id": "order-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal("Failed to create kafka consumer:", err)
	}

	err = consumer.Subscribe("orders", nil)
	if err != nil {
		log.Fatal("Failed to subscribe to topic orders:",err)
	}
	fmt.Println("Kafka consumer started, listening on orders...")

	for{
		message, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Println("Consumer error:",err)
			continue	
		}

		var order models.Order
		if err := json.Unmarshal(message.Value, &order); err != nil {
			log.Println("Failed to parse order JSON:",err)
			continue
		} 

		if err := config.DB.Model(&order).Where("id = ?", order.ID).Update("status", "Processed").Error; err != nil{
			log.Println("Failed to update order in DB:", err)
		} else {
			fmt.Printf("Order ID %d processed\n", order.ID)
		}
	}
}
