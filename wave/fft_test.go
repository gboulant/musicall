package wave

import (
	"testing"
)

func Test_fftfreq(t *testing.T) {
	N := 10
	r := 100
	T := 1. / float64(r)
	fftf := fftfreq(N, T)
	exp := []float64{0., 10, 20, 30, 40, -50, -40, -30, -20, -10}
	for i, v := range exp {
		if !almostEqual(fftf[i], v, 1e-6) {
			t.Errorf("value of item %d is %.6f (should be %.6f)", i, fftf[i], v)
		}
	}

	fftf = fftfreq(N+1, T)
	exp = []float64{0, 9.090909, 18.181818, 27.272727, 36.363636, 45.454545, -45.454545, -36.363636, -27.272727, -18.181818, -9.090909}
	for i, v := range exp {
		if !almostEqual(fftf[i], v, 1e-6) {
			t.Errorf("value of item %d is %.6f (should be %.6f)", i, fftf[i], v)
		}
	}
}

func Test_fft(t *testing.T) {
	N := 10
	samples := make([]float64, N)
	for i := range N {
		samples[i] = float64(2*i + 1)
	}

	r := 1
	T := 1. / float64(r)
	fftf := fftfreq(N, T)
	ffta := fft(samples)

	if len(fftf) != N {
		t.Errorf("the size of fftf is %d (should be %d)", len(fftf), N)
	}
	if len(ffta) != N {
		t.Errorf("the size of ffta is %d (should be %d)", len(ffta), N)
	}
}

func Test_spectrum(t *testing.T) {
	f := 100. // Hz
	r := DefaultSampleRate
	d := 2.
	a := 1.

	s := SineWaveSignal(f, a, d, r)
	frequencies, amplitudes := Spectrum(s, r)

	// On essaye de retrouver la fréquence ayant l'amplitude maximale
	// (et on s'attend à trouver une valeur proche de la valeur de f)
	imax := 0
	vmax := 0.
	for i, v := range amplitudes {
		if v > vmax {
			vmax = v
			imax = i
		}
	}
	maxamplfreq := frequencies[imax]
	if !almostEqual(maxamplfreq, f, 1.) {
		t.Errorf("Max Amplitude Frequency is %.2f (should be %.2f)", maxamplfreq, f)
	}
}
