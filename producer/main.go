package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"real-time-order-processing-system/config"
)

func init() {
	config.InitDB()
	config.MigrateDB()
	InitKafkaProducer()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", CreateOrder).Methods("POST")

	fmt.Println("Producer running on port 8080")
	http.ListenAndServe(":8080", router)
}
