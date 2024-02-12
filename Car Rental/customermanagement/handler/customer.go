package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

// Customer struct represents a customer record
type Customer struct {
	ID          int    `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

var customerTableName = "Customer"

// CreateCustomer creates a new customer record
func CreateCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var customer *Customer
	_ = json.NewDecoder(request.Body).Decode(&customer)
	customerData := Customer{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
		CreatedAt:   customer.CreatedAt,
		UpdatedAt:   customer.UpdatedAt,
	}
	// Perform database operation to add customer
	dbClient.Table(customerTableName).Create(&customerData)

	// Simulated database operation
	logger.Infof("Inserted data")
	json.NewEncoder(response).Encode(customerData)
}

// GetCustomerByID retrieves a specific customer's details by its ID
func GetCustomerByID(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var customer *Customer
	// Perform database operation to fetch customer by ID
	dbClient.Table(customerTableName).First(&customer, id)

	// Simulated database operation
	logger.Infof("Fetched Customer With Id:" + id)
	json.NewEncoder(response).Encode(customer)
}

// GetCustomers retrieves a list of all customers
func GetCustomers(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	var customers []*Customer
	// Perform database operation to fetch all customers
	dbClient.Table(customerTableName).Find(&customers)

	// Simulated database operation
	logger.Infof("Fetched all Customers")
	json.NewEncoder(response).Encode(customers)
}

// UpdateCustomer updates an existing customer's details by its ID
func UpdateCustomer(response http.ResponseWriter, request *http.Request) {
	// Parse request body and decode JSON into a Customer object
	var customer *Customer
	params := mux.Vars(request)
	id := params["id"]
	if err := json.NewDecoder(request.Body).Decode(&customer); err != nil {
		// Handle JSON decoding error
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform database operation to update customer
	result := dbClient.Table(customerTableName).Where("id = ?", id).Updates(&customer)

	// Check for errors
	if result.Error != nil {
		http.Error(response, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		http.Error(response, "No customer found with the given ID", http.StatusNotFound)
		return
	}

	// Simulated database operation
	logger.Infof("Updated Customer with Id: %d", id)
	json.NewEncoder(response).Encode(map[string]string{"message": fmt.Sprintf("Updated Customer with Id: %s", id)})
}

// DeleteCustomer deletes a specific customer by its ID
func DeleteCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	queryParams := request.URL.Query() // Use request.URL.Query() to parse query parameters directly
	idValues, ok := queryParams["id"]
	if !ok || len(idValues) == 0 {
		http.Error(response, "Missing or empty 'id' parameter", http.StatusBadRequest)
		return
	}
	id := idValues[0]
	// Perform database operation to delete customer by ID
	result := dbClient.Table(customerTableName).Where("id = ?", id).Delete(&Customer{})
	if result.Error != nil {
		http.Error(response, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(response, "No customer found with the given ID", http.StatusNotFound)
		return
	}
	// Simulated database operation
	logger.Infof("Deleted Customer with Id: %s", id)
	json.NewEncoder(response).Encode(map[string]string{"message": "Deleted Customer with Id: " + id})
}
