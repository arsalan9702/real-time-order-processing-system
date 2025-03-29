package producer

import (
	"encoding/json"
	"net/http"
	"real-time-order-processing-system/config"
	"real-time-order-processing-system/models"
)

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	order.Status = "Pending"

	config.DB.Create(&order)

	orderJSON, _ := json.Marshal(order)
	SendMessage("orders", orderJSON)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}