package wave

import (
	"os"
	"testing"
)

func d01_KarplusStrong(r int) []float64 {
	f := 440.
	a := 2.
	d := 4.
	return KarplusStrongSignal(f, a, d, r)
}

func d02_sweepfrequency(r int, reverse bool) []float64 {
	a := 2.
	d := 10.   // sec.
	fmin := 1. // Hz
	fmax := 5. // Hz
	return SweepFrequencySignal(fmin, fmax, a, d, r, reverse)
}

func decimate(samples []float64, samplerate int, step int) ([]float64, int) {
	return Decimate(samples, step), samplerate / step
}

func TestPlotToFile(t *testing.T) {
	// We generate the signal with the default sample rate (so that it
	// can be played if needed), but we decimate the signal to plot, to
	// make a plot not to heavy to display.
	samplerate := DefaultSampleRate
	decimatestep := 100

	samples := d01_KarplusStrong(samplerate)
	outfilepath := "output.d01_KarplusStrong.html"
	pltsamples, pltsamplerate := decimate(samples, samplerate, decimatestep)
	if err := PlotToFile(outfilepath, pltsamples, pltsamplerate, "KarplusStrong"); err != nil {
		t.Error(err)
	}

	samples = d02_sweepfrequency(samplerate, false)
	outfilepath = "output.d02_sweepfrequency.html"
	pltsamples, pltsamplerate = decimate(samples, samplerate, decimatestep)
	if err := PlotToFile(outfilepath, pltsamples, pltsamplerate, "sweepfrequency"); err != nil {
		t.Error(err)
	}
}

func TestWaveSynthesizers(t *testing.T) {
	f := 10.
	a := 2.
	d := 4.
	r := int(1000 * f) // 100 points by cycle

	s := NewSineWaveSynthesizer(f, a, r)
	samples := s.Synthesize(d)
	xdata, ydata := data(samples, r)
	chart := line(xdata, ydata, "Sine")

	s = NewSquareWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	_, ydata = data(samples, r)
	chart.AddSeries("Square", ydata)

	s = NewKarplusStrongSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	_, ydata = data(samples, r)
	chart.AddSeries("KarplusStrong", ydata)

	s = NewTriangleWaveSynthesizer(f, a, r, 0.5)
	samples = s.Synthesize(d)
	_, ydata = data(samples, r)
	chart.AddSeries("Triangle", ydata)

	s = NewSawtoothWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	_, ydata = data(samples, r)
	chart.AddSeries("SawTooth", ydata)

	outpath := "output.TestWaveSynthesizers.html"
	outfile, _ := os.Create(outpath)
	defer outfile.Close()

	chart.Render(outfile)
}
