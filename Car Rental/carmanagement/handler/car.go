package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

// Car struct represents a car record
type Car struct {
	ID           int    `json:"id,omitempty"`
	Make         string `json:"make,omitempty"`
	Model        string `json:"model,omitempty"`
	Year         int    `json:"year,omitempty"`
	LicensePlate string `json:"license_plate,omitempty"`
	Color        string `json:"color,omitempty"`
	Mileage      int    `json:"mileage,omitempty"`
	Status       string `json:"status,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

var carTableName = "Car"

// AddCar adds a new car record
func AddCar(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var car *Car
	_ = json.NewDecoder(request.Body).Decode(&car)
	carData := Car{
		Make:         car.Make,
		Model:        car.Model,
		Year:         car.Year,
		LicensePlate: car.LicensePlate,
		Color:        car.Color,
		Mileage:      car.Mileage,
		Status:       car.Status,
		CreatedAt:    car.CreatedAt,
		UpdatedAt:    car.UpdatedAt,
	}
	// Perform database operation to add car
	dbClient.Table(carTableName).Create(&carData)

	// Simulated database operation
	logger.Infof("Inserted data")
	json.NewEncoder(response).Encode(carData)
}

// GetCarByID retrieves a specific car's details by its ID
func GetCarByID(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var car *Car
	// Perform database operation to fetch car by ID
	dbClient.Table(carTableName).First(&car, id)

	// Simulated database operation
	logger.Infof("Fetched Car With Id:" + id)
	json.NewEncoder(response).Encode(car)
}

// GetCars retrieves a list of all cars
func GetCars(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type", "application/json")
	var cars []*Car
	// Perform database operation to fetch all cars
	dbClient.Table(carTableName).Find(&cars)

	// Simulated database operation
	logger.Infof("Fetched all Cars")
	json.NewEncoder(response).Encode(cars)
}

// UpdateCar updates an existing car's details by its ID
func UpdateCar(response http.ResponseWriter, request *http.Request) {
	// Parse request body and decode JSON into a Car object
	var car *Car
	params := mux.Vars(request)
	id := params["id"]
	if err := json.NewDecoder(request.Body).Decode(&car); err != nil {
		// Handle JSON decoding error
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform database operation to update car
	result := dbClient.Table(carTableName).Where("id = ?", id).Updates(&car)

	// Check for errors
	if result.Error != nil {
		http.Error(response, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		http.Error(response, "No car found with the given ID", http.StatusNotFound)
		return
	}

	// Simulated database operation
	logger.Infof("Updated Car with Id: %d", id)
	json.NewEncoder(response).Encode(map[string]string{"message": fmt.Sprintf("Updated Car with Id: %s", id)})
}

// DeleteCar deletes a specific car by its ID
func DeleteCar(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	queryParams := request.URL.Query() // Use request.URL.Query() to parse query parameters directly
	idValues, ok := queryParams["id"]
	if !ok || len(idValues) == 0 {
		http.Error(response, "Missing or empty 'id' parameter", http.StatusBadRequest)
		return
	}
	id := idValues[0]
	// Perform database operation to delete car by ID
	result := dbClient.Table(carTableName).Where("id = ?", id).Delete(&Car{})
	if result.Error != nil {
		http.Error(response, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(response, "No car found with the given ID", http.StatusNotFound)
		return
	}
	// Simulated database operation
	logger.Infof("Deleted Car with Id: %s", id)
	json.NewEncoder(response).Encode(map[string]string{"message": "Deleted Car with Id: " + id})
}
