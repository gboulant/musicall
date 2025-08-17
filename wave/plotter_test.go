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
	if err := PlotToFile(outfilepath, pltsamples, pltsamplerate); err != nil {
		t.Error(err)
	}

	samples = d02_sweepfrequency(samplerate, false)
	outfilepath = "output.d02_sweepfrequency.html"
	pltsamples, pltsamplerate = decimate(samples, samplerate, decimatestep)
	if err := PlotToFile(outfilepath, pltsamples, pltsamplerate); err != nil {
		t.Error(err)
	}
}
