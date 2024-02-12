package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

// Payment struct represents a payment record
type Payment struct {
	ID            int     `json:"id,omitempty"`
	ReservationID int     `json:"reservation_id,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	PaymentDate   string  `json:"payment_date,omitempty"`
	PaymentMethod string  `json:"payment_method,omitempty"`
	Status        string  `json:"status,omitempty"`
	CreatedAt     string  `json:"created_at,omitempty"`
	UpdatedAt     string  `json:"updated_at,omitempty"`
}

var paymentTableName = "Payment"

// CreatePayment creates a new payment record
func CreatePayment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var payment *Payment
	_ = json.NewDecoder(request.Body).Decode(&payment)
	paymentData := Payment{
		ReservationID: payment.ReservationID,
		Amount:        payment.Amount,
		PaymentDate:   payment.PaymentDate,
		PaymentMethod: payment.PaymentMethod,
		Status:        payment.Status,
		CreatedAt:     payment.CreatedAt,
		UpdatedAt:     payment.UpdatedAt,
	}
	// Perform database operation to add payment
	dbClient.Table(paymentTableName).Create(&paymentData)

	// Simulated database operation
	logger.Infof("Inserted data")
	json.NewEncoder(response).Encode(paymentData)
}

// GetPaymentByID retrieves a specific payment's details by its ID
func GetPaymentByID(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var payment *Payment
	// Perform database operation to fetch payment by ID
	dbClient.Table(paymentTableName).First(&payment, id)

	// Simulated database operation
	logger.Infof("Fetched Payment With Id:" + id)
	json.NewEncoder(response).Encode(payment)
}

// GetPayments retrieves a list of all payments
func GetPayments(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	var payments []*Payment
	// Perform database operation to fetch all payments
	dbClient.Table(paymentTableName).Find(&payments)

	// Simulated database operation
	logger.Infof("Fetched all Payments")
	json.NewEncoder(response).Encode(payments)
}

// UpdatePayment updates an existing payment's details by its ID
func UpdatePayment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var payment *Payment
	_ = json.NewDecoder(request.Body).Decode(&payment)
	// Perform database operation to update payment
	dbClient.Table(paymentTableName).Model(&payment).Updates(payment)

	// Simulated database operation
	logger.Infof("Updated Payment with Id:" + string(payment.ID))
	json.NewEncoder(response).Encode("Updated Payment with Id:" + string(payment.ID))
}

// DeletePayment deletes a specific payment by its ID
func DeletePayment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	queryParams := request.URL.Query() // Use request.URL.Query() to parse query parameters directly
	idValues, ok := queryParams["id"]
	if !ok || len(idValues) == 0 {
		http.Error(response, "Missing or empty 'id' parameter", http.StatusBadRequest)
		return
	}
	id := idValues[0]
	// Perform database operation to delete payment by ID
	dbClient.Table(paymentTableName).Where("id = ?", id).Delete(&Payment{})

	// Simulated database operation
	logger.Infof("Deleted Payment with Id:" + id)
	json.NewEncoder(response).Encode("Deleted Payment with Id:" + id)
}
