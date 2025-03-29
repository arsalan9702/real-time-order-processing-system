package main

import (
	"fmt"
	"net/http"
	"real-time-order-processing-system/config"
	"github.com/gorilla/mux"
)

func init(){
	config.InitDB()
	InitKafkaProducer()
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/orders", createOrder).Methods("POST")

	fmt.Println("Producer running on port 8080")
	http.ListenAndServe(":8080", router)
}