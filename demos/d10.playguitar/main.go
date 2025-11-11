package main

import applet "github.com/gboulant/dingo-applet"

const defaultExampleName string = "D01"

func init() {
	applet.AddApplet("T01", "Play all open strings", T01_play_open_strings)
	applet.AddApplet("T02", "Play the main chords", T02_main_chords)
	applet.AddApplet("T03", "Play the pentatonic scale from La", T03_pentatonic_scale_La)

	applet.AddApplet("D01", "Nocking on the heaven's door", D01_Nocking_on_the_heavens_door)
	applet.AddApplet("D02", "U2, One", D02_U2_One)
	applet.AddApplet("D03", "ACDC, Thunderstruck", D03_ACDC_Thunderstruck)
	applet.AddApplet("D04", "NoirDesir, Tostaky", D04_NoirDesir_Tostaky)
	applet.AddApplet("D05", "NoirDesir, Un jour en France", D05_NoirDesir_Un_jour_en_France)
	applet.AddApplet("D06", "NoirDesir, Le vent l'emportera (pont)", D06_NoirDesir_Le_vent_l_emportera)
	applet.AddApplet("D07", "U2, Bloody Sunday", D07_U2_Bloody_Sunday)
	applet.AddApplet("D08", "Rythm Bas & Bas - Haut & Haut - Bas", D08_Rythm_UpDown)
	applet.AddApplet("D09", "Johnny Cash, Hurt", D09_Johnny_Cash_Hurt)
}

func main() {
	applet.StartApplication(defaultExampleName)
}
