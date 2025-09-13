package wave

// SineWaveSignal creates a sine wave of frequency f (Hz), amplitude a
// (1=nominal), duration d (sec), and with a samplerate r (number of
// samples in a sec). If r is zero or negative, then the default value
// DefaultSampleRate is applied.
func SineWaveSignal(f, a, d float64, r int) []float64 {
	w := NewSineWaveSynthesizer(f, a, SampleRate(r))
	s := w.Synthesize(d)
	return s
}

// SquareWaveSignal creates a sine wave of frequency f (Hz), amplitude a
// (1=nominal), duration d (sec), and with a samplerate r (number of
// samples in a sec). If r is zero or negative, then the default value
// DefaultSampleRate is applied.
func SquareWaveSignal(f, a, d float64, r int) []float64 {
	w := NewSquareWaveSynthesizer(f, a, SampleRate(r))
	s := w.Synthesize(d)
	return s
}

// KarplusStrongSignal creates a Karplus-Strong wave (emulation of a
// guitar sound) of frequency f (Hz), amplitude a (1=nominal), duration
// d (sec), and with a samplerate r (number of samples in a sec). If r
// is zero or negative, then the default value DefaultSampleRate is
// applied.
func KarplusStrongSignal(f, a, d float64, r int) []float64 {
	w := NewKarplusStrongSynthesizer(f, a, SampleRate(r))
	s := w.Synthesize(d)
	return s
}

// SweepFrequencySignal creates a sine wave whose frequency sweeps from
// fstart (Hz) to fstop (Hz), with amplitude a (1=nominal), duration d
// (sec), and with a samplerate r (number of samples in a sec). If r is
// zero or negative, then the default value DefaultSampleRate is
// applied. The frequency fstart can be greater than fstop (for a decreasing
// frequencies series).
func SweepFrequencySignal(fstart, fstop, a, d float64, r int) []float64 {
	w := NewSweepFrequencySynthesizer(fstart, fstop, a, SampleRate(r))
	s := w.Synthesize(d)
	return s
}

// PWMWaveSignal creates a PWM (Pulse Width Modulation) square wave with a
// dutycycle to be given as a number between 0 (no ON cycle) and 1 (full ON
// cycle). A square wave is a special case of the PWM with dutycycle = 0.5.
func PWMWaveSignal(f, a, d float64, dutycycle float64, r int) []float64 {
	w := NewPWMWaveSynthesizer(f, a, SampleRate(r), dutycycle)
	s := w.Synthesize(d)
	return s
}

// TriangleWaveSignal creates a triangle wave from with a rising cycle equal to
// risingrate, between 0 (vertical slope at the begining) and 1 (vertical slope
// at the end of the cycle. The regular triangle wave form is obtained using
// risingrate=0.5. The SawTooth wave form is a special case with risingrate = 1.
func TriangleWaveSignal(f, a, d float64, risingrate float64, r int) []float64 {
	w := NewTriangleWaveSynthesizer(f, a, SampleRate(r), risingrate)
	s := w.Synthesize(d)
	return s
}

// SawToothWaveSignal creates a SawTooth wave form (special case of triangle
// wave form with risingrate = 1).
func SawToothWaveSignal(f, a, d float64, r int) []float64 {
	w := NewSawtoothWaveSynthesizer(f, a, SampleRate(r))
	s := w.Synthesize(d)
	return s
}
