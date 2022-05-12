package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello World!")
	}
}

func wdym(w http.ResponseWriter, r *http.Request) {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	fmt.Fprintf(w, "Hello %s!", m["name"][0])
}

func main() {
	http.HandleFunc("/wdym", wdym)
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
