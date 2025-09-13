package wave

import (
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
	return SweepFrequencySignal(fmin, fmax, a, d, reverse, r)
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

func TestWavePlotter(t *testing.T) {
	f := 10.
	a := 2.
	d := 4.
	r := int(100 * f) // 100 points by cycle

	p := NewPlotter()
	p.SetTitle("Different Wave Forms")

	s := NewSineWaveSynthesizer(f, a, r)
	samples := s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "Sine")

	s = NewSquareWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "Square")

	s = NewKarplusStrongSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "KarplusStrong")

	// ERROR: problème de pente avec le triangle
	s = NewTriangleWaveSynthesizer(f, a, r, 0.5)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "Triangle")

	s = NewSawtoothWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "SawTooth")

	s = NewPWMWaveSynthesizer(f, a, r, 0.2)
	samples = s.Synthesize(d)
	p.AddLineSampledValues(samples, r, "PulseWidthModulation")

	outpath := "output.TestWavePlotter.html"
	p.Save(outpath)
}
