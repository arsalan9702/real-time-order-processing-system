package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Order represents the order structure matching your API
type Order struct {
	CustomerID        int     `json:"user_id"`
	ProductID         int     `json:"product_id"`
	Quantity          int     `json:"quantity"`
}


func main() {
	// Configure the seeding
	baseURL := "http://localhost:8080"
	numOrders := 10 // Number of orders to generate
	
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	fmt.Printf("Sending %d random orders to %s/orders\n", numOrders, baseURL)
	
	// Generate and send random orders
	for i := 0; i < numOrders; i++ {
		// Create a random order
		order := generateRandomOrder()
		
		// Convert order to JSON
		orderJSON, err := json.Marshal(order)
		if err != nil {
			fmt.Printf("Error marshaling order %d: %v\n", i+1, err)
			continue
		}
		
		// Send POST request
		resp, err := http.Post(
			fmt.Sprintf("%s/orders", baseURL),
			"application/json",
			bytes.NewBuffer(orderJSON),
		)
		
		if err != nil {
			fmt.Printf("Error sending order %d: %v\n", i+1, err)
			continue
		}
		
		// Process response
		var createdOrder map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&createdOrder)
		resp.Body.Close()
		
		// Print result
		fmt.Printf("Order %d created with ID: %v (Status code: %d)\n", 
			i+1, createdOrder["ID"], resp.StatusCode)
		
		// Add a small delay between requests to avoid overwhelming the server
		// time.Sleep(100 * time.Millisecond)
	}
	
	fmt.Println("Seeding completed!")
}

// generateRandomOrder creates a randomized order
func generateRandomOrder() Order {
	return Order{
		CustomerID:        rand.Intn(10000) + 1000,        // Random customer ID between 1000-10999
		ProductID:         rand.Intn(500) + 100,           // Random product ID between 100-599
		Quantity:          rand.Intn(5) + 1,     // Random address ID between 500-5499
	}
}