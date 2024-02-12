package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

// Reservation struct represents a reservation record
type Reservation struct {
	ID            int    `json:"id,omitempty"`
	CarID         int    `json:"car_id,omitempty"`
	CustomerID    int    `json:"customer_id,omitempty"`
	StartDate     string `json:"start_date,omitempty"`
	EndDate       string `json:"end_date,omitempty"`
	Status        string `json:"status,omitempty"`
	PaymentStatus string `json:"payment_status,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

var reservationTableName = "Reservation"

// AddReservation creates a new reservation record
func AddReservation(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var reservation *Reservation
	_ = json.NewDecoder(request.Body).Decode(&reservation)
	// Perform database operation to add reservation
	dbClient.Table(reservationTableName).Create(&reservation)

	// Simulated database operation
	logger.Infof("Inserted data")
	json.NewEncoder(response).Encode(reservation)
}

// GetReservationByID retrieves a specific reservation's details by its ID
func GetReservationByID(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var reservation *Reservation
	// Perform database operation to fetch reservation by ID
	dbClient.Table(reservationTableName).First(&reservation, id)

	// Simulated database operation
	logger.Infof("Fetched Reservation With Id:" + id)
	json.NewEncoder(response).Encode(reservation)
}

// GetReservations retrieves a list of all reservations
func GetReservations(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	var reservations []*Reservation
	// Perform database operation to fetch all reservations
	dbClient.Table(reservationTableName).Find(&reservations)

	// Simulated database operation
	logger.Infof("Fetched all Reservations")
	json.NewEncoder(response).Encode(reservations)
}

// UpdateReservation updates an existing reservation's details by its ID
func UpdateReservation(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var reservation *Reservation
	_ = json.NewDecoder(request.Body).Decode(&reservation)
	// Perform database operation to update reservation
	dbClient.Table(reservationTableName).Model(&reservation).Updates(reservation)

	// Simulated database operation
	logger.Infof("Updated Reservation with Id:" + string(reservation.ID))
	json.NewEncoder(response).Encode("Updated Reservation with Id:" + string(reservation.ID))
}

// DeleteReservation deletes a specific reservation by its ID
func DeleteReservation(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var reservation *Reservation
	queryParams, err := url.ParseQuery(request.URL.RawQuery)
	if err != nil {
		http.Error(response, "Error parsing query parameters", http.StatusBadRequest)
		return
	}
	idValues, ok := queryParams["id"]
	if !ok || len(idValues) == 0 {
		http.Error(response, "Missing or empty 'id' parameter", http.StatusBadRequest)
		return
	}
	id := idValues[0]
	// Perform database operation to delete reservation by ID
	dbClient.Table(reservationTableName).Where("id = ?", id).Delete(&reservation)

	// Simulated database operation
	logger.Infof("Deleted Reservation with Id:" + id)
	json.NewEncoder(response).Encode("Deleted Reservation with Id:" + id)
}
