package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Worlds")
	})
	router.HandleFunc("/employees", getEmployeesHandler).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)

}

func getEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	employees := getEmployees()
	employeesJSON, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(employeesJSON)
}
