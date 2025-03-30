package producer

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

func InitKafkaProducer() {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		log.Fatal("Failed to create Kafka producer: ", err)
	}

	fmt.Println("Kafka Producer Initialized")
}

func SendMessage(topic string, message []byte) {
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          message,
	}, nil)

	if err != nil {
		log.Println("Failed to send kafka message: ", err)
	} else {
		fmt.Println("Message sent to kafka: ", string(message))
	}
}
