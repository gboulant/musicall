package main

/*
We illustrate in this program how to use the external package go-echarts to plot
the timeseries of a signal generated using the wave package.
*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"galuma.net/synthetic/wave"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func testsamples() (samples []float64, samplerate int) {
	f := 5.
	a := 2.
	d := 4.
	samples = wave.KarplusStrongSignal(f, a, d)
	samplerate = wave.DefaultSampleRate
	return samples, samplerate
}

// wavedata creates a echart timeseries from a float dataset
func wavedata(samples []float64, samplerate int) (xdata []float64, ydata []opts.LineData) {
	xdata = make([]float64, len(samples))
	ydata = make([]opts.LineData, len(samples))
	for i := range samples {
		xdata[i] = float64(i) / float64(samplerate)
		ydata[i] = opts.LineData{Value: samples[i]}
	}
	return xdata, ydata
}

func plotdata(w io.Writer, samples []float64, samplerate int) error {
	// Prepare the data
	xdata, ydata := wavedata(samples, samplerate)

	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	var zoomstart float32 = 20. // percentage of the window range
	var zoomend float32 = 80.   // percentage of the window range

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "MusicHall - Synthetic",
			Subtitle: "Line chart rendering wave functions",
		}),
		charts.WithDataZoomOpts(
			opts.DataZoom{
				Type:       "inside",
				Start:      zoomstart,
				End:        zoomend,
				XAxisIndex: []int{0},
			},
			opts.DataZoom{
				Type:       "slider",
				Start:      zoomstart,
				End:        zoomend,
				XAxisIndex: []int{0},
			},
			opts.DataZoom{
				Type:       "slider",
				YAxisIndex: []int{0},
			},
		),
	)

	// Put data into instance
	line.SetXAxis(xdata)
	line.AddSeries("Sine Wave", ydata)

	//smooth := true
	//line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: &smooth}))
	return line.Render(w)
}

// ----------------------------------------------------------------
// Plot the data using a standalone configuration (create a local html file)
func standalone() error {
	// Prepare the data
	samples, samplerate := testsamples()
	// Create the plot with rendering onto an html local file

	// Create the plot with rendering into a local html file
	outfilepath := "output.wavechart.html"
	f, _ := os.Create(outfilepath)
	defer f.Close()

	if err := plotdata(f, samples, samplerate); err != nil {
		return err
	}
	log.Printf("Result available in file %s", outfilepath)
	return nil
}

// ----------------------------------------------------------------
// Plot the data using a server configuration (plot into the http writer)
func httpserver(w http.ResponseWriter, _ *http.Request) {
	// Prepare the data
	samples, samplerate := testsamples()

	// Create the plot with rendering into the http writer
	plotdata(w, samples, samplerate)
}

func server() error {
	http.HandleFunc("/", httpserver)
	log.Println("Plot server is running on http://localhost:8081")
	return http.ListenAndServe(":8081", nil)
}

// ----------------------------------------------------------------
func main() {
	var program func() error

	program = server
	program = standalone // comment this line for activating the server mode

	if err := program(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
