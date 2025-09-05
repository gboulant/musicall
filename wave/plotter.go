package wave

import (
	"io"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// chartdata creates a echart timeseries from a float dataset
func chartdata(samples []float64, samplerate int) (xdata []float64, ydata []opts.LineData) {
	xdata = make([]float64, len(samples))
	ydata = make([]opts.LineData, len(samples))
	for i := range samples {
		xdata[i] = float64(i) / float64(samplerate)
		ydata[i] = opts.LineData{Value: samples[i]}
	}
	return xdata, ydata
}

// chartline creates a echart chartline from the specified echart timeseries
func chartline(xdata []float64, ydata []opts.LineData, label string) *charts.Line {

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
	line.AddSeries(label, ydata)

	//smooth := true
	//line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: &smooth}))
	return line
}

type WavePlotter struct {
	samplerate int
	chart      *charts.Line
}

func NewPlotter(samplerate int) *WavePlotter {
	return &WavePlotter{samplerate: samplerate}
}

func (p *WavePlotter) AddSeries(samples []float64, label string) error {
	xdata, ydata := chartdata(samples, p.samplerate)
	if p.chart == nil {
		p.chart = chartline(xdata, ydata, label)
		return nil
	}
	p.chart.AddSeries(label, ydata)
	return nil
}

// Plot draws the chart into the specified writer. The writer could be an html
// file writer or an http writer
func (p *WavePlotter) Plot(w io.Writer) error {
	return p.chart.Render(w)
}

// Save creates an html file that display the chart. Technically speaking, it
// executes the Plot function with a writer opened on the specified output file.
func (p *WavePlotter) Save(htmlpath string) error {
	f, err := os.Create(htmlpath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := p.Plot(f); err != nil {
		return err
	}
	log.Printf("Result available in file %s", htmlpath)
	return nil
}

// PlotToFile creates a HTML file that displays the chart of the specified
// sample. It is a short instruction for handling a WavePlotter in the case
// where you have a single dataset to plot.
func PlotToFile(htmlpath string, samples []float64, samplerate int, label string) error {
	p := NewPlotter(samplerate)
	p.AddSeries(samples, label)
	return p.Save(htmlpath)
}
