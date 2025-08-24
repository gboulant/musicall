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
		g.Silence(0.2),
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
		g.Silence(0.2),
		g.Chord(StandardChord("Do"), duration, delai),
		g.Chord(StandardChord("Re"), duration, delai),
		g.Chord(StandardChord("Mi"), duration, delai),
		g.Chord(StandardChord("Fa"), duration, delai),
		g.Chord(StandardChord("Sol"), duration, delai),
		g.Chord(StandardChord("La"), duration, delai),
		g.Silence(1.),
	}
	streamer := beep.Seq(streamers...)
	if err := sound.Play(streamer); err != nil {
		t.Error(err)
	}
}

func TestGuitar_Reverse(t *testing.T) {
	g := NewGuitar(sampleRate)

	duration := 0.4
	delai := 0.03

	DoDown := StandardChord("Do")
	DoUp := Reverse(DoDown)
	ReDown := StandardChord("Re")
	ReUp := Reverse(ReDown)

	streamers := []beep.Streamer{
		g.Silence(0.2),
		g.Chord(DoDown, duration, delai),
		g.Chord(DoUp, duration, delai),
		g.Silence(0.2),
		g.Chord(ReDown, duration, delai),
		g.Chord(ReUp, duration, delai),
	}
	streamer := beep.Seq(streamers...)
	if err := sound.Play(streamer); err != nil {
		t.Error(err)
	}

}
