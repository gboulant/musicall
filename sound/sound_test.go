package sound

import (
	"log"
	"testing"
	"time"

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

func silence(duration float64) beep.Streamer {
	return generators.Silence(int(duration * float64(sampleRate)))
}

func TestNewSound(t *testing.T) {
	a := 1.

	sinesound := func(f float64, d float64) beep.Streamer {
		synthesizer := wave.NewSineWaveSynthesizer(f, a, int(sampleRate))
		return SynthSound(d, synthesizer)
	}
	squaresound := func(f float64, d float64) beep.Streamer {
		synthesizer := wave.NewSquareWaveSynthesizer(f, a, int(sampleRate))
		return SynthSound(d, synthesizer)
	}
	guitarsound := func(f float64, d float64) beep.Streamer {
		synthesizer := wave.NewKarplusStrongSynthesizer(f, a, int(sampleRate))
		return SynthSound(d, synthesizer)
	}

	f := 440.
	d := 1.

	streamers := []beep.Streamer{
		silence(0.2), sinesound(f, d),
		silence(0.2), squaresound(f, d),
		silence(0.2), guitarsound(f, d),
	}

	streamer := beep.Seq(streamers...)
	if err := Play(streamer); err != nil {
		t.Error(err)
	}
}

func TestSoundStruct(t *testing.T) {
	// The sound can be created directly from the raw signal without a
	// predefined syntesizer (even if in this example we use one for
	// initializing the signal array)
	f := 440.
	a := 1.
	d := 1.

	sound1 := NewSound(wave.SineWaveSignal(f, a, d, 0))
	sound2 := NewSound(wave.SquareWaveSignal(f, a, d, 0))
	sound3 := NewSound(wave.KarplusStrongSignal(f, a, d, 0))

	streamers := []beep.Streamer{
		silence(0.2), sound1,
		silence(0.2), sound2,
		silence(0.2), sound3,
	}

	streamer := beep.Seq(streamers...)
	if err := Play(streamer); err != nil {
		t.Error(err)
	}
}

func TestSoundWithNoise(t *testing.T) {
	f := 440.
	a := 1.
	d := 2.

	signal0 := wave.SineWaveSignal(f, a, d, 0)
	signal1 := make([]float64, len(signal0))
	copy(signal1, signal0)
	wave.AddNoise(&signal1, 0.2*a)

	streamers := []beep.Streamer{
		silence(0.2), NewSound(signal0),
		silence(0.2), NewSound(signal1),
	}

	streamer := beep.Seq(streamers...)
	if err := Play(streamer); err != nil {
		t.Error(err)
	}
}

func TestSweepFrequency01(t *testing.T) {
	f0 := 60. // Hz
	f1 := 280.
	a := 5.
	d := 2.0 // seconds

	var w wave.Synthesizer
	var s []float64

	w = wave.NewSweepFrequencySynthesizer(f0, f1, a, int(sampleRate))
	s = w.Synthesize(d)
	Play(NewSound(s))

	wave.Reverse(&s)
	Play(NewSound(s))
}

func TestSweepFrequency02(t *testing.T) {
	// Exemple 2, using streamers

	f0 := 60. // Hz
	f1 := 280.
	a := 0.5
	d := 2. // seconds

	var w wave.Synthesizer
	var s []float64

	w = wave.NewSweepFrequencySynthesizer(f0, f1, a, int(sampleRate))

	var streamers []beep.Streamer
	streamers = append(streamers, silence(0.4))

	s = w.Synthesize(d)
	streamers = append(streamers, NewSound(s))

	s = w.Synthesize(d)
	wave.Reverse(&s)
	streamers = append(streamers, NewSound(s))

	streamer := beep.Seq(streamers...)
	if err := Play(streamer); err != nil {
		t.Error(err)
	}
}
