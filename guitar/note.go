package guitar

import (
	"github.com/gboulant/musicall"
	"github.com/gboulant/musicall/music"
)

type StringNumber int
type FretNumber music.Interval

// Les cordes de guitare sont numérotées de 1 (la corde la plus
// aigue, le Mi3) à 6 la corde la plus grave (Mi1)
const (
	Mi3 StringNumber = iota + 1
	Si2
	Sol2
	Re2
	La1
	Mi1
)

// Les cordes de guitare sont Mi1, La1, Re2, Sol2, Si2, Mi3
// soit en notation chromatique: (1,4), (1,9), (2,2), (2,7), (2,11),
// (3,4). On peut remarquer que l'intervalle entre deux notes
// successives est de 5 demi-tons, sauf entre Sol2 et Si2, où
// l'intervalle est de 4 demi-tons. Soit un rapport de fréquences
// 2^5/12 ~ 1.334, soit environ 4/3, et 2⁴/12 ~ 1.259, soit environ
// 5/4.

// noteOfOpenString returns the music.Note played when we pluck the
// specified open string (without pressing any fret).
func noteOfOpenString(stringNum StringNumber) music.Note {
	note, ok := openStringNotes[stringNum]
	if !ok {
		musicall.LogError("err: (noteOfOpenString) the string number %d is not defined\n", stringNum)
	}
	return note
}

var openStringNotes map[StringNumber]music.Note = map[StringNumber]music.Note{
	Mi3:  {Octave: 3, Index: music.Label2Index("Mi")},
	Si2:  {Octave: 2, Index: music.Label2Index("Si")},
	Sol2: {Octave: 2, Index: music.Label2Index("Sol")},
	Re2:  {Octave: 2, Index: music.Label2Index("Re")},
	La1:  {Octave: 1, Index: music.Label2Index("La")},
	Mi1:  {Octave: 1, Index: music.Label2Index("Mi")},
}

// Note defines the musical note from a guitar point of view, i.e. by
// specifying the string number to pluck and the fret number to press.
// Concerning the string number, the convention is to designate the bass
// Mi string (string at the top of the guitar neck) as the string number
// 6, and the high Mi string (at the bottom of the neck) as the string
// number 1. To avoid error, you should use the constant definitions
// above (Mi3=1, Si2=2, ..., Mi1=6). A fret number of 0 means "don't
// press any fret (play the open string)".
type Note struct {
	StringNum StringNumber
	FretNum   FretNumber
}

// MusicNote returns the music note corresponding to this guitar note. A music
// note is defined in terms of an octave number and an index in this octave.
func (n Note) MusicNote() music.Note {
	// 1. On récupère la note de la corde à vide
	note := noteOfOpenString(n.StringNum)
	// 2. On ajoute l'intervalle de la frette (nb de demi-tons)
	note.Add(music.Interval(n.FretNum))
	return note
}

// Frequency return the major frequency of the sound corresponding to
// this note. This frequency can be used in a synthesiser for generating
// the sound signal.
func (n Note) Frequency() float64 {
	return n.MusicNote().Frequency()
}

// Name return the name of this note (Do, Ré, Mi, etc.).
func (n Note) Name() string {
	return n.MusicNote().Name()
}
