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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rand.Seed(time.Now().UnixNano())
	response = []byte(fmt.Sprintf("Hallo, Welt! Meine magische Zahl ist %v\n", rand.Int()))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
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
	mux.HandleFunc("/second", Handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
