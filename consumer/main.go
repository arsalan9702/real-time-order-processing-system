package main

import (
	"fmt"
	"real-time-order-processing-system/config"
)

func main() {
	config.InitDB()

	fmt.Println("Consumer service starting")
	StartKafkaConsumer()
}
