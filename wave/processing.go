package wave

import "math/rand/v2"

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
