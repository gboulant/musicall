package main

import (
	"log"
	"math"
	"time"

	"galuma.net/synthetic/sound"
	"galuma.net/synthetic/wave"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/generators"
	"github.com/gopxl/beep/speaker"
)

const sampleRate = beep.SampleRate(wave.DefaultSampleRate)

func init() {
	// Le speaker est initialisé avec un sample rate donnée. Tous les
	// signaux ([]float64) joués par ce speaker seront considérés comme
	// des sons avec ce sample rate. On doit donc générer des signaux
	// avec ce sample rate.
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
func DEMO01_quintes() error {

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
}

// ----------------------------------------------------------------
func DEMO02_vibrato() error {
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
		sound.NewSound(vibrato),
	)
	if err := sound.Play(streamer); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------
func DEMO03_amplitude_modulation() error {

	desimate := false

	f := 440.
	a := 1.
	d := 3.
	r := int(sampleRate)

	if desimate {
		f = f / 10.
		r = int(sampleRate / 100)
	}

	mf := f * 0.1 // fréquence de la modulation d'amplitude
	ma := a * 0.2 // amplitude de la modulation d'amplitude

	size := int(d * float64(r))
	samples := make([]float64, size)
	var angle float64 = math.Pi * 2 / float64(r)

	var angleModulation float64 = 0.
	var angleSample float64 = 0.
	var amplitude float64 = a
	for i := range samples {
		samples[i] = amplitude * math.Sin(angleSample)
		angleSample += angle * f
		angleModulation += angle * mf
		amplitude = a + ma*math.Sin(angleModulation)
	}

	wave.PlotToFile("output.DEMO03_amplitude_modulation.html", samples, r)

	if desimate {
		// Do not play the sound, the sample rate is not consistent with
		// the speaker
		return nil
	}

	if err := sound.Play(sound.NewSound(samples)); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------
func DEMO04_frequency_modulation() error {
	desimate := false

	f := 440.
	a := 1.
	d := 3.
	r := int(sampleRate)

	if desimate {
		r = int(sampleRate / 100)
	}

	mf := f * 0.1 // fréquence de la modulation de frequence
	ma := f * 0.2 // amplitude de la modulation de fréquence

	size := int(d * float64(r))
	samples := make([]float64, size)
	var angle float64 = math.Pi * 2 / float64(r)

	var angleModulation float64 = 0.
	var angleSample float64 = 0.
	var frequency float64 = f
	for i := range samples {
		samples[i] = a * math.Sin(angleSample)
		angleModulation += angle * mf
		frequency = f + ma*math.Sin(angleModulation)
		angleSample += angle * frequency
	}

	wave.PlotToFile("output.DEMO04_frequency_modulation.html", samples, int(r))

	if desimate {
		// Do not play the sound, the sample rate is not consistent with
		// the speaker
		return nil
	}
	if err := sound.Play(sound.NewSound(samples)); err != nil {
		return err
	}

	return nil
}
