package wave

import (
	"log"
	"math"
	"testing"
)

func TestAddNoise(t *testing.T) {
	f := 120.
	a := 1.
	d := 4.
	r := DefaultSampleRate
	samples := SquareWaveSignal(f, a, d, r)
	log.Println(samples[:10])
	AddNoise(&samples, 0.2)
	log.Println((samples)[:10])
}

const float64EqualityThreshold = 1e-6

func almostEqual(a, b float64, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}

func TestDecimate(t *testing.T) {
	f := 120.
	a := 1.
	d := 4.
	r := DefaultSampleRate
	s := SineWaveSignal(f, a, d, r)

	// Note that the number of sample points by cycle is Nc = r*T = r/f.
	// After decimation, the number of points by cycle is reduced with a
	// factor step: Nc = r/(f*step). Then be carrefull when choosing the
	// decimation step, and take care of the cycle frequencies.
	//
	// For example, if the frequency is f=440 (La) and the sample rate
	// is r=44100~44000, then the original Nc is Nc=r/f=44000/440=100.
	// If you choose a decimation step around 100, you will have one
	// point by cycle.

	step := 10
	decimate := Decimate(s, step)

	inputsize := int(d * float64(r))
	outputsize := int(float64(inputsize) / float64(step))

	reslen := len(decimate)
	explen := outputsize
	if reslen != explen {
		t.Errorf("len is %d (should be %d)", reslen, explen)
	}

	// Create a signal with sample rate reduced of step. The create
	// signal should be the same than the decimate signal.
	r = r / step
	s = SineWaveSignal(f, a, d, r)

	if len(s) != len(decimate) {
		t.Errorf("len is %d (should be %d)", len(s), len(decimate))
	}
	for i, v := range decimate {
		if !almostEqual(v, s[i], float64EqualityThreshold) {
			t.Errorf("s[%d] is %.8f (should be %.8f)", i, s[i], v)
		}
	}
}

func TestMinMax(t *testing.T) {
	f := 120.
	a := 1.4
	d := 4.
	r := DefaultSampleRate
	samples := SineWaveSignal(f, a, d, r)

	// Add an offset for non trivial min max
	offset := 2.3
	for i := range samples {
		samples[i] = samples[i] + offset
	}

	resmin, resmax, resmed := MinMax(&samples)
	expmin := -a + offset
	expmax := +a + offset
	expmed := offset

	if !almostEqual(resmin, expmin, 1e-3) {
		t.Errorf("min is %.2f (should be %.2f)", resmin, expmin)
	}
	if !almostEqual(resmax, expmax, 1e-3) {
		t.Errorf("max is %.2f (should be %.2f)", resmax, expmax)
	}
	if !almostEqual(resmed, expmed, 1e-3) {
		t.Errorf("med is %.2f (should be %.2f)", resmed, expmed)
	}
}

func TestRescale(t *testing.T) {
	f := 5.
	a := 1.4
	d := 4.
	r := int(f * 100) // 100 point by cycle
	p := NewPlotter()

	samples := SineWaveSignal(f, a, d, r)

	// Add an offset for non trivial min max
	offset := 2.3
	for i := range samples {
		samples[i] = samples[i] + offset
	}
	p.AddLineSampledValues(samples, r, "origin")

	// Rescale to the range -1, +1
	min, max, _ := MinMax(&samples)
	Rescale(&samples, min, max, -1., +1.)
	p.AddLineSampledValues(samples, r, "rescale")
	outpath := "output.TestRescale.html"
	p.Save(outpath)

	resmin, resmax, _ := MinMax(&samples)
	expmin := -1.
	expmax := +1.
	if !almostEqual(resmin, expmin, 1e-3) {
		t.Errorf("min is %.2f (should be %.2f)", resmin, expmin)
	}
	if !almostEqual(resmax, expmax, 1e-3) {
		t.Errorf("max is %.2f (should be %.2f)", resmax, expmax)
	}
}

