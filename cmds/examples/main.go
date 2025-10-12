package main

import applet "github.com/gboulant/dingo-applet"

const defaultExampleName string = "D01"

func init() {
	applet.NewExample("D00", "echelle logarithmique", DEMO00_logscale)
	applet.NewExample("D01", "son de quintes", DEMO01_quintes)
	applet.NewExample("D02", "vibrato", DEMO02_vibrato)
	applet.NewExample("D03", "modulation d'amplitude", DEMO03_amplitude_modulation)
	applet.NewExample("D04", "modulation de fréquence", DEMO04_frequency_modulation)
	applet.NewExample("D05", "sounds like a laser", DEMO05_sounds_like_a_laser)
	applet.NewExample("D06", "echelle musicale", DEMO06_musicalscale)
}

func main() {
	applet.StartExampleApp(defaultExampleName)
}
