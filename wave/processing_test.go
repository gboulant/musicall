package wave

import (
	"fmt"
	"log"
	"math"
	"testing"
)

func TestAddNoise(t *testing.T) {
	f := 120.
	a := 1.
	d := 4.
	r := DefaultSampleRate
	samples := SquareWaveSignal(f, a, d, r)
	log.Println(samples[:10])
	AddNoise(&samples, 0.2)
	log.Println((samples)[:10])
}

const float64EqualityThreshold = 1e-6

func almostEqual(a, b float64, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}

func TestDecimate(t *testing.T) {
	f := 120.
	a := 1.
	d := 4.
	r := DefaultSampleRate
	s := SineWaveSignal(f, a, d, r)

	// Note that the number of sample points by cycle is Nc = r*T = r/f.
	// After decimation, the number of points by cycle is reduced with a
	// factor step: Nc = r/(f*step). Then be carrefull when choosing the
	// decimation step, and take care of the cycle frequencies.
	//
	// For example, if the frequency is f=440 (La) and the sample rate
	// is r=44100~44000, then the original Nc is Nc=r/f=44000/440=100.
	// If you choose a decimation step around 100, you will have one
	// point by cycle.

	step := 10
	decimate := Decimate(s, step)

	inputsize := int(d * float64(r))
	outputsize := int(float64(inputsize) / float64(step))

	reslen := len(decimate)
	explen := outputsize
	if reslen != explen {
		t.Errorf("len is %d (should be %d)", reslen, explen)
	}

	// Create a signal with sample rate reduced of step. The create
	// signal should be the same than the decimate signal.
	r = r / step
	s = SineWaveSignal(f, a, d, r)

	if len(s) != len(decimate) {
		t.Errorf("len is %d (should be %d)", len(s), len(decimate))
	}
	for i, v := range decimate {
		if !almostEqual(v, s[i], float64EqualityThreshold) {
			t.Errorf("s[%d] is %.8f (should be %.8f)", i, s[i], v)
		}
	}
}

func TestMinMax(t *testing.T) {
	f := 120.
	a := 1.4
	d := 4.
	r := DefaultSampleRate
	samples := SineWaveSignal(f, a, d, r)

	// Add an offset for non trivial min max
	offset := 2.3
	for i := range samples {
		samples[i] = samples[i] + offset
	}

	resmin, resmax, resmed := MinMax(&samples)
	expmin := -a + offset
	expmax := +a + offset
	expmed := offset

	if !almostEqual(resmin, expmin, 1e-3) {
		t.Errorf("min is %.2f (should be %.2f)", resmin, expmin)
	}
	if !almostEqual(resmax, expmax, 1e-3) {
		t.Errorf("max is %.2f (should be %.2f)", resmax, expmax)
	}
	if !almostEqual(resmed, expmed, 1e-3) {
		t.Errorf("med is %.2f (should be %.2f)", resmed, expmed)
	}
}

func TestRescale(t *testing.T) {
	f := 5.
	a := 1.4
	d := 4.
	r := int(f * 100) // 100 point by cycle
	p := NewPlotter()

	samples := SineWaveSignal(f, a, d, r)

	// Add an offset for non trivial min max
	offset := 2.3
	for i := range samples {
		samples[i] = samples[i] + offset
	}
	p.AddLineSampledValues(samples, r, "origin")

	// Rescale to the range -1, +1
	min, max, _ := MinMax(&samples)
	Rescale(&samples, min, max, -1., +1.)
	p.AddLineSampledValues(samples, r, "rescale")
	outpath := "output.TestRescale.html"
	p.Save(outpath)

	resmin, resmax, _ := MinMax(&samples)
	expmin := -1.
	expmax := +1.
	if !almostEqual(resmin, expmin, 1e-3) {
		t.Errorf("min is %.2f (should be %.2f)", resmin, expmin)
	}
	if !almostEqual(resmax, expmax, 1e-3) {
		t.Errorf("max is %.2f (should be %.2f)", resmax, expmax)
	}
}

func TestFFT(t *testing.T) {
	a := []complex128{1, 1, 1, 1, 0, 0, 0, 0}
	fmt.Println(fft(a, len(a), 1))
}
