package main

import (
	"log"
	"time"

	"galuma.net/synthetic/sound"
	"galuma.net/synthetic/wave"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

const sampleRate = beep.SampleRate(wave.DefaultSampleRate)

func init() {
	err := speaker.Init(sampleRate, sampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}
}

var d01_dissonance *Program = NewProgram("D01", "test de fréquences dissonantes", func() error {

	sinesound := func(f float64, a float64, d float64) beep.Streamer {
		synthesizer := wave.NewSineWaveSynthesizer(f, a, int(sampleRate))
		return sound.NewSound(d, synthesizer)
	}

	f := 440.
	a := 0.4
	d := 2.

	var streamers []beep.Streamer
	streamers = []beep.Streamer{
		sinesound(f, a, d),
		sinesound(3.*f/2., a, d),
	}

	var streamer beep.Streamer
	streamer = beep.Mix(streamers...)
	if err := sound.Play(streamer); err != nil {
		return err
	}

	streamers = []beep.Streamer{
		sinesound(f, a, d),
		sinesound(3.1*f/2., a, d),
	}

	streamer = beep.Mix(streamers...)
	if err := sound.Play(streamer); err != nil {
		return err
	}

	return nil
})
