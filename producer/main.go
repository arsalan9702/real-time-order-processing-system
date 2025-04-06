package main

import (
	"fmt"
	"net/http"
	"real-time-order-processing-system/config"
	"real-time-order-processing-system/internal/kafka"

	"github.com/gorilla/mux"
)

func init() {
	config.InitDB()
	config.MigrateDB()
	kafka.InitKafkaProducer()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", CreateOrder).Methods("POST")

	fmt.Println("Producer running on port 8080")
	http.ListenAndServe(":8080", router)
}
