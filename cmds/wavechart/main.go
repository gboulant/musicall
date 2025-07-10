package main

/*
We illustrate in this program how to use the external package go-echarts to plot
the timeseries of a signal generated using the wave package.
*/

import (
	"fmt"
	"os"
)

// Plot the data using a standalone configuration (create a local html file)
func fileplot(samples []float64, samplerate int) error {
	outfilepath := "output.wavechart.html"
	return plotToFile(outfilepath, samples, samplerate)
}

// ----------------------------------------------------------------
func main() {
	var program func(samples []float64, samplerate int) error

	// choose the configuration: 1/ html file, or 2/ http server
	//program = httpplot
	program = fileplot

	//samples, samplerate := d01_KarplusStrong()
	samples, samplerate := d02_sweepfrequency(false)

	if err := program(samples, samplerate); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
