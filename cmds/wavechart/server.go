package main

import (
	"fmt"
	"log"
	"net/http"

	"galuma.net/synthetic/wave"
)

type dataset struct {
	samples    []float64
	samplerate int
}

// payload is a local dataset to communicate data between the httpplot
// function (that start the http server) and the httpserver (that serve
// the plot html chart)
var payload dataset

// Plot the data using a server configuration (plot into the http writer)
func httpserver(w http.ResponseWriter, _ *http.Request) {
	// Prepare the data
	samples := payload.samples
	samplerate := payload.samplerate

	// Create the plot with rendering into the http writer
	if err := wave.Plot(w, samples, samplerate); err != nil {
		fmt.Fprintln(w, err)
	}
}

// Plot the data using an http configuration (http server)
func httpplot(samples []float64, samplerate int) error {
	payload.samplerate = samplerate
	payload.samples = samples
	http.HandleFunc("/", httpserver)
	log.Println("Plot server is running on http://localhost:8081")
	return http.ListenAndServe(":8081", nil)
}
