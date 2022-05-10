package main

import (
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
		// exercise 2
	}
}

func main() {
	http.HandleFunc("/wdym", exercise2)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
