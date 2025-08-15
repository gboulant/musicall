package wave

import (
	"testing"
)

func d01_KarplusStrong(r int) []float64 {
	f := 5.
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

func TestPlotToFile(t *testing.T) {
	samplerate := DefaultSampleRate / 1000

	samples := d01_KarplusStrong(samplerate)
	outfilepath := "output.d01_KarplusStrong.html"
	if err := PlotToFile(outfilepath, samples, samplerate); err != nil {
		t.Error(err)
	}

	samples = d02_sweepfrequency(samplerate, false)
	outfilepath = "output.d02_sweepfrequency.html"
	if err := PlotToFile(outfilepath, samples, samplerate); err != nil {
		t.Error(err)
	}
}
