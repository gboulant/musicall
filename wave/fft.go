package wave

// This file provides functions to compute the spectrum of a timeseries,
// i.e. the composition of the timeseries in terms of frequencies and
// their proportions (amplitudes). The API (and part of the
// implementation) is inspired from the Python scipy.fft package.
// examples from github.com/xigh/spectrogram.git

import (
	"math"
	"math/cmplx"
)

func hfft(samples []float64, freqs []complex128, n, step int) {
	if n == 1 {
		freqs[0] = complex(samples[0], 0)
		return
	}

	half := n / 2

	hfft(samples, freqs, half, 2*step)
	hfft(samples[step:], freqs[half:], half, 2*step)

	for k := range half {
		a := -2 * math.Pi * float64(k) / float64(n)
		e := cmplx.Rect(1, a) * freqs[k+half]

		freqs[k], freqs[k+half] = freqs[k]+e, freqs[k]-e
	}
}

// fft computes the FFT amplitudes (complex numbers) of the input signal
func fft(samples []float64) []complex128 {
	n := len(samples)
	freqs := make([]complex128, n)
	hfft(samples, freqs, n, 1)
	return freqs
}

// fftfreq computes the FFT frequencies
func fftfreq(N int, d float64) []float64 {
	// f = [0, 1, ...,   n/2-1,     -n/2, ..., -1] / (d*n)   if n is even
	// f = [0, 1, ..., (n-1)/2, -(n-1)/2, ..., -1] / (d*n)   if n is odd
	f := make([]float64, N)
	dxN := d * float64(N)
	var limit int
	if N%2 == 0 { // N is even
		limit = N/2 - 1
	} else { // N is odd
		limit = (N - 1) / 2
	}
	for i := 0; i <= limit; i++ {
		f[i] = float64(i) / dxN
	}
	for i := limit + 1; i < N; i++ {
		f[i] = -float64(N-i) / dxN
	}
	return f
}

func Spectrum(samples []float64, samplerate int) (frequencies, amplitudes []float64) {
	N := len(samples)
	T := 1. / float64(samplerate)
	fftf := fftfreq(N, T)
	ffta := fft(samples)

	// Pour les signaux à valeurs réelles, le spectre est symétrique et
	// on ne garde que les valeurs de fréquences positives, c'est-à-dire
	// jusque l'indice N/2 - 1 si N est pair, et (N-1)/2 si N est impair.
	var lastidx int
	if N%2 == 0 { // N is even
		lastidx = N/2 - 1
	} else { // N is odd
		lastidx = (N - 1) / 2
	}

	// frequencies = fftf[:N//2-1]
	// amplitudes  = 2.0/N * modulus(ffta[:N//2-1])
	frequencies = fftf[:lastidx+1]
	amplitudes = make([]float64, lastidx+1)
	for i := 0; i <= lastidx; i++ {
		amplitudes[i] = 2.0 * cmplx.Abs(ffta[i]) / float64(N)
	}
	return frequencies, amplitudes
}
