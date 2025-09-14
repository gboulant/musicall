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
	SetFrequency(f float64)
	SetAmplitude(a float64)
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

func (s *harmonicSynthesizer) SetFrequency(f float64) {
	s.frequency = f
}

func (s harmonicSynthesizer) Amplitude() float64 {
	return s.amplitude
}

func (s *harmonicSynthesizer) SetAmplitude(a float64) {
	s.amplitude = a
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
	return &sineWaveSynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude}}
}

// -------------------------------------------------------------
// pwmWaveSynthesizer is a synthesizer for creating a square wave
type pwmWaveSynthesizer struct {
	harmonicSynthesizer
	dutycycle float64
}

// Synthesize creates a Pulse Width Modulation (PWM) wave signal
func (s pwmWaveSynthesizer) Synthesize(duration float64) []float64 {
	size := int(duration * float64(s.sampleRate))
	samples := make([]float64, size)

	period_duration_seconds := 1. / s.frequency
	samples_by_period := period_duration_seconds * float64(s.sampleRate)
	samples_by_dutycycle := samples_by_period * s.dutycycle

	for i := range samples {
		if i%int(samples_by_period) < int(samples_by_dutycycle) {
			samples[i] = 1. * s.amplitude
		} else {
			samples[i] = -1. * s.amplitude
		}
	}
	return samples
}

func NewPWMWaveSynthesizer(frequency float64, amplitude float64, sampleRate int, dutycycle float64) HarmonicSynthesizer {
	return &pwmWaveSynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude},
		dutycycle}
}

func NewSquareWaveSynthesizer(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	dutycycle := 0.5
	return &pwmWaveSynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude},
		dutycycle}
}

// -------------------------------------------------------------
// triangleWaveSynthesizer is a synthesizer for creating a triangle wave
type triangleWaveSynthesizer struct {
	harmonicSynthesizer
	risingrate float64
}

// Synthesize creates a triangle wave signal
func (s triangleWaveSynthesizer) Synthesize(duration float64) []float64 {
	size := int(duration * float64(s.sampleRate))
	samples := make([]float64, size)

	period_duration_seconds := 1. / s.frequency
	samples_by_period := period_duration_seconds * float64(s.sampleRate)
	riseslope_samples := samples_by_period * s.risingrate

	// We have to make sure that there is at least one sample point in
	// each slope (when the rising rate is close to 0 or 1)
	if riseslope_samples < 1. {
		riseslope_samples = 1
	} else if riseslope_samples > (samples_by_period - 1) {
		riseslope_samples = samples_by_period - 1
	}
	downslope_samples := samples_by_period - riseslope_samples

	riseslope_step := 2 * s.amplitude / riseslope_samples
	downslope_step := -2 * s.amplitude / downslope_samples

	value := -s.amplitude
	step := riseslope_step
	// On démarre le signal à la valeur -a avec une pente montante. La
	// pente change de sens (vers le bas) si l'amplitude dépasse
	// l'amplitude max (+a) ou si le temps dans le cycle dépasse la
	// durée montante. Elle change de sens (vers le haut) si l'amplitude
	// descend en dessous de l'amplitude minimum (-a). C'est l'algo qui
	// permet d'obtenir les signaux les plus propres même quand le
	// risingrate est proche de 0 ou 1.
	for i := range samples {
		samples[i] = value
		value += step
		if value > s.amplitude {
			value = s.amplitude
			step = downslope_step
		} else if value < -s.amplitude {
			value = -s.amplitude
			step = riseslope_step
		} else if i%int(samples_by_period) >= int(riseslope_samples) {
			step = downslope_step
		} else {
			step = riseslope_step
		}
	}
	return samples
}

func NewTriangleWaveSynthesizer(frequency float64, amplitude float64, sampleRate int, risingrate float64) HarmonicSynthesizer {
	return &triangleWaveSynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude},
		risingrate,
	}
}

func NewRegularTriangleWaveSynthesizer(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	risingrate := 0.5
	return &triangleWaveSynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude},
		risingrate,
	}
}

func NewSawtoothWaveSynthesizer(frequency float64, amplitude float64, sampleRate int) HarmonicSynthesizer {
	risingrate := 1.
	return &triangleWaveSynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			frequency:  frequency,
			amplitude:  amplitude},
		risingrate}
}

// -------------------------------------------------------------
// karplusStrongSynthesizer is a synthesizer for creating a square wave
type karplusStrongSynthesizer struct {
	harmonicSynthesizer
}

func (s karplusStrongSynthesizer) Synthesize(duration float64) []float64 {
	noise := make([]float64, int(float64(s.sampleRate)/s.frequency))
	for i := range noise {
		noise[i] = s.amplitude * (rand.Float64()*2 - 1)
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
	return &karplusStrongSynthesizer{
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

func NewSweepFrequencySynthesizer(frequencyStart, frequencyEnd float64, amplitude float64, sampleRate int) Synthesizer {
	return &sweepFrequencySynthesizer{
		harmonicSynthesizer{
			sampleRate: sampleRate,
			amplitude:  amplitude},
		frequencyStart,
		frequencyEnd,
	}
}
