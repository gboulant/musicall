package main

import "github.com/gboulant/musicall"

const defaultExampleName string = "D01"

func init() {
	musicall.NewExample("T01", "Play all open strings", T01_play_open_strings)
	musicall.NewExample("T02", "Play the main chords", T02_main_chords)
	musicall.NewExample("T03", "Play the pentatonic scale from La", T03_pentatonic_scale_La)

	musicall.NewExample("D01", "Nocking on the heaven's door", D01_Nocking_on_the_heavens_door)
	musicall.NewExample("D02", "U2, One", D02_U2_One)
	musicall.NewExample("D03", "ACDC, Thunderstruck", D03_ACDC_Thunderstruck)
	musicall.NewExample("D04", "NoirDesir, Tostaky", D04_NoirDesir_Tostaky)
	musicall.NewExample("D05", "NoirDesir, Un jour en France", D05_NoirDesir_Un_jour_en_France)
	musicall.NewExample("D06", "NoirDesir, Le vent l'emportera (pont)", D06_NoirDesir_Le_vent_l_emportera)
	musicall.NewExample("D07", "U2, Bloody Sunday", D07_U2_Bloody_Sunday)
	musicall.NewExample("D08", "Rythm Bas & Bas - Haut & Haut - Bas", D08_Rythm_UpDown)
	musicall.NewExample("D09", "Johnny Cash, Hurt", D09_Johnny_Cash_Hurt)
}

func main() {
	musicall.StartExampleApp(defaultExampleName)
}
