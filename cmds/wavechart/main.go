package main

import (
	"log"
	"net/http"

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

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// Prepare the data
	samples, samplerate := testsamples()
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
	line.Render(w)
}

func test_server() error {
	http.HandleFunc("/", httpserver)
	log.Println("Plot server is running on http://localhost:8081")
	return http.ListenAndServe(":8081", nil)
}

func main() {
	test_server()
}
