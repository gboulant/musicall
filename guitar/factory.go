package guitar

type StringNumber int

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

type Note struct {
	StringNum int
	FretNum   int
}

func (n Note) Frequency() float64 {
	// 1. On récupère la fréquence de la corde à vide
	// 2. On dérive ajoute à cette fréquence l'intervalle de la frette
	return 0.
}
