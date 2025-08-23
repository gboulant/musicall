package sound

import (
	"fmt"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/generators"
	"github.com/gopxl/beep/speaker"
)

// Init intialize le speaker avec un taux d'échantillonage (sample rate)
// fixé par le paramètre sampleRate. Il correspond au nombre de points
// (samples) par seconde. Attention, on ne peux pas initialiser le
// speaker deux fois dans un même programme. Tous les signaux ([]float64) joués par le
// speaker ainsi initialisé seront considérés comme des sons avec ce
// taux d'échantillonage.
func Init(sampleRate int) error {
	beepSampleRate := beep.SampleRate(sampleRate)
	err := speaker.Init(beepSampleRate, beepSampleRate.N(time.Second/10))
	if err != nil {
		return err
	}
	return nil
}

func Play(s beep.Streamer) error {
	// Note that speaker.Play is an asynchronous function, then we play
	// 2 streamers, the second being a callback that triggers the channel
	done := make(chan bool, 1)
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		done <- true
	})))
	<-done
	return nil
}

// LabelledStreamer return the input streamer with a preprocessing
// action that print the specified label on the standard output.
func LabelledStreamer(s beep.Streamer, label string) beep.Streamer {
	return beep.Seq(beep.Callback(func() {
		fmt.Println(label)
	}), s)
}

func Silence(duration float64, sampleRate int) beep.Streamer {
	return generators.Silence(int(duration * float64(sampleRate)))
}
