package main

import (
	"fmt"

	"github.com/gboulant/musicall/sound"
	"github.com/gboulant/musicall/wave"
)

func testsignal(f, a, d float64, r int) []float64 {
	s1 := wave.SineWaveSignal(f, a, d, r)
	s2 := wave.SineWaveSignal(4*f, 0.3*a, d, r)
	s := make([]float64, len(s1))
	for i := range len(s1) {
		s[i] = s1[i] + s2[i]
	}
	wave.Normalize(&s)
	return s
}

func D01_fft() error {
	f := 100. // Hz
	r := wave.DefaultSampleRate
	d := 2.
	a := 1.
	s := testsignal(f, a, d, r)

	// Plot of the time series
	pts := wave.NewPlotter()
	pts.AddLineSampledValues(s, r, "Signal")
	htmlpath := "output.signal.html"
	pts.Save(htmlpath)

	// Computing of the spectrum
	frequencies, amplitudes := wave.Spectrum(s, r)

	// Plot of the spectrum
	psp := wave.NewPlotter()
	psp.AddLineXYValues(frequencies, amplitudes, "Spectre")
	psp.SetXFormatter("{value}Hz")
	htmlpath = "output.fft.html"
	psp.Save(htmlpath)

	// Play the sound of the timeseries
	sound.Init(r)
	streamer := sound.NewSound(s)
	return sound.Play(streamer)
}

func D02_fft_frequencyOfMaxAmplitude() error {
	f := 100. // Hz
	r := wave.DefaultSampleRate
	d := 2.
	a := 1.
	s := testsignal(f, a, d, r)
	frequencies, amplitudes := wave.Spectrum(s, r)

	// On essaye de retrouver la fréquence ayant l'amplitude maximale
	imax := 0
	vmax := 0.
	for i, v := range amplitudes {
		if v > vmax {
			vmax = v
			imax = i
		}
	}
	fmt.Printf("Frequency of Maximal Amplitude = %.2f\n", frequencies[imax])
	return nil
}

func D03_fft_smoothboundaries() error {
	f := 100. // Hz
	r := wave.DefaultSampleRate
	d := 2.
	a := 1.
	s := testsignal(f, a, d, r)

	// --------------------------------------------------
	// Analysis of the signal without smoothing the boundaries
	frequencies, amplitudes := wave.Spectrum(s, r)

	// Plot the signal and spectrum
	pts := wave.NewPlotter()
	pts.AddLineSampledValues(s, r, "Signal")
	psp := wave.NewPlotter()
	psp.AddLineXYValues(frequencies, amplitudes, "Spectre")
	psp.SetXFormatter("{value}Hz")

	// Play the sound of the signal
	sound.Init(r)
	streamer := sound.NewSound(s)
	sound.Play(streamer)

	// --------------------------------------------------
	// Analysis of the signal with smoothing the boundaries
	wave.SmoothBoundaries(&s, r, 10/f)
	frequencies, amplitudes = wave.Spectrum(s, r)

	// Plot the signal and spectrum
	pts.AddLineSampledValues(s, r, "Signal Smooth")
	psp.AddLineXYValues(frequencies, amplitudes, "Spectre Smooth")

	// Play the sound of the signal
	streamer = sound.NewSound(s)
	sound.Play(streamer)

	// --------------------------------------------------
	// Save the plots
	htmlpath := "output.signal.html"
	pts.Save(htmlpath)
	htmlpath = "output.fft.html"
	psp.Save(htmlpath)

	return nil
}
