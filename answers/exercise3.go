package main

import (
	"encoding/json"
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
	personalKey := "PUT YOUR PERSONAL LINK"
	var body []Info
	requ, _ := http.Get(personalKey + "/module/board/?format=json&start=2022-05-09&end=2022-05-16")
	json.NewDecoder(requ.Body).Decode(&body)
	fmt.Println(body)
}

func main() {
	http.HandleFunc("/calendar", exercise3)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
