package guitar

import "galuma.net/synthetic/music"

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

var stringNotes map[StringNumber]music.Note = map[StringNumber]music.Note{
	Mi3:  {Octave: 3, Index: music.Label2Index["Mi"]},
	Si2:  {Octave: 2, Index: music.Label2Index["Si"]},
	Sol2: {Octave: 2, Index: music.Label2Index["Sol"]},
	Re2:  {Octave: 2, Index: music.Label2Index["Ré"]},
	La1:  {Octave: 1, Index: music.Label2Index["La"]},
	Mi1:  {Octave: 1, Index: music.Label2Index["Mi"]},
}

type Note struct {
	StringNum StringNumber
	FretNum   FretNumber
}

func (n Note) Frequency() float64 {
	// 1. On récupère la note de la corde à vide
	note := stringNotes[n.StringNum]
	// 2. On ajoute l'intervalle de la frette (nb de demi-tons)
	note.Add(music.Interval(n.FretNum))
	// 3. On calcul la fréquence
	return note.Frequency()
}
