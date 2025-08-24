package music

import (
	"math"

	"galuma.net/synthetic"
)

// We have to defined at least one frequency, and the other frequencies
// can be derived from the interval with this reference frequency using
// the following (we consider the tempered equal scale):
//
// log(Frequency) = log(FrequencyLa3) + interval/12
//
// We use the La3 as reference frequency
//

const FrequencyLa3 float64 = 440.0 // Herz
var logFrequencyLa3 float64 = math.Log2(FrequencyLa3)

var La3 Note = Note{3, 9}

type Interval int

var HalfTone Interval = 1
var Tone Interval = 2 * HalfTone
var Octave Interval = 6 * Tone
var Quinte Interval = 3*Tone + HalfTone

type NoteIndex Interval

// Label2Index can be used to get the note index in an octave from its
// symbolic name (Do, Ré, etc.). We designate the "Ré" as "Re" (without
// accents) so that it can be used even with a qwerty keyboard ;-).
func Label2Index(label string) NoteIndex {
	index, ok := label2Index[label]
	if !ok {
		synthetic.LogError("err: (Label2Index) the label %s does not exist\n", label)
	}
	return index
}

var label2Index map[string]NoteIndex = map[string]NoteIndex{
	"Do":   0,
	"Do#":  1,
	"Re":   2,
	"Re#":  3,
	"Mi":   4,
	"Fa":   5,
	"Fa#":  6,
	"Sol":  7,
	"Sol#": 8,
	"La":   9,
	"La#":  10,
	"Si":   11,
}

// A noter (pour la culture):
// L'ensemble de ces notes (de 0 à 11) compose la gamme dite chromatique.
// La restriction aux notes non altérées compose la gamme dites diatonique

type Note struct {
	Octave int       // index of the octave where is considered the note
	Index  NoteIndex // index of the note in the octave, counted in number of half-tones
}

func (n *Note) Add(interval Interval) {
	i := Interval(n.Octave*int(Octave) + int(n.Index))
	i += interval
	n.Octave = int(i / Octave)
	n.Index = NoteIndex(i % Octave)
}

func (n Note) IntervalTo(other Note) Interval {
	interval1 := n.Octave*int(Octave) + int(n.Index) // intervals from the reference Do0
	interval2 := other.Octave*int(Octave) + int(other.Index)
	return Interval(interval2 - interval1)
}

func (n Note) Derived(interval Interval) Note {
	note := Note{n.Octave, n.Index}
	note.Add(interval)
	return note
}

func (n Note) Frequency() float64 {
	interval := La3.IntervalTo(n)
	logF := logFrequencyLa3 + float64(interval)/float64(Octave)
	return math.Pow(2, logF)
}
