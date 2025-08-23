package guitar

import (
	"testing"

	"galuma.net/synthetic/music"
)

func TestNote_Frequency(t *testing.T) {
	type fields struct {
		StringNum StringNumber
		FretNum   FretNumber
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"Mi3", fields{Mi3, 0}, music.Note{Octave: 3, Index: music.Label2Index["Mi"]}.Frequency()},
		{"Si2", fields{Si2, 0}, music.Note{Octave: 2, Index: music.Label2Index["Si"]}.Frequency()},
		{"Sol2", fields{Sol2, 0}, music.Note{Octave: 2, Index: music.Label2Index["Sol"]}.Frequency()},
		{"Ré2", fields{Re2, 0}, music.Note{Octave: 2, Index: music.Label2Index["Ré"]}.Frequency()},
		{"La1", fields{La1, 0}, music.Note{Octave: 1, Index: music.Label2Index["La"]}.Frequency()},
		{"Mi1", fields{Mi1, 0}, music.Note{Octave: 1, Index: music.Label2Index["Mi"]}.Frequency()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Note{
				StringNum: tt.fields.StringNum,
				FretNum:   tt.fields.FretNum,
			}
			if got := n.Frequency(); got != tt.want {
				t.Errorf("Note.Frequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
