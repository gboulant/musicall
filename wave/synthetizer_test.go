package wave

import (
	"testing"
)

func TestNewSineWave(t *testing.T) {
	f := 440. // Hz
	a := 10.
	w := NewSineWaveSynthesizer(f, a, DefaultSampleRate)

	d := 2.0 // seconds
	s := w.Synthesize(d)

	explen := int(d * float64(DefaultSampleRate))
	reslen := len(s)
	if reslen != explen {
		t.Errorf("len is %d (should be %d)", reslen, explen)
	}
}
