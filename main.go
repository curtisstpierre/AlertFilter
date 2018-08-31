package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Alert struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Customer   string `json:"customer,omitempty"`
	Host       string `json:"host,omitempty"`
	Datacenter string `json:"datacenter,omitempty"`
}

var alerts []Alert

// GetAlerts all from the alerts var
func GetAlerts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(alerts)
}

// GetAlert a single data
func GetAlert(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range alerts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Alert{})
}

// CreateAlert a new item
func CreateAlert(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Alert
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	alerts = append(alerts, person)
	json.NewEncoder(w).Encode(alerts)
}

// DeleteAlert an item
func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range alerts {
		if item.ID == params["id"] {
			alerts = append(alerts[:index], alerts[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(alerts)
	}
}

// fun main()
func main() {
	router := mux.NewRouter()
	alerts = append(alerts, Alert{ID: "1", Name: "John", Customer: "Doe", Host: "Instance1", Datacenter: "dc1"})
	alerts = append(alerts, Alert{ID: "2", Name: "Koko", Customer: "Doe", Host: "Instance1", Datacenter: "dc2"})
	alerts = append(alerts, Alert{ID: "3", Name: "Francis", Customer: "Sunday", Host: "Instance1", Datacenter: "dc1"})
	router.HandleFunc("/alerts", GetAlerts).Methods("GET")
	router.HandleFunc("/alerts/{id}", GetAlert).Methods("GET")
	router.HandleFunc("/alerts/{id}", CreateAlert).Methods("POST")
	router.HandleFunc("/alerts/{id}", DeleteAlert).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
