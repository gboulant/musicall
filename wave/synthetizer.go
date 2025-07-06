package wave

import "math"

const DefaultSampleRate int = 44100

// Synthetizer is an interface for functions that create a signal at a
// given samplerate for a given duration
type Synthesizer interface {
	SampleRate() int
	Synthesize(duration float64) []float64
}

type HarmonicSynthesizer interface {
	Synthesizer
	Frequency() float64
	Amplitude() float64
}

// -------------------------------------------------------------
type SineWave struct {
	sampleRate int
	frequency  float64
	amplitude  float64
}

// Synthesize creates a sine wave signal
func (s SineWave) Synthesize(duration float64) []float64 {
	size := int(duration * float64(s.sampleRate))
	samples := make([]float64, size)
	var angle float64 = math.Pi * 2 / float64(s.sampleRate)

	for i := range samples {
		samples[i] = s.amplitude * math.Sin(angle*s.frequency*float64(i))
	}
	return samples
}

func (s SineWave) SampleRate() int {
	return s.sampleRate
}

func (s SineWave) Frequency() float64 {
	return s.frequency
}

func (s SineWave) Amplitude() float64 {
	return s.amplitude
}

func NewSineWave(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	return SineWave{sampleRate: sampleRate, frequency: frequency, amplitude: amplitude}
}
