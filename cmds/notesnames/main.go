package main

import (
	"fmt"

	"github.com/gboulant/musicall/guitar"
)

var stringNumbers = []guitar.StringNumber{
	guitar.Mi3,
	guitar.Si2,
	guitar.Sol2,
	guitar.Re2,
	guitar.La1,
	guitar.Mi1,
}

func NoteName(stringNumber guitar.StringNumber, fretNumber guitar.FretNumber) string {
	note := guitar.Note{StringNum: stringNumber, FretNum: fretNumber}
	return note.Name()
}

func demo01_printnames() {
	for _, stringNumber := range stringNumbers {
		for fretNumber := range 8 {
			name := NoteName(stringNumber, guitar.FretNumber(fretNumber))
			fmt.Printf("Note(s=%d, f=%d): %s\n", stringNumber, fretNumber, name)
		}
	}
}

func demo02_guitarmanche() {
	nbfrets := 16
	var header0 string = "String \\ Fret | "
	var header1 string = "----------------"
	for fretNumber := range nbfrets {
		header0 += fmt.Sprintf("F%.2d   ", fretNumber)
		header1 += "------"
	}
	fmt.Println(header0)
	fmt.Println(header1)

	for _, stringNumber := range stringNumbers {
		var line string = fmt.Sprintf("S%d %-5s      | ", stringNumber, NoteName(stringNumber, 0))
		for fretNumber := range nbfrets {
			name := NoteName(stringNumber, guitar.FretNumber(fretNumber))
			line += fmt.Sprintf("%-6s", name)
		}
		fmt.Println(line)
	}

}

func main() {
	//demo01_printnames()
	demo02_guitarmanche()
}
