package main

import (
	"fmt"
	"net/http"

	"github.com/gboulant/musicall/wave"
)

type dataset struct {
	samples    []float64
	samplerate int
	label      string
}

// payload is a local dataset to communicate data between the httpplot
// function (that start the http server) and the httpserver (that serve
// the plot html chart). It is to emulate a post REST function with a dataset in
// the payload and the plot of the chart as a response
var payload dataset

// Plot the data using a server configuration (plot into the http writer)
func httpserver(w http.ResponseWriter, _ *http.Request) {
	// Prepare the data
	samples := payload.samples
	samplerate := payload.samplerate
	label := payload.label

	p := wave.NewPlotter()
	p.AddLineSampledValues(samples, samplerate, label)

	// Create the plot with rendering into the http writer
	if err := p.Plot(w); err != nil {
		fmt.Fprintln(w, err)
	}
}
