package wave

// TODO, the signal function should return their samplerate, or let the
// possibility to choose it?

func SineWaveSignal(f float64, a float64, d float64) []float64 {
	w := NewSineWaveSynthesizer(f, a, DefaultSampleRate)
	s := w.Synthesize(d)
	return s
}

func SquareWaveSignal(f float64, a float64, d float64) []float64 {
	w := NewSquareWaveSynthesizer(f, a, DefaultSampleRate)
	s := w.Synthesize(d)
	return s
}

func KarplusStrongSignal(f float64, a float64, d float64) []float64 {
	w := NewKarplusStrongSynthesizer(f, a, DefaultSampleRate)
	s := w.Synthesize(d)
	return s
}
