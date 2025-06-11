package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
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

	router.HandleFunc("/info", getSystemInfo).Methods(http.MethodGet)

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

func getSystemInfo(w http.ResponseWriter, r *http.Request) {
	remoteAddr := r.RemoteAddr

	// Split host and port
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Fprintf(w, "Error parsing address: %v", err)
		return
	}

	fmt.Fprintf(w, "Client IP: %s\n", host)
	fmt.Fprintf(w, "Remote Address: %s\n", remoteAddr)

}
