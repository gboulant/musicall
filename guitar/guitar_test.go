package guitar

import (
	"testing"

	"galuma.net/synthetic/sound"
	"galuma.net/synthetic/wave"
	"github.com/gopxl/beep"
)

const sampleRate = wave.DefaultSampleRate

func init() {
	sound.Init(sampleRate)
}

func TestGuitar_Pluck(t *testing.T) {
	g := NewGuitar(sampleRate)

	duration := 0.4

	// play the pentatonic sequence
	streamers := []beep.Streamer{
		sound.Silence(0.2, sampleRate),
		g.Pluck(Note{Mi1, 5}, duration),
		g.Pluck(Note{Mi1, 8}, duration),
		g.Pluck(Note{La1, 5}, duration),
		g.Pluck(Note{La1, 7}, duration),
		g.Pluck(Note{Re2, 5}, duration),
		g.Pluck(Note{Re2, 7}, duration),
		g.Pluck(Note{Sol2, 5}, duration),
		g.Pluck(Note{Sol2, 7}, duration),
		g.Pluck(Note{Si2, 5}, duration),
		g.Pluck(Note{Si2, 8}, duration),
		g.Pluck(Note{Mi3, 5}, duration),
		g.Pluck(Note{Mi3, 8}, duration),
	}

	streamer := beep.Seq(streamers...)
	if err := sound.Play(streamer); err != nil {
		t.Error(err)
	}

}

func TestGuitar_StandardChord(t *testing.T) {
	g := NewGuitar(sampleRate)

	duration := 0.8
	delai := 0.05
	streamers := []beep.Streamer{
		sound.Silence(0.2, sampleRate),
		g.Chord(StandardChords["Do"], duration, delai),
		g.Chord(StandardChords["Re"], duration, delai),
		g.Chord(StandardChords["Mi"], duration, delai),
		g.Chord(StandardChords["Fa"], duration, delai),
		g.Chord(StandardChords["Sol"], duration, delai),
		g.Chord(StandardChords["La"], duration, delai),
		sound.Silence(1., sampleRate),
	}
	streamer := beep.Seq(streamers...)
	if err := sound.Play(streamer); err != nil {
		t.Error(err)
	}

}
