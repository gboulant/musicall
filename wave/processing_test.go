package wave

import (
	"log"
	"testing"
)

func TestAddNoise(t *testing.T) {
	f := 120.
	a := 1.
	d := 4.
	samples := SquareWaveSignal(f, a, d)
	log.Println(samples[:10])
	AddNoise(&samples, 0.2)
	log.Println((samples)[:10])

}
