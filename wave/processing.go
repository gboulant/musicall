package wave

import (
	"math"
	"math/rand/v2"
)

func AddNoise(samples *[]float64, amplitude float64) {
	var noise float64
	for i := range *samples {
		noise = amplitude * (rand.Float64()*2 - 1)
		(*samples)[i] = (*samples)[i] + noise
	}
}

func Reverse(samples *[]float64) {
	size := len(*samples)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		(*samples)[i], (*samples)[j] = (*samples)[j], (*samples)[i]
	}
}

func Decimate(samples []float64, step int) []float64 {
	inputsize := len(samples)
	outputsize := int(float64(inputsize) / float64(step))
	decimate := make([]float64, outputsize)
	for i := range outputsize {
		decimate[i] = samples[i*step]
	}
	return decimate
}

// MinMax returns the minimum, maximum and medium values of the series
func MinMax(samples *[]float64) (min, max, med float64) {
	min = math.Inf(+1)
	max = math.Inf(-1)
	med = 0.
	for _, v := range *samples {
		med += v
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	med = med / float64(len(*samples))
	return min, max, med
}

// Rescale changes the range of the input signal by applying a
// multiplication factor and adding an offset, so that the input range
// (inmin, inmax) is transposed to the output range (outmin, outmax).
// Mathematically, it corresponds to an affin transformation:
//
//	Vout = a * Vin +b
//
// With:
//
//	a = (outmax - outmin) / (inmax - inmin)
//	b = (inmax*outmin - inmin*outmax) / (inmax - inmin)
func Rescale(samples *[]float64, inmin, inmax float64, outmin, outmax float64) (a, b float64) {
	a = (outmax - outmin) / (inmax - inmin)
	b = (inmax*outmin - inmin*outmax) / (inmax - inmin)
	for i := range *samples {
		(*samples)[i] = a*(*samples)[i] + b
	}
	return a, b
}

// Normalize rescales the signal in a range -1, +1
func Normalize(samples *[]float64) (a, b float64) {
	min, max, _ := MinMax(samples)
	return Rescale(samples, min, max, -1., +1.)
}

// Times returns a set of timestamps in second, considering a series of the
// specified size (total number of points) sampled to the specified sample rate,
// and starting at t0 (in seconds). If d is the duration in seconds, then the
// size is d*samplerate. If you have a dataset of values (samples) then the size
// is the number of points of this dataset (len(samples))
func Times(size int, samplerate int, t0 float64) []float64 {
	times := make([]float64, size)
	for i := range size {
		times[i] = t0 + float64(i)/float64(samplerate)
	}
	return times
}
