package main

import (
	"galuma.net/synthetic"
)

const defaultExampleName string = "D01"

func init() {
	synthetic.NewExample("D00", "echelle logarithmique", DEMO00_logscale)
	synthetic.NewExample("D01", "son de quintes", DEMO01_quintes)
	synthetic.NewExample("D02", "vibrato", DEMO02_vibrato)
	synthetic.NewExample("D03", "modulation d'amplitude", DEMO03_amplitude_modulation)
	synthetic.NewExample("D04", "modulation de fréquence", DEMO04_frequency_modulation)
	synthetic.NewExample("D05", "sounds like a laser", DEMO05_sounds_like_a_laser)
}

func main() {
	synthetic.StartExampleApp(defaultExampleName)
}
