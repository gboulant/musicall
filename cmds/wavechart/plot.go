package main

import (
	"io"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// data creates a echart timeseries from a float dataset
func data(samples []float64, samplerate int) (xdata []float64, ydata []opts.LineData) {
	xdata = make([]float64, len(samples))
	ydata = make([]opts.LineData, len(samples))
	for i := range samples {
		xdata[i] = float64(i) / float64(samplerate)
		ydata[i] = opts.LineData{Value: samples[i]}
	}
	return xdata, ydata
}

// line plots the samples series with rendering into the specified
// writer. The writer could be an html file writer or an http writer.
func line(xdata []float64, ydata []opts.LineData) *charts.Line {

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
	return line
}

func plot(w io.Writer, samples []float64, samplerate int) error {
	xdata, ydata := data(samples, samplerate)
	chart := line(xdata, ydata)
	return chart.Render(w)
}

func plotToFile(htmlpath string, samples []float64, samplerate int) error {
	f, _ := os.Create(htmlpath)
	defer f.Close()

	if err := plot(f, samples, samplerate); err != nil {
		return err
	}
	log.Printf("Result available in file %s", htmlpath)
	return nil
}
