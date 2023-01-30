package main

import (
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	Current_weather Weather `json:"current_weather"`
}

type Weather struct {
	Temperature   float32 `json:"temperature"`
	Windspeed     float32 `json:"windspeed"`
	Winddirection float32 `json:"winddirection"`
	Weathercode   int     `json:"weathercode"`
	Time          string  `json:"time"`
}

func exercise3(w http.ResponseWriter, r *http.Request) {
	// call to the https://open-meteo.com api and print out the info using the structures above
	return
}

func main() {
	http.HandleFunc("/weather", exercise3)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
