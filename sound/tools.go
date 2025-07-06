package sound

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

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
