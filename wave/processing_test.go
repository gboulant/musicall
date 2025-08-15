package wave

import (
	"log"
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
