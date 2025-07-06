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
		synthesizer := wave.NewSineWave(f, a, int(sampleRate))
		return NewSound(d, synthesizer)
	}
	squaresound := func(f float64, d float64) beep.Streamer {
		synthesizer := wave.NewSquareWave(f, a, int(sampleRate))
		return NewSound(d, synthesizer)
	}
	guitarsound := func(f float64, d float64) beep.Streamer {
		synthesizer := wave.NewKarplusStrongWave(f, a, int(sampleRate))
		return NewSound(d, synthesizer)
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

	signal := wave.SineWaveSignal(f, a, d)
	sound1 := &Sound{signal, 0}
	signal = wave.SquareWaveSignal(f, a, d)
	sound2 := &Sound{signal, 0}
	signal = wave.KarplusStrongWaveSignal(f, a, d)
	sound3 := &Sound{signal, 0}

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
