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

/* IMPORTANT REMARK for the computation of sinus angle

A sine wave is a signal whose value S at time t is:

 S(t) = A * sin(2*Pi * t/T)

Where:

- A is the amplitude of the signal
- T is the period of the signal (when t=T, a cycle is done)

Our parameters are f the frequency and r the samplerate, i.e. the number
of samples in one second. Then we can rewrite the variables:

- t = i / r (where i is the index of the sample, when i=r then t=1)
- T = 1 / f

Then the value of index i of the signal is:

 S(i) = A * sin(2*Pi * i*f/r)

The index should go from i=0 (first sample) to i=d * r (last sample)
where d is the duration of the signal in second (r points in one
second).

IMPORTANT REMARK (to understand the implementations we use for computing
the angle of the sinus):

The angle of the sinus is a(i) = 2*Pi * i*f/r. It is a first way to
calculate this value and compute the sinus value. An alternative is to
consider that we add an increment of 2*Pi*f/r to the angle value at each
sample. This second manner should be prefered in the case where we
execute a sweeping of the frequency, or a frequency modulation to be
sure that the angle increase at each step. With the formula a(i) =
2*Pi*i*f/r and a frequency rising down, you may experiment an angle
value that could decrease for some values of i even if i is increasing,
resulting to an unpedictable signal (at least not the signal you would
expect).

In the following, we use the second way to implement the angle
increment, i.e. we add the value 2*Pi*f/r at each step.

*/

// Synthesize creates a sine wave signal
func (s sineWaveSynthesizer) Synthesize(duration float64) []float64 {
	size := int(duration * float64(s.sampleRate))
	samples := make([]float64, size)
	var angleIncrement float64 = math.Pi * 2 * s.frequency / float64(s.sampleRate)

	var angle float64 = 0.
	for i := range samples {
		samples[i] = s.amplitude * math.Sin(angle)
		angle += angleIncrement
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
	frequencyStart float64
	frequencyEnd   float64
}

func (s sweepFrequencySynthesizer) Synthesize(duration float64) []float64 {
	size := int(duration * float64(s.sampleRate))
	samples := make([]float64, size)
	var angleIncrementFactor float64 = math.Pi * 2 / float64(s.sampleRate)

	deltafreq := (s.frequencyEnd - s.frequencyStart) / float64(size)
	var angle float64 = 0.
	var frequency float64 = s.frequencyStart
	for i := range samples {
		samples[i] = s.amplitude * math.Sin(angle)
		frequency += deltafreq
		angle += angleIncrementFactor * frequency
	}
	return samples
}

func NewSweepFrequencySynthesizer(frequencyStart, frequencyEnd float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	return sweepFrequencySynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			amplitude:  amplitude},
		frequencyStart,
		frequencyEnd,
	}
}
