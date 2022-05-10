package main

import (
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	Title_module string `json:"title_module"`
	Begin_event  string `json:"begin_event"`
	End_event    string `json:"end_event"`
	Seats        string `json:"seats"`
	Type_acti    string `json:"type_acti"`
	acti_title   string `json:"acti_title"`
	Begin_acti   string `json:"begin_acti"`
	End_acti     string `json:"end_acti"`
	Registered   int    `json:"registered"`
}

type Request struct {
	PersonalKey string `json:"personalKey"`
}

func exercise3(w http.ResponseWriter, r *http.Request) {
	// use this to extract your key from the request
	var req Request
	// use this to extract the calendar info from the intranet
	var body []Info
}

func main() {
	http.HandleFunc("/calendar", exercise3)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
