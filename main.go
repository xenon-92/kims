package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

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
	router.HandleFunc("/graceKill", killServerGracefully).Methods(http.MethodPost)
	router.HandleFunc("/kill", killServer).Methods(http.MethodPost)

	router.HandleFunc("/livez", liveCheckHandler).Methods(http.MethodGet)
	router.HandleFunc("/livezSlow", liveCheckHandlerSlow).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)

	log.Println("Server started on port 8080")
}

func liveCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Live Check OK")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Live Check OK"))
}

func liveCheckHandlerSlow(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	log.Println("Live Check slow OK")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Live Check Slow OK"))
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

func killServerGracefully(w http.ResponseWriter, r *http.Request) {
	log.Println("Server is gracefully shutting down...")
	os.Exit(0)
	log.Println("Server has been gracefully shut down.")
}

func killServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Server is shutting down...")
	os.Exit(1)
	log.Println("Server has been shut down.")
}
