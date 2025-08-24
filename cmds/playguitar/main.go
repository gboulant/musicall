package main

import (
	"galuma.net/synthetic"
)

const defaultExampleName string = "D01"

func init() {
	synthetic.NewExample("T01", "Play all open strings", T01_play_open_strings)
	synthetic.NewExample("T02", "Play the main chords", T02_main_chords)
	synthetic.NewExample("T03", "Play the pentatonic scale from La", T03_pentatonic_scale_La)

	synthetic.NewExample("D01", "Nocking on the heaven's door", D01_Nocking_on_the_heavens_door)
	synthetic.NewExample("D02", "U2, One", D02_U2_One)
	synthetic.NewExample("D03", "ACDC, Thunderstruck", D03_ACDC_Thunderstruck)
	synthetic.NewExample("D04", "NoirDesir, Tostaky", D04_NoirDesir_Tostaky)
	synthetic.NewExample("D05", "NoirDesir, Un jour en France", D05_NoirDesir_Un_jour_en_France)
	synthetic.NewExample("D06", "NoirDesir, Le vent l'emportera (pont)", D06_NoirDesir_Le_vent_l_emportera)
	synthetic.NewExample("D07", "U2, Bloody Sunday", D07_U2_Bloody_Sunday)
	synthetic.NewExample("D08", "Rythm Bas & Bas - Haut & Haut - Bas", D08_Rythm_UpDown)
	synthetic.NewExample("D09", "Johnny Cash, Hurt", D09_Johnny_Cash_Hurt)
}

func main() {
	synthetic.StartExampleApp(defaultExampleName)
}
