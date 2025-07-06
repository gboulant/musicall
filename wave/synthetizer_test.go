package wave

import (
	"log"
	"testing"
)

func TestNewSineWave(t *testing.T) {
	f := 440. // Hz
	a := 10.
	w := NewSineWave(f, a, DefaultSampleRate)
	d := 2.0 // seconds
	s := w.Synthesize(d)
	log.Println(s)
}
