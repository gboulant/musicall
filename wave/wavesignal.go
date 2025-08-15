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
	w := NewSquareWaveSynthesizer(f, a, DefaultSampleRate)
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
// fmin (Hz) to fmax (Hz), with amplitude a (1=nominal), duration d
// (sec), and with a samplerate r (number of samples in a sec). If r is
// zero or negative, then the default value DefaultSampleRate is
// applied. If reverse is true, then the sweep go from fmax to fmin.
func SweepFrequencySignal(fmin, fmax, a, d float64, r int, reverse bool) []float64 {
	w := NewSweepFrequencySynthesizer(fmin, fmax, a, SampleRate(r))
	s := w.Synthesize(d)
	if reverse {
		Reverse(&s)
	}
	return s
}
