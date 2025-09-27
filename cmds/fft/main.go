package main

import (
	"fmt"

	"github.com/gboulant/musicall/sound"
	"github.com/gboulant/musicall/wave"
)

func main() {
	f := 100. // Hz
	r := wave.DefaultSampleRate
	d := 2.
	a := 1.

	s1 := wave.SineWaveSignal(f, a, d, r)
	s2 := wave.SineWaveSignal(4*f, 0.3*a, d, r)
	s := make([]float64, len(s1))
	for i := range len(s1) {
		s[i] = s1[i] + s2[i]
	}
	wave.Normalize(&s)

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

	// Plot of the spectrum
	psp := wave.NewPlotter()
	psp.AddLineXYValues(frequencies, amplitudes, "Spectre")
	psp.SetXFormatter("{value}Hz")
	htmlpath := "output.fft.html"
	psp.Save(htmlpath)

	// Plot of the time series
	pts := wave.NewPlotter()
	pts.AddLineSampledValues(s, r, "Signal")
	htmlpath = "output.signal.html"
	pts.Save(htmlpath)

	// Play the sound of the timeseries
	sound.Init(r)
	streamer := sound.NewSound(s)
	sound.Play(streamer)
}
