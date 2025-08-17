package main

import (
	"flag"
	"fmt"
	"os"
)

const defaultExampleName string = "D01"

func init() {
	NewExample("D01", "son de quintes", DEMO01_quintes)
	NewExample("D02", "vibrato", DEMO02_vibrato)
	NewExample("D03", "modulation d'amplitude", DEMO03_amplitude_modulation)
	NewExample("D04", "modulation de fréquence", DEMO04_frequency_modulation)
	NewExample("D05", "sounds like a laser", DEMO05_sounds_like_a_laser)
}

func main() {
	var listExamples bool
	var exampleName string
	flag.BoolVar(&listExamples, "l", false, "list of demo examples")
	flag.StringVar(&exampleName, "n", defaultExampleName, "name of the demo example to execute")
	flag.Parse()

	if listExamples {
		ListExamples()
		os.Exit(0)
	}

	example, err := GetExample(exampleName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Executing demo %s ...\n", example.Name)
	if err := example.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
