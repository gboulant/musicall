package main

/*
We illustrate in this program how to use the external package go-echarts to plot
the timeseries of a signal generated using the wave package.
*/

import (
	"fmt"
	"os"

	"github.com/gboulant/musicall/wave"
)

// Plot the data using a standalone configuration (create a local html file)
func fileplot(samples []float64, samplerate int) error {
	outfilepath := "output.wavechart.html"
	return wave.PlotToFile(outfilepath, samples, samplerate, "Wave Chart")
}

// ----------------------------------------------------------------
func main() {
	var program func(samples []float64, samplerate int) error

	// choose the configuration: 1/ html file, or 2/ http server
	//program = httpplot
	program = fileplot

	samplerate := 440
	//samples, samplerate := d01_KarplusStrong(samplerate)
	samples := d02_sweepfrequency(samplerate, false)

	if err := program(samples, samplerate); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
