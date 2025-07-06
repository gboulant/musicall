package wave

func SineWaveSignal(f float64, a float64, d float64) []float64 {
	w := NewSineWave(f, a, DefaultSampleRate)
	s := w.Synthesize(d)
	return s
}

func SquareWaveSignal(f float64, a float64, d float64) []float64 {
	w := NewSquareWave(f, a, DefaultSampleRate)
	s := w.Synthesize(d)
	return s
}

func KarplusStrongWaveSignal(f float64, a float64, d float64) []float64 {
	w := NewKarplusStrongWave(f, a, DefaultSampleRate)
	s := w.Synthesize(d)
	return s
}
