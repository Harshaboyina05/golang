package controllers

import (
	"encoding/json"
	"net/http"

	"com.carmanagement/auth"
	"com.carmanagement/handler"

	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

type EventController struct{}

func (t EventController) RegisterRoutes(r *mux.Router) {
	r.Handle("/car/AddCar", auth.Protect(http.HandlerFunc(handler.AddCar))).Methods(http.MethodPost)
	r.Handle("/car/GetCars", auth.Protect(http.HandlerFunc(handler.GetCars))).Methods(http.MethodGet)
	r.Handle("/car/{id}", auth.Protect(http.HandlerFunc(handler.GetCarByID))).Methods(http.MethodGet)
	r.Handle("/car/{id}", auth.Protect(http.HandlerFunc(handler.UpdateCar))).Methods(http.MethodPut)
	r.Handle("/car/{id}", auth.Protect(http.HandlerFunc(handler.DeleteCar))).Methods(http.MethodDelete)

	r.HandleFunc("/management/health/readiness", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "UP",
			"components": map[string]interface{}{
				"readinessState": map[string]interface{}{
					"status": "UP",
				},
			},
		})
	}).Methods(http.MethodGet)
	r.Handle("/rest/services/carmanagement", auth.Protect(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		logger.Infof("response sent")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"server": "UP"})
	}))).Methods(http.MethodGet)

	r.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode("helloworld")
	}).Methods(http.MethodGet)
	r.HandleFunc("/management/health/liveness", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "UP",
			"components": map[string]interface{}{
				"livenessState": map[string]interface{}{
					"status": "UP",
				},
			},
		})
	}).Methods(http.MethodGet)
}
