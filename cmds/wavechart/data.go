package main

import (
	"galuma.net/synthetic/wave"
)

func d01_KarplusStrong() (samples []float64, samplerate int) {
	f := 5.
	a := 2.
	d := 4.
	samples = wave.KarplusStrongSignal(f, a, d)
	samplerate = wave.DefaultSampleRate
	return samples, samplerate
}

func d02_sweepfrequency(reverse bool) (samples []float64, samplerate int) {
	a := 2.
	d := 10.   // sec.
	fmin := 1. // Hz
	fmax := 5. // Hz

	w := wave.NewSweepFrequencySynthesizer(fmin, fmax, a, int(wave.DefaultSampleRate))
	samples = w.Synthesize(d)
	samplerate = w.SampleRate()

	// Uncomment to reverse the signal (sweep from max to min)
	if reverse {
		wave.Reverse(&samples)
	}
	return samples, samplerate
}
