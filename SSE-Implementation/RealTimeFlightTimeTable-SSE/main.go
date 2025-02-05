package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Flight struct {
	FlightNo string `json:"flightNo"`
	Status   string `json:"status"`
	ETA      string `json:"eta"`
}

var statuses = []string{"On Time", "Delayed", "Departed", "Landed", "Cancelled"}

func generateFlightUpdates() Flight {
	flightNumbers := []string{"AI101", "EK202", "LH303", "BA404", "UA505"}
	return Flight{
		FlightNo: flightNumbers[rand.Intn(len(flightNumbers))],
		Status:   statuses[rand.Intn(len(statuses))],
		ETA:      time.Now().Add(time.Duration(rand.Intn(5)) * time.Minute).Format("15:04"),
	}
}

func flightUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}
	for {
		flightUpdate := generateFlightUpdates()
		data, _ := json.Marshal(flightUpdate)

		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()
		time.Sleep(3 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/flights", flightUpdatesHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
