package wave

import (
	"math"
	"math/rand/v2"
)

const DefaultSampleRate int = 44100

// SampleRate returns a validated sample rate (number of samples per
// second). If the input value is zero or negative, then it returns the
// default sample rate value DefaultSampleRate=44100, otherwize it
// returns the input value itself.
func SampleRate(rate int) int {
	if rate <= 0 {
		return DefaultSampleRate
	}
	return rate
}

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
// sineWaveSynthesizer is a synthesizer for creating a sine wave
type sineWaveSynthesizer struct {
	harmonicSynthesizer
}

// Synthesize creates a sine wave signal
func (s sineWaveSynthesizer) Synthesize(duration float64) []float64 {
	size := int(duration * float64(s.sampleRate))
	samples := make([]float64, size)
	var angle float64 = math.Pi * 2 / float64(s.sampleRate)

	for i := range samples {
		samples[i] = s.amplitude * math.Sin(angle*s.frequency*float64(i))
	}
	return samples
}

func NewSineWaveSynthesizer(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	return sineWaveSynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude}}
}

// -------------------------------------------------------------
// squareWaveSynthesizer is a synthesizer for creating a square wave
type squareWaveSynthesizer struct {
	harmonicSynthesizer
}

// Synthesize creates a square wave signal
func (s squareWaveSynthesizer) Synthesize(duration float64) []float64 {
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

func NewSquareWaveSynthesizer(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	return squareWaveSynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude}}
}

// -------------------------------------------------------------
// karplusStrongSynthesizer is a synthesizer for creating a square wave
type karplusStrongSynthesizer struct {
	harmonicSynthesizer
}

func (s karplusStrongSynthesizer) Synthesize(duration float64) []float64 {
	noise := make([]float64, int(float64(s.sampleRate)/s.frequency))
	for i := range noise {
		noise[i] = s.amplitude*rand.Float64()*2 - 1
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

func NewKarplusStrongSynthesizer(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	return karplusStrongSynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude}}
}

// -------------------------------------------------------------
type sweepFrequencySynthesizer struct {
	harmonicSynthesizer
	frequencyMin float64
	frequencyMax float64
}

func (s sweepFrequencySynthesizer) Synthesize(duration float64) []float64 {
	size := int(duration * float64(s.sampleRate))
	samples := make([]float64, size)
	var angle float64 = math.Pi * 2 / float64(s.sampleRate)

	deltafreq := (s.frequencyMax - s.frequencyMin) / float64(size)
	var frequency float64
	for i := range samples {
		frequency = s.frequencyMin + deltafreq*float64(i)
		samples[i] = s.amplitude * math.Sin(angle*frequency*float64(i))
	}
	return samples
}

func NewSweepFrequencySynthesizer(frequencyMin, frequencyMax float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	return sweepFrequencySynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			amplitude:  amplitude},
		frequencyMin,
		frequencyMax,
	}
}
