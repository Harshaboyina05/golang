package controllers

import (
	"encoding/json"
	"net/http"

	"com.reservationmanagement/auth"
	"com.reservationmanagement/handler"

	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

type EventController struct{}

func (t EventController) RegisterRoutes(r *mux.Router) {
	r.Handle("/reservation/AddReservation", auth.Protect(http.HandlerFunc(handler.AddReservation))).Methods(http.MethodPost)
	r.Handle("/reservation/GetReservations", auth.Protect(http.HandlerFunc(handler.GetReservations))).Methods(http.MethodGet)
	r.Handle("/reservation/{id}", auth.Protect(http.HandlerFunc(handler.GetReservationByID))).Methods(http.MethodGet)
	r.Handle("/reservation/{id}", auth.Protect(http.HandlerFunc(handler.UpdateReservation))).Methods(http.MethodPut)
	r.Handle("/reservation/{id}", auth.Protect(http.HandlerFunc(handler.DeleteReservation))).Methods(http.MethodDelete)

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
	r.Handle("/rest/services/reservationmanagement", auth.Protect(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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
