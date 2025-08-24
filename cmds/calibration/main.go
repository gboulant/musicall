package main

import (
	"fmt"
	"math"
)

// Timiskhakov uses hard coded values for the frequencies of notes while
// on our side, we calculate the frequency by adding or substracting
// intervals to/from a reference frequency (the La3). See packages
// guitar and music. In this program, we check the difference between
// the hard coded values (from Timiskhakov project) and the computed
// values (from our project).

func TEST00() {
	cf := getCalculatedFrequency(6, 1)
	hf := getHardCodedFrequency(6, 1)
	fmt.Printf("(%d, %d): cf=%.2f Hz hf=%.2f Hz\n", 6, 1, cf, hf)
}

// sequenceOfInt generates a list of integer starting from start and
// ending with end (included)
func sequenceOfInt(start int, end int) []int {
	size := end - start + 1
	s := make([]int, int(size))
	for i := range size {
		s[i] = start + i
	}
	return s
}

func almostEqual(a, b float64, accuracy float64) bool {
	return math.Abs(a-b) < accuracy
}

func TEST01() {
	stringNumbers := sequenceOfInt(1, 6)
	fretNumbers := sequenceOfInt(0, 19)

	accuracy := 1e-2

	msgpattern := "(%d, %2d): cf = %6.2f Hz  hf = %6.2f Hz  - %s\n"
	for _, stringNum := range stringNumbers {
		for _, fretNum := range fretNumbers {
			cf := getCalculatedFrequency(stringNum, fretNum)
			hf := getHardCodedFrequency(stringNum, fretNum)
			var msg string
			if !almostEqual(cf, hf, accuracy) {
				msg = "DIFF"
			} else {
				msg = "OK"
			}
			fmt.Printf(msgpattern, stringNum, fretNum, cf, hf, msg)
		}
	}
}

func main() {
	//TEST00()
	TEST01()
}
