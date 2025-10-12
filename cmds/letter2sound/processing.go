package main

import (
	"fmt"
	"log"

	"github.com/gboulant/musicall/music"
	"github.com/gboulant/musicall/sound"
	"github.com/gboulant/musicall/wave"
	"github.com/gopxl/beep"
)

var sampleRate = wave.DefaultSampleRate

func init() {
	err := sound.Init(sampleRate)
	if err != nil {
		log.Fatal(err)
	}
}

// -------------------------------------------------------------------
// Model with a frequency value linear with the ascii code value

const (
	fmin = 100. // minimal frequency of the range
	fmax = 350. // maximal frequency of the range
	cmin = 32   // minimal character ascii code (mapped to fmin)
	cmax = 126  // maximal character ascii code (mapped to fmax)
)

var c2f_slope = (fmax - fmin) / (cmax - cmin)
var c2f_offset = (cmax*fmin - cmin*fmax) / (cmax - cmin)

func code2freq_linear(code int) float64 {
	return c2f_slope*float64(code) + c2f_offset
}

// -------------------------------------------------------------------
// Model with a frequency value equal to a music note
var Do0 music.Note = music.Note{Octave: 0, Index: 0}

func code2freq_interval(code int) float64 {
	interval := music.Interval(code - spacecode)
	log.Printf("code=%3d interval=%d", code, interval)
	n := Do0.Derived(interval)
	return n.Frequency()
}

// -------------------------------------------------------------------
// var synthetizer = wave.NewSquareWaveSynthesizer(0., 1., sampleRate)
var synthetizer = wave.NewSineWaveSynthesizer(0., 1., sampleRate)

//var synthetizer = wave.NewKarplusStrongSynthesizer(0., 1., sampleRate)

func letter2sound(letter rune, duration float64) beep.Streamer {
	code2freq := code2freq_linear
	//code2freq := code2freq_interval

	code := char2code(letter)
	if code == spacecode {
		label := " : silence"
		return sound.LabelledStreamer(sound.Silence(duration, sampleRate), label)
	}
	freq := code2freq(code)
	synthetizer.SetFrequency(freq)
	samples := synthetizer.Synthesize(duration)
	wave.SmoothBoundaries(&samples, sampleRate, 0.1*duration)
	label := fmt.Sprintf("%c: c=%3d f=%.1f Hz", letter, code, freq)
	return sound.LabelledStreamer(sound.NewSound(samples), label)
}

func string2streamer(phrase string) beep.Streamer {
	streamers := make([]beep.Streamer, len(phrase))
	duration := 0.2

	for i, char := range phrase {
		streamers[i] = letter2sound(char, duration)
	}
	return beep.Seq(streamers...)
}
