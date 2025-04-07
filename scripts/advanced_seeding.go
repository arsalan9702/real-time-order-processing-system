package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"gorm.io/gorm"
)

// Order represents the order structure matching your API
// Note: gorm.Model is not needed in the JSON representation
// gorm.Model contains ID, CreatedAt, UpdatedAt, and DeletedAt fields
// which are handled by the server
type Order struct {
	gorm.Model
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status,omitempty"`
}

func main() {
	// Define command line flags
	numOrders := flag.Int("count", 1000, "Number of orders to generate")
	baseURL := flag.String("url", "http://localhost:8080", "Base URL of the API")
	// delay := flag.Int("delay", 100, "Delay between requests in milliseconds")
	minUserID := flag.Int("min-user", 1, "Minimum user ID")
	maxUserID := flag.Int("max-user", 1000, "Maximum user ID")
	minProductID := flag.Int("min-product", 1, "Minimum product ID")
	maxProductID := flag.Int("max-product", 100, "Maximum product ID")
	minQuantity := flag.Int("min-quantity", 1, "Minimum quantity")
	maxQuantity := flag.Int("max-quantity", 10, "Maximum quantity")
	verbose := flag.Bool("verbose", false, "Print detailed output")
	flag.Parse()
	
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	fmt.Printf("Sending %d random orders to %s/orders\n", *numOrders, *baseURL)
	
	// Count successful and failed requests
	successful := 0
	failed := 0
	
	startTime := time.Now()
	
	// Generate and send random orders
	for i := 0; i < *numOrders; i++ {
		// Create a random order
		order := Order{
			UserID:    rand.Intn(*maxUserID-*minUserID+1) + *minUserID,
			ProductID: rand.Intn(*maxProductID-*minProductID+1) + *minProductID,
			Quantity:  rand.Intn(*maxQuantity-*minQuantity+1) + *minQuantity,
		}
		
		// Convert order to JSON
		orderJSON, err := json.Marshal(order)
		if err != nil {
			fmt.Printf("Error marshaling order %d: %v\n", i+1, err)
			failed++
			continue
		}
		
		// Send POST request
		resp, err := http.Post(
			fmt.Sprintf("%s/orders", *baseURL),
			"application/json",
			bytes.NewBuffer(orderJSON),
		)
		
		if err != nil {
			fmt.Printf("Error sending order %d: %v\n", i+1, err)
			failed++
			continue
		}
		
		// Check response status
		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
			fmt.Printf("Order %d received non-success status code: %d\n", i+1, resp.StatusCode)
			resp.Body.Close()
			failed++
			continue
		}
		
		// Process response
		var createdOrder map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&createdOrder); err != nil {
			fmt.Printf("Error decoding response for order %d: %v\n", i+1, err)
			resp.Body.Close()
			failed++
			continue
		}
		resp.Body.Close()
		
		successful++
		
		// Print result only if verbose flag is set
		if *verbose {
			fmt.Printf("Order %d created with ID: %v (User: %d, Product: %d, Quantity: %d)\n", 
				i+1, createdOrder["ID"], order.UserID, order.ProductID, order.Quantity)
		} else if i%10 == 0 || i == *numOrders-1 {
			// Print progress update every 10 orders
			fmt.Printf("Progress: %d/%d orders processed\n", i+1, *numOrders)
		}
		
		// Add a delay between requests
		// time.Sleep(time.Duration(*delay) * time.Millisecond)
	}
	
	elapsedTime := time.Since(startTime)
	
	fmt.Printf("\nSeeding completed in %v!\n", elapsedTime)
	fmt.Printf("Successful: %d, Failed: %d, Total: %d\n", successful, failed, *numOrders)
	fmt.Printf("Throughput: %.2f requests/second\n", float64(*numOrders)/elapsedTime.Seconds())
}