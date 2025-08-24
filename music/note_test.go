package music

import (
	"math"
	"testing"

	"galuma.net/synthetic/sound"
	"galuma.net/synthetic/wave"
	"github.com/gopxl/beep"
)

const sampleRate = wave.DefaultSampleRate

func init() {
	sound.Init(sampleRate)
}

func TestNote_Interval(t *testing.T) {
	n1 := Note{3, 9} // La3
	n2 := Note{0, 0} // Do0

	res := n1.IntervalTo(n2)
	exp := Interval(-45)
	if res != exp {
		t.Errorf("interval is %d (should be %d)", res, exp)
	}

}

func almostEqual(a, b float64, accuracy float64) bool {
	return math.Abs(a-b) < accuracy
}

func TestNote_Frequency(t *testing.T) {
	type fields struct {
		Octave int
		Index  NoteIndex
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"Do0", fields{0, Label2Index("Do")}, 32.703},
		{"Ré1", fields{1, Label2Index("Re")}, 73.416},
		{"Mi2", fields{2, Label2Index("Mi")}, 164.813},
		{"Fa2", fields{2, Label2Index("Fa")}, 174.614},
		{"Sol2", fields{2, Label2Index("Sol")}, 195.997},
		{"La3", fields{3, Label2Index("La")}, 440.000},
		{"Si3", fields{3, Label2Index("Si")}, 493.883},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Note{
				Octave: tt.fields.Octave,
				Index:  tt.fields.Index,
			}

			if got := n.Frequency(); !almostEqual(got, tt.want, 1e-3) {
				t.Errorf("Note.Frequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNote_GammeChromatique(t *testing.T) {

	type labelledNote struct {
		note  Note
		label string
	}

	newLabelledNote := func(octave int, label string) labelledNote {
		return labelledNote{
			note:  Note{octave, Label2Index(label)},
			label: label,
		}
	}

	octave := 3
	notes := []labelledNote{
		newLabelledNote(octave, "Do"),
		newLabelledNote(octave, "Re"),
		newLabelledNote(octave, "Mi"),
		newLabelledNote(octave, "Fa"),
		newLabelledNote(octave, "Sol"),
		newLabelledNote(octave, "La"),
		newLabelledNote(octave, "Si"),
		newLabelledNote(octave+1, "Do"),
	}

	var newSynthesizer func(f, a float64, r int) wave.HarmonicSynthesizer

	// Uncomment the desired synthesizer builder
	newSynthesizer = wave.NewSineWaveSynthesizer
	newSynthesizer = wave.NewKarplusStrongSynthesizer

	a := 2.
	r := sampleRate
	s := newSynthesizer(0., a, r)
	d := 0.8

	var streamers []beep.Streamer
	streamers = append(streamers, sound.Silence(0.2, sampleRate))

	for _, note := range notes {
		s.SetFrequency(note.note.Frequency())
		samples := s.Synthesize(d)

		stream := sound.LabelledStreamer(sound.NewSound(samples), note.label)
		streamers = append(streamers, stream)
	}

	streamer := beep.Seq(streamers...)
	if err := sound.Play(streamer); err != nil {
		t.Error(err)
	}

}

func TestNote_Add(t *testing.T) {
	n := Note{0, 0} // Do0
	n.Add(45)

	exp := Note{3, 9} // La3
	if n.Octave != exp.Octave || n.Index != exp.Index {
		t.Errorf("result is %v (should be %v)", n, exp)
	}

}
