package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gboulant/musicall/sound"
	"github.com/gboulant/musicall/wave"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/generators"
	"github.com/gopxl/beep/speaker"
)

const sampleRate = beep.SampleRate(wave.DefaultSampleRate)

func init() {
	// Le speaker est initialisé avec un sample rate fixé. Tous les
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
func decimate(samples []float64, samplerate int, step int) ([]float64, int) {
	return wave.Decimate(samples, step), samplerate / step
}

func labeledStreamer(f float64, duration float64, name string) beep.Streamer {
	s := wave.NewKarplusStrongSynthesizer(f, 1., int(sampleRate))
	samples := s.Synthesize(duration)
	label := fmt.Sprintf("note: %-4s (f=%.1f Hz)", name, f)
	return sound.LabelledStreamer(sound.NewSound(samples), label)
}

func DEMO00_logscale() error {
	duration := 1.0
	streamers := make([]beep.Streamer, 0)
	streamers = append(streamers, generators.Silence(int(duration*float64(sampleRate))))

	// En echelle logarithmique (log à base 2), les notes sont espacées
	// d'un intervalle de 1/12 (entre Do et Do# par exemple, et 2/12
	// entre Do et Ré, T = 2/12):
	//
	// Do -T-> Ré -T-> Mi -T/2-> Fa -T-> Sol -T-> La -T-> Si -T/2-> Do
	//

	fLa := 440.0 // Hz
	lLa := math.Log2(fLa)
	lDo := lLa - 9/12.
	notes := []struct {
		logf float64
		name string
	}{
		{logf: lDo, name: "Do"},
		{logf: lDo + 1/12., name: "Do#"},
		{logf: lDo + 2/12., name: "Ré"},
		{logf: lDo + 3/12., name: "Ré#"},
		{logf: lDo + 4/12., name: "Mi"},
		{logf: lDo + 5/12., name: "Fa"},
		{logf: lDo + 6/12., name: "Fa#"},
		{logf: lDo + 7/12., name: "Sol"},
		{logf: lDo + 8/12., name: "Sol#"},
		{logf: lDo + 9/12., name: "La"},
		{logf: lDo + 10/12., name: "La#"},
		{logf: lDo + 11/12., name: "Si"},
		{logf: lDo + 12/12., name: "Do"},
	}

	for _, note := range notes {
		f := math.Pow(2, note.logf)
		s := labeledStreamer(f, duration, note.name)
		streamers = append(streamers, s)
	}

	streamer := beep.Seq(streamers...)
	if err := sound.Play(streamer); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------
func DEMO01_quintes() error {

	f := 440.
	a := 0.8
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
		sinesound(3.03*f/2., a, d),
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
	a := 1.
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

	f := 440.
	a := 1.
	d := 3.
	r := int(sampleRate)

	mf := f * 0.1 // fréquence de la modulation d'amplitude
	ma := a * 0.2 // amplitude de la modulation d'amplitude

	size := int(d * float64(r))
	samples := make([]float64, size)
	angleIncrementFactor := math.Pi * 2 / float64(r)
	angleIncrement := angleIncrementFactor * f

	var angleModulation float64 = 0.
	var angleSample float64 = 0.
	var amplitude float64 = a
	for i := range samples {
		samples[i] = amplitude * math.Sin(angleSample)
		angleSample += angleIncrement
		angleModulation += angleIncrementFactor * mf
		amplitude = a + ma*math.Sin(angleModulation)
	}

	plts, pltr := decimate(samples, r, 10)
	label := "DEMO03_amplitude_modulation"
	outpath := fmt.Sprintf("output.%s.html", label)
	wave.PlotToFile(outpath, plts, pltr, label)

	if err := sound.Play(sound.NewSound(samples)); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------
func DEMO04_frequency_modulation() error {

	f := 440.
	a := 1.
	d := 3.
	r := int(sampleRate)

	mf := f * 0.1 // fréquence de la modulation de frequence
	ma := f * 0.4 // amplitude de la modulation de fréquence

	size := int(d * float64(r))
	samples := make([]float64, size)
	var angleInclementFactor float64 = math.Pi * 2 / float64(r)

	var angleModulation float64 = 0.
	var angleSample float64 = 0.
	var frequency float64 = f
	for i := range samples {
		samples[i] = a * math.Sin(angleSample)
		angleModulation += angleInclementFactor * mf
		frequency = f + ma*math.Sin(angleModulation)
		angleSample += angleInclementFactor * frequency
	}

	plts, pltr := decimate(samples, r, 5)
	label := "DEMO04_frequency_modulation"
	outpath := fmt.Sprintf("output.%s.html", label)
	wave.PlotToFile(outpath, plts, pltr, label)

	if err := sound.Play(sound.NewSound(samples)); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------

// DEMO05_sounds_like_a_laser emulates the sound of a laser saber
// starting. In fact it is the same implementation than the frequency
// modulation above, but with a buggy computation of the sinus angle.
// Indeed when computing the angle like a factor of products (using the
// step i as product) then you may see the angle phase decrease at some
// points because the frequency is rising down, even if i increases). It
// was a buggy version that let me discover a kind of laser sound.
// Chance.
func DEMO05_sounds_like_a_laser() error {

	f := 440.
	a := 1.
	d := 8.
	r := int(sampleRate)

	mf := f * 0.1 // fréquence de la modulation de frequence
	ma := f * 0.1 // amplitude de la modulation de fréquence

	// We observe:
	// - the more the ma is the longer the starting step is (and then
	// the period of the shuffle afterward).)

	size := int(d * float64(r))
	samples := make([]float64, size)
	var angleIncrementFactor float64 = math.Pi * 2 / float64(r)

	var angleModulation float64
	var angleSample float64
	var frequency float64
	for i := range samples {
		angleModulation = angleIncrementFactor * mf * float64(i)
		frequency = f + ma*math.Sin(angleModulation)
		angleSample = angleIncrementFactor * frequency * float64(i)
		samples[i] = a * math.Sin(angleSample)
	}

	plts, pltr := decimate(samples, r, 5)
	label := "DEMO05_sounds_like_a_laser"
	outpath := fmt.Sprintf("output.%s.html", label)
	wave.PlotToFile(outpath, plts, pltr, label)

	if err := sound.Play(sound.NewSound(samples)); err != nil {
		return err
	}

	return nil
}
