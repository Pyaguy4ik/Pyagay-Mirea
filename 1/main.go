package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()
	// Текстовый ответ
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	// ...
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(user{
			ID:   uuid.NewString(), // теперь реальный UUID
			Name: "Gopher"})
	})
	addr := ":8080"
	log.Printf("Starting on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
