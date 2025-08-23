package guitar

import (
	"galuma.net/synthetic/sound"
	"galuma.net/synthetic/wave"
	"github.com/gopxl/beep"
)

type Guitar struct {
	synthesizer wave.HarmonicSynthesizer
}

func NewGuitar(sampleRate int) *Guitar {
	f := 0. // no specific frequency at initialize step
	a := 1.
	r := wave.SampleRate(sampleRate)
	s := wave.NewKarplusStrongSynthesizer(f, a, r)
	g := Guitar{synthesizer: s}
	return &g
}

//func (g Guitarilence(duration float64) beep.Streamer {
//	return generators.Silence(int(duration * float64(sampleRate)))
//}

func (g Guitar) Pluck(note Note, duration float64) beep.Streamer {
	frequency := 10. // should be computed from Note
	g.synthesizer.SetFrequency(frequency)
	samples := g.synthesizer.Synthesize(duration)
	return sound.NewSound(samples)
}

func (g Guitar) Chord(notes []Note, duration float64, delay float64) beep.Streamer {
	streamers := make([]beep.Streamer, len(notes))
	for i, note := range notes {
		starttime := delay * float64(i)
		start := sound.Silence(starttime, g.synthesizer.SampleRate())
		sound := g.Pluck(note, duration-starttime)
		streamers[i] = beep.Seq(start, sound)
	}
	return beep.Mix(streamers...)
}
