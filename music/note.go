package music

import (
	"math"
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
// symbolic name (Do, Ré, etc.).
var Label2Index map[string]NoteIndex = map[string]NoteIndex{
	"Do":   0,
	"Do#":  1,
	"Ré":   2,
	"Ré#":  3,
	"Mi":   4,
	"Fa":   5,
	"Fa#":  6,
	"Sol":  7,
	"Sol#": 8,
	"La":   9,
	"La#":  10,
	"Si":   11,
}

type Note struct {
	Octave int       // index of the octave where is considered the note
	Index  NoteIndex // index of the note in the octave, counted in number of half-tones
}

func (n *Note) Add(interval Interval) {
	n.Octave = n.Octave + int(interval/Octave)
	n.Index = NoteIndex(int(n.Index) % int(Octave))
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
