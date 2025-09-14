package main

import "github.com/gboulant/musicall"

const defaultExampleName string = "D01"

func init() {
	musicall.NewExample("D00", "echelle logarithmique", DEMO00_logscale)
	musicall.NewExample("D01", "son de quintes", DEMO01_quintes)
	musicall.NewExample("D02", "vibrato", DEMO02_vibrato)
	musicall.NewExample("D03", "modulation d'amplitude", DEMO03_amplitude_modulation)
	musicall.NewExample("D04", "modulation de fréquence", DEMO04_frequency_modulation)
	musicall.NewExample("D05", "sounds like a laser", DEMO05_sounds_like_a_laser)
	musicall.NewExample("D06", "echelle musicale", DEMO06_musicalscale)
}

func main() {
	musicall.StartExampleApp(defaultExampleName)
}
