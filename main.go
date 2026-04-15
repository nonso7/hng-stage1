package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response structs
type MessageResponse struct {
	Message string `json:"message"`
}

type MeResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Github string `json:"github"`
}

func jsonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Message: "API is running"})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Message: "healthy"})
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MeResponse{
		Name:   "Chinonso Nwajagu",
		Email:  "felixmedia78@gmail.com",
		Github: "https://github.com/nonso7",
	})
}

func main() {
	http.HandleFunc("/", jsonMiddleware(rootHandler))
	http.HandleFunc("/health", jsonMiddleware(healthHandler))
	http.HandleFunc("/me", jsonMiddleware(meHandler))

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
