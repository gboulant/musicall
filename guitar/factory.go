package guitar

import "github.com/gboulant/musicall"

// ----------------------------------------------------------------------
// Définition des accord principaux

// Chord defines a sequence of notes to play as a chord. The notes are
// plucked in the order defined by the list.
type Chord []Note

// Reverse can be used to create a chord as the reverse of the input
// chord. For example, if you define a chord as the standard chord,
// plucking the string from the top to the bottom (down chord), then the
// reverse chord is this same chord but plucking the strings from the
// bottom to the top (up chord).
func Reverse(notes Chord) Chord {
	l := len(notes)
	r := make(Chord, l)
	for i, n := range notes {
		r[l-i-1] = n
	}
	return r
}

// StandardChords is a table of the most standard chords, identified
// with the french naming convention (Do, Re, Mi, etc.).
//
// The first index of a note is the string number to pluck and the
// second is the fret number to press: Note{string number to pluck, fret
// number to press}.
//
// Remember that the convention is to designate the bass Mi string
// (string at the top of the guitar neck) as the string number 6, and
// the high Mi string (at the bottom of the neck) as the string number
// 1. For example the first note of the Do chord Note{5, 3} means:
// "pluck the string number 5 (the La1, counting starting from the
// bottom), and press the fret number 3". A fret number of 0 means
// "don't press any fret (play the open string)". If a string number
// does not appear in a chord list, it means that the string must not be
// plucked.
func StandardChord(name string) Chord {
	chord, ok := standardChords[name]
	if !ok {
		musicall.LogError("err: (StandardChord) no chord with name %s\n", name)
	}
	return chord
}

var standardChords map[string]Chord = map[string]Chord{
	"Do": {
		Note{5, 3},
		Note{4, 2},
		Note{3, 0},
		Note{2, 1},
		Note{1, 0},
	},
	"Re": {
		Note{4, 0},
		Note{3, 2},
		Note{2, 3},
		Note{1, 2},
	},
	"Mi": {
		Note{6, 0},
		Note{5, 2},
		Note{4, 2},
		Note{3, 1},
		Note{2, 0},
		Note{1, 0},
	},
	"Mim": {
		Note{6, 0},
		Note{5, 2},
		Note{4, 2},
		Note{3, 0},
		Note{2, 0},
		Note{1, 0},
	},
	"Fa": {
		Note{5, 0},
		Note{4, 3},
		Note{3, 2},
		Note{2, 1},
	},
	"Sol": {
		Note{6, 3},
		Note{5, 2},
		Note{4, 0},
		Note{3, 0},
		Note{2, 0},
		Note{1, 3},
	},
	"La": {
		Note{5, 0},
		Note{4, 2},
		Note{3, 2},
		Note{2, 2},
		Note{1, 0},
	},
	"Lam": {
		Note{5, 0},
		Note{4, 2},
		Note{3, 2},
		Note{2, 1},
		Note{1, 0},
	},
}

func PowerChord(stringnum StringNumber, fretnum FretNumber) Chord {
	return Chord{
		{StringNum: stringnum, FretNum: fretnum},
		{StringNum: stringnum - 1, FretNum: fretnum + 2},
	}
}
