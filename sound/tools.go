package sound

import (
	"fmt"
	"os"
	"time"

	"github.com/gboulant/musicall/wave"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/generators"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

var samplerate int = wave.DefaultSampleRate

// Init intialize le speaker avec un taux d'échantillonage (sample rate)
// fixé par le paramètre sampleRate. Il correspond au nombre de points
// (samples) par seconde. Attention, on ne peux pas initialiser le
// speaker deux fois dans un même programme. Tous les signaux ([]float64) joués par le
// speaker ainsi initialisé seront considérés comme des sons avec ce
// taux d'échantillonage.
func Init(sampleRate int) error {
	beepSampleRate := beep.SampleRate(sampleRate)
	if err := speaker.Init(beepSampleRate, beepSampleRate.N(time.Second/10)); err != nil {
		return err
	}
	samplerate = sampleRate
	return nil
}

func Silence(duration float64, sampleRate int) beep.Streamer {
	return generators.Silence(int(duration * float64(sampleRate)))
}

func Play(s beep.Streamer) error {
	// Note that speaker.Play is an asynchronous function, then we play
	// 2 streamers, the second being a callback that triggers the
	// channel, so that this Play function is synchronous
	done := make(chan bool, 1)
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		done <- true
	})))
	<-done
	return nil
}

func Format(sampleRate int) beep.Format {
	r := beep.SampleRate(sampleRate)
	return beep.Format{SampleRate: r, NumChannels: 2, Precision: 3}
}

func Save(s beep.Streamer, outpath string) error {
	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	format := Format(samplerate)
	if err = wav.Encode(f, s, format); err != nil {
		return err
	}
	fmt.Printf("Sound saved to %s\n", outpath)
	return nil
}

// -------------------------------------------------------------
// Smart streamers

// LabelledStreamer return the input streamer with a preprocessing
// action that print the specified label on the standard output.
func LabelledStreamer(s beep.Streamer, label string) beep.Streamer {
	return beep.Seq(beep.Callback(func() {
		fmt.Println(label)
	}), s)
}

// VolumeStreamer returns the input streamer with a volume parameter.
// The volume is set to 0 by default, meaning the standard volume. Add
// +1 means a multiplication of the sound power by Base = 2.
func VolumeStreamer(s beep.Streamer) *effects.Volume {
	return &effects.Volume{
		Streamer: s,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
}
