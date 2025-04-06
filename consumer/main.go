package main

import (
	"fmt"
	"real-time-order-processing-system/config"
	"real-time-order-processing-system/internal/kafka"
)

func main() {
	config.InitDB()

	fmt.Println("Consumer service starting")
	kafka.InitKafkaProducer()
	kafka.StartKafkaConsumer()
}
