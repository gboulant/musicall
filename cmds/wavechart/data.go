package main

import "github.com/gboulant/musicall/wave"

func d01_KarplusStrong(samplerate int) []float64 {
	f := 5.
	a := 2.
	d := 4.
	return wave.KarplusStrongSignal(f, a, d, samplerate)
}

func d02_sweepfrequency(samplerate int, reverse bool) []float64 {
	a := 2.
	d := 10.   // sec.
	fmin := 1. // Hz
	fmax := 5. // Hz
	return wave.SweepFrequencySignal(fmin, fmax, a, d, reverse, samplerate)
}
