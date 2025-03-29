package producer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"real-time-order-processing-system/config"
	"real-time-order-processing-system/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
)

func init(){
	config.InitDB()
	producer = createKafkaProducer()
}

func createKafkaProducer() *kafka.Producer{
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatal("Failed to create kafka producer: ", err)
	}

	return producer
}

func createOrder(w http.ResponseWriter, r *http.Request){
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	order.Status = "Pending"

	config.DB.Create(&order)

	orderJSON, _ := json.Marshal(order)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &"orders", Partition: kafka.PartitionAny},
		Value: orderJSON,
	}, nil)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/orders", createOrder).Methods("POST")

	fmt.Println("Producer running on port 8080")
	http.ListenAndServe(":8080", router)
}