func TestSigmoidFilterSandbox(t *testing.T) {
	f := 10.
	a := 1.
	d := 4.
	r := int(f * 100) // 100 point by cycle
	p := NewPlotter()

	samples := SineWaveSignal(f, a, d, r)
	p.AddLineSampledValues(samples, r, "origin")

	mksigmoid := func(origin float64, lambda float64) func(x float64) float64 {
		return func(x float64) float64 {
			return 1. / (1. + math.Exp(-lambda*(x-origin)))
		}
	}

	// On peut démontrer que pour avoir une montée de la sigmoide de la
	// valeur alpha à la valeur 1-alpha (où 0 < alpha < 0.5) sur une
	// longeur d'abscice D (ici une durée), alors le paramètre lambda
	// doit être égale à:
	//
	// lambda = (2/D) * ln((1-alpha)/alpha)
	//
	// Par exemple, si on souhaite une montée caractéristique d'une
	// durée D entre 0.1 (alpha=10% du maximum) et 0.9 (90% du maximum),
	// alors on doit fixer le paramètre lambda à:
	//
	// lambda = (2/D) * ln(9)
	//
	D := 3. / f // montée sur 3 cycles
	lambda := math.Log(9.) * 2. / D
	origin := 0.5

	sigmoidfunc1 := mksigmoid(origin, lambda)
	sigmoiddata := make([]float64, len(samples))
	for i := 0; i < len(samples); i += 1 {
		t := float64(i) / float64(r)
		sigmoiddata[i] = sigmoidfunc1(t)
	}
	p.AddLineSampledValues(sigmoiddata, r, "sigmoid1")

	filtered1 := make([]float64, len(samples))
	for i := 0; i < len(samples); i += 1 {
		t := float64(i) / float64(r)
		filtered1[i] = samples[i] * sigmoidfunc1(t)
	}
	p.AddLineSampledValues(filtered1, r, "filtered1")

	origin = d - origin
	lambda = -lambda
	sigmoidfunc2 := mksigmoid(origin, lambda)
	sigmoiddata = make([]float64, len(samples))
	for i := 0; i < len(samples); i += 1 {
		t := float64(i) / float64(r)
		sigmoiddata[i] = sigmoidfunc2(t)
	}
	p.AddLineSampledValues(sigmoiddata, r, "sigmoid2")

	filtered2 := make([]float64, len(samples))
	for i := 0; i < len(samples); i += 1 {
		t := float64(i) / float64(r)
		filtered2[i] = samples[i] * sigmoidfunc2(t)
	}
	p.AddLineSampledValues(filtered2, r, "filtered2")

	filtered3 := make([]float64, len(samples))
	for i := 0; i < len(samples); i += 1 {
		t := float64(i) / float64(r)
		filtered3[i] = samples[i] * sigmoidfunc1(t) * sigmoidfunc2(t)
	}
	p.AddLineSampledValues(filtered3, r, "filtered3")

	outpath := "output.TestSigmoidFilter.html"
	p.Save(outpath)
}

func TestApplyFilter_sigmoid(t *testing.T) {
	f := 10.
	a := 1.
	d := 4.
	r := int(f * 100) // 100 point by cycle
	p := NewPlotter()

	samples := SineWaveSignal(f, a, d, r)
	p.AddLineSampledValues(samples, r, "origin")

	// Apply a first filter, rising up at the begining
	origin := 0.5
	risingtime := 3. / f // montée sur 3 cycles
	filter1 := NewSigmoidFilterByRisingTime(origin, risingtime)
	ApplyTimeFilter(&samples, r, filter1)
	p.AddLineSampledValues(samples, r, "filtered1")

	// Apply a second filter, rising down at the end. To have a rising
	// down sigmoid just get the inverse of the lambda parameter
	origin = d - origin
	risingtime = -risingtime
	filter2 := NewSigmoidFilterByRisingTime(origin, risingtime)
	ApplyTimeFilter(&samples, r, filter2)
	p.AddLineSampledValues(samples, r, "filtered2")

	outpath := "output.TestApplyFilter_sigmoid.html"
	p.Save(outpath)
}

func TestSmoothBoundaries(t *testing.T) {
	f := 10.
	a := 1.
	d := 4.
	r := int(f * 100) // 100 point by cycle
	p := NewPlotter()

	samples := SineWaveSignal(f, a, d, r)
	p.AddLineSampledValues(samples, r, "origin")

	smoothtime := 3. / f // fondu sur 3 cycles
	SmoothBoundaries(&samples, r, smoothtime)
	p.AddLineSampledValues(samples, r, "smooth")

	outpath := "output.TestSmoothBoundaries.html"
	p.Save(outpath)
}
