package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Name        string `json:"name"`
	LastName    string `json:"lastname"`
	DateOfBirth string `json:"dateofbirth"`
}

func exercise2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var body Request
		json.NewDecoder(r.Body).Decode(&body)
		fmt.Fprintf(w, "Hello %s %s born %s.", body.Name, body.LastName, body.DateOfBirth)
	}
}

func main() {
	http.HandleFunc("/hehe", exercise2)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
