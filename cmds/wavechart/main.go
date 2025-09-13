package main

/*
We illustrate in this program how to use the external package go-echarts to plot
the timeseries of a signal generated using the wave package.
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gboulant/musicall"
	"github.com/gboulant/musicall/wave"
)

const defaultExampleName string = "D01"

func init() {
	musicall.NewExample("D01", "Plot into a HTML file", demo01_appli)
	musicall.NewExample("D02", "Plot into a web browser", demo02_server)
}

func demo01_appli() error {
	a := 2.
	d := 10. // sec.
	r := 1000
	fmin := 1.  // Hz
	fmax := 10. // Hz

	p := wave.NewPlotter()

	s := wave.NewSweepFrequencySynthesizer(fmin, fmax, a, r)
	samples := s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "Sweep")

	f := 1.
	s = wave.NewSquareWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "Square")

	outfilepath := "output.wavechart.html"
	return p.Save(outfilepath)
}

func demo02_server() error {
	a := 2.
	d := 10. // sec.
	r := 1000
	fmin := 1.  // Hz
	fmax := 10. // Hz
	s := wave.NewSweepFrequencySynthesizer(fmin, fmax, a, r)
	samples := s.Synthesize(d)

	// Just to emulate a payload
	payload.samplerate = r
	payload.samples = samples
	payload.label = "Sweep"

	httpport := 8081
	address := fmt.Sprintf(":%d", httpport)
	http.HandleFunc("/", httpserver)
	log.Printf("Plot server is running on http://localhost:%d", httpport)
	return http.ListenAndServe(address, nil)
}

// ----------------------------------------------------------------
func main() {
	musicall.StartExampleApp(defaultExampleName)
}
