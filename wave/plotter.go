package wave

import (
	"io"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type WavePlotter struct {
	chart *charts.Line
}

func applyDefaultOptions(chart *charts.Line) {
	// set some global options like Title/Legend/ToolTip or anything else
	var zoomstart float32 = 0. // percentage of the window range
	var zoomend float32 = 100. // percentage of the window range

	chart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{Formatter: "{value}s"},
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

	smooth := true
	chart.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
		Smooth: &smooth,
	}))
}

func NewPlotter() *WavePlotter {
	chart := charts.NewLine()
	applyDefaultOptions(chart)
	return &WavePlotter{chart: chart}
}

func (p *WavePlotter) SetXFormatter(formatter string) {
	p.chart.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{Formatter: types.FuncStr(formatter)},
		}),
	)
}

func (p *WavePlotter) SetTitle(title string) {
	p.chart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
	)
}

func (p *WavePlotter) AddLineXYValues(x, y []float64, label string) {
	data := make([]opts.LineData, len(y))
	for i := range y {
		data[i] = opts.LineData{Value: []float64{x[i], y[i]}, Symbol: "none"}
	}
	p.chart.AddSeries(label, data)
}

func (p *WavePlotter) AddLineTimedValues(samples, times []float64, label string) {
	p.AddLineXYValues(times, samples, label)
}

func (p *WavePlotter) AddLineSampledValues(samples []float64, samplerate int, label string) {
	data := make([]opts.LineData, len(samples))
	var timestamp float64
	for i := range samples {
		timestamp = float64(i) / float64(samplerate)
		data[i] = opts.LineData{Value: []float64{timestamp, samples[i]}, Symbol: "none"}
	}
	p.chart.AddSeries(label, data)
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
	p := NewPlotter()
	p.AddLineSampledValues(samples, samplerate, label)
	return p.Save(htmlpath)
}
