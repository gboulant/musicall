package main

import (
	"log"
	"time"

	"galuma.net/synthetic/sound"
	"galuma.net/synthetic/wave"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/generators"
	"github.com/gopxl/beep/speaker"
)

const sampleRate = beep.SampleRate(wave.DefaultSampleRate)

func init() {
	err := speaker.Init(sampleRate, sampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}
}

func sinesound(f float64, a float64, d float64) beep.Streamer {
	synthesizer := wave.NewSineWaveSynthesizer(f, a, int(sampleRate))
	return sound.SynthSound(d, synthesizer)
}
func silence(duration float64) beep.Streamer {
	return generators.Silence(int(duration * float64(sampleRate)))
}

// ----------------------------------------------------------------
var _ *Program = NewProgram("D01", "son de quintes", func() error {

	f := 440.
	a := 0.4
	d := 2.

	// quinte juste (rapport 3/2)
	var streamers []beep.Streamer
	streamers = []beep.Streamer{
		sinesound(f, a, d),
		sinesound(3.*f/2., a, d),
	}

	var streamer beep.Streamer
	streamer = beep.Seq(
		silence(0.5),
		beep.Mix(streamers...),
	)
	if err := sound.Play(streamer); err != nil {
		return err
	}

	// quinte non juste
	streamers = []beep.Streamer{
		sinesound(f, a, d),
		sinesound(3.01*f/2., a, d),
	}

	streamer = beep.Seq(
		silence(0.5),
		beep.Mix(streamers...),
	)
	if err := sound.Play(streamer); err != nil {
		return err
	}

	return nil
})

// ----------------------------------------------------------------
var _ *Program = NewProgram("D02", "vibrato", func() error {
	f := 80.
	a := 2.
	d := 3.
	deltaf := 3.

	// 1. Superposition of two waves using the Mix of streamers
	streamers := []beep.Streamer{
		sinesound(f, a, d),
		sinesound(f+deltaf, a, d),
	}

	streamer := beep.Seq(
		silence(0.5),
		beep.Mix(streamers...),
	)
	if err := sound.Play(streamer); err != nil {
		return err
	}

	// 2. Superposition of two waves using the addition of signals
	signal1 := wave.NewSineWaveSynthesizer(f, a, int(sampleRate)).Synthesize(d)
	signal2 := wave.NewSineWaveSynthesizer(f+deltaf, a, int(sampleRate)).Synthesize(d)
	vibrato := make([]float64, len(signal1))
	for i := range vibrato {
		vibrato[i] = signal1[i] + signal2[i]
	}
	streamer = beep.Seq(
		silence(0.5),
		&sound.Sound{TotalSamples: vibrato, Processed: 0},
	)
	if err := sound.Play(streamer); err != nil {
		return err
	}

	return nil
})
