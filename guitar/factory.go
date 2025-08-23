package guitar

// ----------------------------------------------------------------------
// Définition des accord principaux
type Notes []Note

var StandardChords map[string]Notes = map[string]Notes{
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
}
