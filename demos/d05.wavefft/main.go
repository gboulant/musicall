package main

import applet "github.com/gboulant/dingo-applet"

const defaultExampleName = "D01"

func init() {
	applet.AddApplet("D01", "Signal à 2 fréquences", D01_fft)
	applet.AddApplet("D02", "Fréquence d'amplitude Max", D02_fft_frequencyOfMaxAmplitude)
	applet.AddApplet("D03", "Augmentation du contraste", D03_fft_smoothboundaries)
}

func main() {
	applet.StartApplication(defaultExampleName)
}
