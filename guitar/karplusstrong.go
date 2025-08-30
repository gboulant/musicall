package guitar

// THANKS: This Karplus Strong implementation is copied/adapted from the
// timiskhakov project https://github.com/timiskhakov/music (the
// extended implementation).

import (
	"math"
	"math/rand"

	"github.com/gboulant/musicall/wave"
)

// -----------------------------------------------------------------
// Implementation of the HarmonicSynthesizer interface
func NewKarplusStrongSynthesizer(frequency, amplitude, level float64, sampleRate int) wave.HarmonicSynthesizer {
	return &karplusStrongSynthesizer{
		frequency:  frequency,
		amplitude:  amplitude,
		level:      level,
		sampleRate: float64(sampleRate)}
}

type karplusStrongSynthesizer struct {
	frequency  float64
	amplitude  float64
	level      float64
	sampleRate float64
}

func (e karplusStrongSynthesizer) Amplitude() float64 {
	return e.amplitude
}
func (e *karplusStrongSynthesizer) SetAmplitude(a float64) {
	e.amplitude = a
}

func (e karplusStrongSynthesizer) Frequency() float64 {
	return e.frequency
}
func (e *karplusStrongSynthesizer) SetFrequency(f float64) {
	e.frequency = f
}

func (e karplusStrongSynthesizer) SampleRate() int {
	return int(e.sampleRate)
}

func (e karplusStrongSynthesizer) Synthesize(duration float64) []float64 {
	return e.synthesize(e.Frequency(), duration)
}

// -----------------------------------------------------------------
// Synthesize implementation

// The following is copied from/adapted from the timiskhakov project
// https://github.com/timiskhakov/music (the extended implementation).

const (
	p = 0.9
	b = 0.1
	s = 0.5
	c = 0.1
)

func (e *karplusStrongSynthesizer) synthesize(frequency float64, duration float64) []float64 {
	// Create initial noise
	noise := make([]float64, int(e.sampleRate/frequency))
	for i := range noise {
		noise[i] = e.amplitude * (rand.Float64()*2 - 1)
	}

	// Apply noise filters
	pickDirectionLowpass(noise)
	pickPositionComb(noise)

	// Create samples and add some noise in the beginning
	samples := make([]float64, int(e.sampleRate*duration))
	for i := range noise {
		samples[i] = noise[i]
	}

	// Apply single sample filters
	for i := len(noise); i < len(samples); i++ {
		samples[i] = firstOrderStringTuningAllpass(samples, i, len(noise))
	}

	// Apply all samples filters
	dynamicLevelLowpass(samples, math.Pi*frequency/e.sampleRate, e.level)

	return samples
}

func pickDirectionLowpass(noise []float64) {
	buffer := make([]float64, len(noise))
	buffer[0] = (1 - p) * noise[0]
	for i := 1; i < len(noise); i++ {
		buffer[i] = (1-p)*noise[i] + p*buffer[i-1]
	}
	noise = buffer
}

func pickPositionComb(noise []float64) {
	pick := int(b*float64(len(noise)) + 1./2.)
	if pick == 0 {
		pick = len(noise)
	}
	buffer := make([]float64, len(noise))
	for i := range noise {
		if i-pick < 0 {
			buffer[i] = noise[i]
		} else {
			buffer[i] = noise[i] - noise[i-pick]
		}
	}
	noise = buffer
}

func delayLine(samples []float64, n, N int) float64 {
	if n-N < 0 {
		return 0
	}
	return samples[n-N]
}

func oneZeroStringDamping(samples []float64, n, N int) float64 {
	return 0.996 * ((1-s)*delayLine(samples, n, N) + s*delayLine(samples, n-1, N))
}

func firstOrderStringTuningAllpass(samples []float64, n, N int) float64 {
	return c*(oneZeroStringDamping(samples, n, N)-samples[n-1]) + oneZeroStringDamping(samples, n-1, N)
}

func dynamicLevelLowpass(samples []float64, w float64, l float64) {
	buffer := make([]float64, len(samples))
	buffer[0] = w / (1 + w) * samples[0]
	for i := 1; i < len(samples); i++ {
		buffer[i] = w/(1+w)*(samples[i]+samples[i-1]) + (1-w)/(1+w)*buffer[i-1]
	}

	for i := range samples {
		samples[i] = (math.Pow(l, 4/3) * samples[i]) + (1-l)*buffer[i]
	}
}
