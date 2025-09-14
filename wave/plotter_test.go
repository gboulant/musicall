package wave

import (
	"fmt"
	"testing"
)

func d01_KarplusStrong(r int) []float64 {
	f := 440.
	a := 2.
	d := 4.
	return KarplusStrongSignal(f, a, d, r)
}

func d02_sweepfrequency(r int) []float64 {
	a := 2.
	d := 10.   // sec.
	fmin := 1. // Hz
	fmax := 5. // Hz
	return SweepFrequencySignal(fmin, fmax, a, d, r)
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

	samples = d02_sweepfrequency(samplerate)
	outfilepath = "output.d02_sweepfrequency.html"
	pltsamples, pltsamplerate = decimate(samples, samplerate, decimatestep)
	if err := PlotToFile(outfilepath, pltsamples, pltsamplerate, "sweepfrequency"); err != nil {
		t.Error(err)
	}
}

func TestWavePlotter(t *testing.T) {
	f := 3.
	a := 2.
	d := 4.
	r := int(100 * f) // 100 points by cycle

	p := NewPlotter()
	//p.SetTitle("Different Wave Forms")

	var s Synthesizer
	s = NewSineWaveSynthesizer(f, a, r)
	samples := s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "Sine")

	s = NewSquareWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "Square")

	s = NewKarplusStrongSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "KarplusStrong")

	s = NewTriangleWaveSynthesizer(f, a, r, 0.7)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "Triangle")

	s = NewSawtoothWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "SawTooth")

	s = NewPWMWaveSynthesizer(f, a, r, 0.2)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "PulseWidthModulation")

	s = NewSweepFrequencySynthesizer(f, 2*f, a, r)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "SweepFrequencySynthesizer")

	outpath := "output.TestWavePlotter.html"
	p.Save(outpath)
}

func TestTriangleWave(t *testing.T) {
	f := 1.
	a := 1.
	d := 1.
	n := 1000                // number of point in one f cycle
	r := int(float64(n) * f) // sample rate

	p := NewPlotter()

	var s Synthesizer
	var samples []float64

	s = NewSquareWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "Square")

	var risingrate float64
	var label string
	for i := range 10 {
		risingrate = float64(i) / 10
		s = NewTriangleWaveSynthesizer(f, a, r, risingrate)
		samples = s.Synthesize(d)
		label = fmt.Sprintf("Triangle %.2f", risingrate)
		p.AddLineSampledValues(samples, r, label)
	}

	/*
		s = NewSawtoothWaveSynthesizer(f, a, r)
		samples = s.Synthesize(d)
		p.AddLineSampledValues(samples, r, "SawTooth")

		s = NewRegularTriangleWaveSynthesizer(f, a, r)
		samples = s.Synthesize(d)
		p.AddLineSampledValues(samples, r, "Regular")
	*/
	outpath := "output.TestTriangleWave.html"
	p.Save(outpath)
}
