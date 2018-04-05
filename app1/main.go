package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var response []byte

func init() {
	rand.Seed(time.Now().UnixNano())
	response = []byte(fmt.Sprintf("Hello, World! My magic number is %v\n", rand.Int()))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/first", Handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
