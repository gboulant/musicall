package main

import (
	"fmt"
	"log"

	"galuma.net/synthetic/sound"
	"galuma.net/synthetic/wave"
	"github.com/gopxl/beep"
)

var sampleRate = wave.DefaultSampleRate

const (
	fmin = 100. // minimal frequency of the range
	fmax = 350. // maximal frequency of the range
	cmin = 32   // minimal character ascii code (mapped to fmin)
	cmax = 126  // maximal character ascii code (mapped to fmax)
)

var c2f_slope = (fmax - fmin) / (cmax - cmin)
var c2f_offset = (cmax*fmin - cmin*fmax) / (cmax - cmin)

var synthetizer = wave.NewSquareWaveSynthesizer(0., 1., sampleRate)

func init() {
	err := sound.Init(sampleRate)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("c2f_slope  = %.2f\n", c2f_slope)
	log.Printf("c2f_offset = %.2f\n", c2f_offset)
}

func char2code(character rune) int {
	return int(character)
}

func code2freq(code int) float64 {
	return c2f_slope*float64(code) + c2f_offset
}

func string2streamer(phrase string) beep.Streamer {
	streamers := make([]beep.Streamer, len(phrase))
	duration := 0.2

	for i, char := range phrase {
		code := char2code(char)
		freq := code2freq(code)

		synthetizer.SetFrequency(freq)
		samples := synthetizer.Synthesize(duration)
		label := fmt.Sprintf("%c: f=%.1f Hz", char, freq)
		streamers[i] = sound.LabelledStreamer(sound.NewSound(samples), label)
	}
	return beep.Seq(streamers...)
}
