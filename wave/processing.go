package wave

import "math/rand/v2"

func AddNoise(samples *[]float64, amplitude float64) {
	var noise float64
	for i := range *samples {
		noise = amplitude * (rand.Float64()*2 - 1)
		(*samples)[i] = (*samples)[i] + noise
	}
}
