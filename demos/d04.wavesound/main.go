package main

import applet "github.com/gboulant/dingo-applet"

const defaultExampleName string = "D01"

func init() {
	applet.AddApplet("D00", "echelle logarithmique", DEMO00_logscale)
	applet.AddApplet("D01", "son de quintes", DEMO01_quintes)
	applet.AddApplet("D02", "vibrato", DEMO02_vibrato)
	applet.AddApplet("D03", "modulation d'amplitude", DEMO03_amplitude_modulation)
	applet.AddApplet("D04", "modulation de fréquence", DEMO04_frequency_modulation)
	applet.AddApplet("D05", "sounds like a laser", DEMO05_sounds_like_a_laser)
	applet.AddApplet("D06", "echelle musicale", DEMO06_musicalscale)
	applet.AddApplet("D07", "filtre sigmoide", DEMO07_sigmoidfilter)
	applet.AddApplet("D08", "sequence de signaux adoucis", DEMO08_sequence_smoot_signal)
}

func main() {
	applet.StartApplication(defaultExampleName)
}
