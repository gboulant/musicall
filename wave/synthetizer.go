package wave

import (
	"math"
	"math/rand/v2"
)

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
type harmonicSynthesizer struct {
	sampleRate int
	frequency  float64
	amplitude  float64
}

func (s harmonicSynthesizer) SampleRate() int {
	return s.sampleRate
}

func (s harmonicSynthesizer) Frequency() float64 {
	return s.frequency
}

func (s harmonicSynthesizer) Amplitude() float64 {
	return s.amplitude
}

// -------------------------------------------------------------
type SineWave struct {
	harmonicSynthesizer
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

func NewSineWave(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	return SineWave{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude}}
}

// -------------------------------------------------------------
type SquareWave struct {
	harmonicSynthesizer
}

// Synthesize creates a square wave signal
func (s SquareWave) Synthesize(duration float64) []float64 {
	size := int(duration * float64(s.sampleRate))
	samples := make([]float64, size)

	period_seconds := 1. / s.frequency
	period_samples := period_seconds * float64(s.sampleRate)
	halfperiod_samples := period_samples * 0.5

	for i := range samples {
		if i%int(period_samples) < int(halfperiod_samples) {
			samples[i] = 1. * s.amplitude
		} else {
			samples[i] = -1. * s.amplitude
		}
	}
	return samples

}

func NewSquareWave(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	return SquareWave{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude}}
}

// -------------------------------------------------------------
type KarplusStrongWave struct {
	harmonicSynthesizer
}

func (s KarplusStrongWave) Synthesize(duration float64) []float64 {
	noise := make([]float64, int(float64(s.sampleRate)/s.frequency))
	for i := range noise {
		noise[i] = rand.Float64()*2 - 1
	}
	// the buffer noise has a duration equal to the period of the signal
	// (1/f). And then we repeatedly copy this buffer for any period that
	// constitutes the whole signal

	size := int(duration * float64(s.sampleRate))
	samples := make([]float64, size)
	copy(samples, noise)

	for i := len(noise) + 1; i < len(samples); i++ {
		samples[i] = (samples[i-len(noise)] + samples[i-len(noise)-1]) / 2
	}
	return samples
}

func NewKarplusStrongWave(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	return KarplusStrongWave{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude}}
}
