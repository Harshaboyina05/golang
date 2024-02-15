package controllers

import (
	"encoding/json"
	"net/http"

	"com.paymentmanagement/auth"
	"com.paymentmanagement/handler"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

type EventController struct{}

func (t EventController) RegisterRoutes(r *mux.Router) {
	r.Handle("/payment", auth.Protect(http.HandlerFunc(handler.CreatePayment))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/payment", auth.Protect(http.HandlerFunc(handler.GetPayments))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/payment/{id}", auth.Protect(http.HandlerFunc(handler.GetPaymentByID))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/payment/{id}", auth.Protect(http.HandlerFunc(handler.UpdatePayment))).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/payment/{id}", auth.Protect(http.HandlerFunc(handler.DeletePayment))).Methods(http.MethodDelete, http.MethodOptions)

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
	r.Handle("/rest/services/paymentmanagement", auth.Protect(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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
