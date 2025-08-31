package main

import (
	"fmt"
	"log"

	"github.com/gboulant/musicall/sound"
	"github.com/gboulant/musicall/wave"
)

const sampleRate = wave.DefaultSampleRate

func init() {
	err := sound.Init(sampleRate)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	f := 404.
	a := 1.
	s := wave.NewSquareWaveSynthesizer(f, a, sampleRate)

	d := 4.
	samples := s.Synthesize(d)
	label := fmt.Sprintf("f=%.1f Hz", s.Frequency())
	streamer := sound.LabelledStreamer(sound.NewSound(samples), label)

	if err := sound.Play(streamer); err != nil {
		log.Fatal(err)
	}

}
