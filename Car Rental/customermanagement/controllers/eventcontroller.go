package controllers

import (
	"encoding/json"
	"net/http"

	"com.customermanagement/auth"
	"com.customermanagement/handler"
	"com.customermanagement/middleware"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

type EventController struct{}

func (t EventController) RegisterRoutes(r *mux.Router) {
	// Apply CORS middleware to all routes under this router
	r.Use(middleware.CORSMiddleware)

	// Register routes
	r.Handle("/customer", auth.Protect(http.HandlerFunc(handler.CreateCustomer))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/customer", auth.Protect(http.HandlerFunc(handler.GetCustomers))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/customer/{id}", auth.Protect(http.HandlerFunc(handler.GetCustomerByID))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/customer/{id}", auth.Protect(http.HandlerFunc(handler.UpdateCustomer))).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/customer/{id}", auth.Protect(http.HandlerFunc(handler.DeleteCustomer))).Methods(http.MethodDelete, http.MethodOptions)

	// Other routes...

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

	r.Handle("/rest/services/customermanagement", auth.Protect(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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
