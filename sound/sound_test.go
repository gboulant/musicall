package sound

import (
	"log"
	"testing"
	"time"

	"github.com/gboulant/musicall/wave"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/generators"
	"github.com/gopxl/beep/speaker"
)

const testSampleRate = wave.DefaultSampleRate

func init() {
	// Le speaker est initialisé avec un sample rate fixé. Tous les
	// signaux ([]float64) joués par ce speaker seront considérés comme
	// des sons avec ce sample rate. On doit donc générer des signaux
	// avec ce sample rate.
	if err := Init(testSampleRate); err != nil {
		log.Fatal(err)
	}
}

func silence(duration float64) beep.Streamer {
	return generators.Silence(int(duration * float64(testSampleRate)))
}

func TestNewSound(t *testing.T) {
	a := 1. // DO NOT set a>1, it will be trucated by the beep speaker

	sinesound := func(f float64, d float64) beep.Streamer {
		synthesizer := wave.NewSineWaveSynthesizer(f, a, testSampleRate)
		return SynthSound(d, synthesizer)
	}
	squaresound := func(f float64, d float64) beep.Streamer {
		synthesizer := wave.NewSquareWaveSynthesizer(f, a, testSampleRate)
		return SynthSound(d, synthesizer)
	}
	guitarsound := func(f float64, d float64) beep.Streamer {
		synthesizer := wave.NewKarplusStrongSynthesizer(f, a, testSampleRate)
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

func TestWaveSynthesisers(t *testing.T) {
	f := 440.
	a := 1.
	d := 1.
	r := wave.DefaultSampleRate

	pause := 0.2
	var streamers []beep.Streamer

	var s wave.Synthesizer

	s = wave.NewSineWaveSynthesizer(f, a, r)
	samples := s.Synthesize(d)
	stream := LabelledStreamer(NewSound(samples), "Sine Wave")
	streamers = append(streamers, silence(pause))
	streamers = append(streamers, stream)

	s = wave.NewSquareWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	stream = LabelledStreamer(NewSound(samples), "Square Wave")
	streamers = append(streamers, silence(pause))
	streamers = append(streamers, stream)

	s = wave.NewPWMWaveSynthesizer(f, a, r, 0.1)
	samples = s.Synthesize(d)
	stream = LabelledStreamer(NewSound(samples), "PWM Wave")
	streamers = append(streamers, silence(pause))
	streamers = append(streamers, stream)

	s = wave.NewRegularTriangleWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	stream = LabelledStreamer(NewSound(samples), "Regular Triangle Wave")
	streamers = append(streamers, silence(pause))
	streamers = append(streamers, stream)

	s = wave.NewSawtoothWaveSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	stream = LabelledStreamer(NewSound(samples), "Saw Tooth Wave")
	streamers = append(streamers, silence(pause))
	streamers = append(streamers, stream)

	s = wave.NewKarplusStrongSynthesizer(f, a, r)
	samples = s.Synthesize(d)
	stream = LabelledStreamer(NewSound(samples), "Karplus Strong Wave")
	streamers = append(streamers, silence(pause))
	streamers = append(streamers, stream)

	s = wave.NewSweepFrequencySynthesizer(f, 2*f, a, r)
	samples = s.Synthesize(d)
	stream = LabelledStreamer(NewSound(samples), "Sweep Frequency Wave")
	streamers = append(streamers, silence(pause))
	streamers = append(streamers, stream)

	streamer := beep.Seq(streamers...)
	if err := Play(streamer); err != nil {
		t.Error(err)
	}
}

func TestSweepFrequency01(t *testing.T) {
	f0 := 60. // Hz
	f1 := 280.
	a := 1.
	d := 2.0 // seconds

	var w wave.Synthesizer
	var s []float64

	w = wave.NewSweepFrequencySynthesizer(f0, f1, a, testSampleRate)
	s = w.Synthesize(d)
	Play(NewSound(s))

	w = wave.NewSweepFrequencySynthesizer(f1, f0, a, testSampleRate)
	s = w.Synthesize(d)
	Play(NewSound(s))
}

func TestSweepFrequency02(t *testing.T) {
	// Exemple 2, using streamers

	f0 := 60. // Hz
	f1 := 280.
	a := 1.
	d := 2. // seconds

	var w wave.Synthesizer
	var s []float64

	w = wave.NewSweepFrequencySynthesizer(f0, f1, a, testSampleRate)

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

func TestToolsSave(t *testing.T) {
	f0 := 60. // Hz
	f1 := 280.
	a := 1.
	d := 2.0 // seconds

	w := wave.NewSweepFrequencySynthesizer(f0, f1, a, testSampleRate)
	s := w.Synthesize(d)

	outpath := "output.TestToolsSave.wav"
	if err := Save(NewSound(s), outpath); err != nil {
		t.Error(err)
	}
}

func TestVolumeStreamer(t *testing.T) {
	// This test creates 2 consecutive different sound with different
	// volumes. See TestVolumeStreamerAsync for an illustration of how
	// change the volume of the current streamer.
	f := 440.
	a := 1.
	d := 1.
	r := wave.DefaultSampleRate
	s := wave.NewSineWaveSynthesizer(f, a, r)
	samples := s.Synthesize(d)

	var streamers []beep.Streamer

	stream := VolumeStreamer(NewSound(samples))
	stream.Volume = 0
	streamers = append(streamers, stream)

	stream = VolumeStreamer(NewSound(samples))
	stream.Volume = 3
	streamers = append(streamers, stream)

	streamer := beep.Seq(streamers...)
	if err := Play(streamer); err != nil {
		t.Error(err)
	}
}

func TestVolumeStreamerAsync(t *testing.T) {
	// This test play one streamer sound and change the volume during
	// playing. We use for that the low level Play function of the
	// speaker with is asynchronous.
	f := 440.
	a := 1.
	d := 2.
	r := wave.DefaultSampleRate
	s := wave.NewSineWaveSynthesizer(f, a, r)
	samples := s.Synthesize(d)

	stream := VolumeStreamer(NewSound(samples))
	stream.Volume = 0
	speaker.Play(stream)

	time.Sleep(time.Duration(1 * time.Second))
	speaker.Lock()
	stream.Volume += 1
	speaker.Unlock()

	time.Sleep(time.Duration(1 * time.Second))
}

func TestSaturation(t *testing.T) {
	f := 440.
	a := 1.
	d := 1.

	// ---------------------------------------------------
	// Saturation test using sine. We observe that the sound becomes a square sound
	sound1 := NewSound(wave.SineWaveSignal(f, a, d, 0))
	sound2 := NewSound(wave.SineWaveSignal(f, 2*a, d, 0))
	sound3 := NewSound(wave.SineWaveSignal(f, 4*a, d, 0))

	streamers := []beep.Streamer{
		silence(0.2), sound1,
		silence(0.2), sound2,
		silence(0.2), sound3,
	}

	streamer := beep.Seq(streamers...)
	if err := Play(streamer); err != nil {
		t.Error(err)
	}

	// ---------------------------------------------------
	// Saturation test using KS/ It seems we can reproduce the sound of a
	// sur-saturated / distorsion guitar
	sound1 = NewSound(wave.KarplusStrongSignal(f, a, d, 0))
	sound2 = NewSound(wave.KarplusStrongSignal(f, 4*a, d, 0))
	sound3 = NewSound(wave.KarplusStrongSignal(f, 28*a, 4*d, 0))
	streamers = []beep.Streamer{
		silence(0.2), sound1,
		silence(0.2), sound2,
		silence(0.2), sound3,
	}

	streamer = beep.Seq(streamers...)
	if err := Play(streamer); err != nil {
		t.Error(err)
	}
}
