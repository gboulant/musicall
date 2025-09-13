package wave

import (
	"testing"
)

func TestWaveSignal(t *testing.T) {
	n := 40                  // nb sample points in a period T
	f := 240.                // frequency of the sound to play
	r := int(float64(n) * f) // samplerate of the signal to create
	d := 2.0                 // duration in sec
	expsize := int(d * float64(r))
	var ressize int

	var samples []float64

	samples = SineWaveSignal(f, 1., d, r)
	ressize = len(samples)
	if ressize != expsize {
		t.Errorf("SineWaveSignal: size is %d (should be %d)", ressize, expsize)
	}

	samples = SquareWaveSignal(f, 1., d, r)
	ressize = len(samples)
	if ressize != expsize {
		t.Errorf("SquareWaveSignal: size is %d (should be %d)", ressize, expsize)
	}

	samples = KarplusStrongSignal(f, 1., d, r)
	ressize = len(samples)
	if ressize != expsize {
		t.Errorf("KarplusStrongSignal: size is %d (should be %d)", ressize, expsize)
	}

	samples = SweepFrequencySignal(f, 2*f, 1., d, r)
	ressize = len(samples)
	if ressize != expsize {
		t.Errorf("SweepFrequencySignal: size is %d (should be %d)", ressize, expsize)
	}

	samples = PWMWaveSignal(f, 1., d, 0.2, r)
	ressize = len(samples)
	if ressize != expsize {
		t.Errorf("PWMWaveSignal: size is %d (should be %d)", ressize, expsize)
	}

	samples = TriangleWaveSignal(f, 1., d, 0.5, r)
	ressize = len(samples)
	if ressize != expsize {
		t.Errorf("TriangleWaveSignal: size is %d (should be %d)", ressize, expsize)
	}

	samples = SawToothWaveSignal(f, 1., d, r)
	ressize = len(samples)
	if ressize != expsize {
		t.Errorf("SawToothWaveSignal: size is %d (should be %d)", ressize, expsize)
	}

}
