package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Worlds")
	})

	router.HandleFunc("/employees", getEmployeesHandler).Methods(http.MethodGet)

	router.HandleFunc("/health", healthCheckHandler).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)

	log.Println("Server started on port 8080")
}

func getEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	employees := getEmployees()
	log.Println(employees)
	employeesJSON, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(employeesJSON)
}
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
