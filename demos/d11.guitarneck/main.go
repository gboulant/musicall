package main

import (
	"encoding/csv"
	"fmt"
	"os"

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

// demo01_printnames print all notes for all couple (string, fret)
func demo01_printnames() error {
	for _, stringNumber := range stringNumbers {
		for fretNumber := range 8 {
			name := NoteName(stringNumber, guitar.FretNumber(fretNumber))
			fmt.Printf("Note(s=%d, f=%d): %s\n", stringNumber, fretNumber, name)
		}
	}
	return nil
}

// demo02_guitarneck prints the notes in a table emulating the guitar neck
func demo02_guitarneck() error {
	nbfrets := 16
	var header0 string = "   | "
	var header1 string = "----"
	for fretNumber := range nbfrets {
		header0 += fmt.Sprintf("F%.2d   ", fretNumber)
		header1 += "------"
	}
	fmt.Println(header0)
	fmt.Println(header1)

	for _, stringNumber := range stringNumbers {
		var line string = fmt.Sprintf("S%d | ", stringNumber)
		for fretNumber := range nbfrets {
			name := NoteName(stringNumber, guitar.FretNumber(fretNumber))
			line += fmt.Sprintf("%-6s", name)
		}
		fmt.Println(line)
	}
	return nil
}

// demo03_guitarneck_ToCSV does like demo02 but records the result in a
// CSV file. The cell (icol,iline) contains the note of the string
// s=iline for the fret f=icol+1. The header is:
//
// Sn; F0; F1; ...; F16
//
// where the column Sn contains the number of the guitar string.
func demo03_guitarneck_ToCSV() error {
	nbfrets := 16

	// -------------------------------------------
	// STEP 01: creating the records
	var records [][]string = make([][]string, len(stringNumbers)+1)

	// Header line
	records[0] = make([]string, nbfrets+1)
	records[0][0] = "Sn"
	for fretNumber := range nbfrets {
		records[0][fretNumber+1] = fmt.Sprintf("F%.2d", fretNumber)
	}

	// Lines of data
	for _, sn := range stringNumbers {
		records[sn] = make([]string, nbfrets+1)
		records[sn][0] = fmt.Sprintf("S%d", sn)
		for fretNumber := range nbfrets {
			name := NoteName(sn, guitar.FretNumber(fretNumber))
			records[sn][fretNumber+1] = name
		}
	}

	// -------------------------------------------
	// STEP 02: save the records in CSF formated file
	csvpath := "output.guitarneck.csv"
	csvfile, err := os.OpenFile(csvpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	w := csv.NewWriter(csvfile)
	w.Comma = ';'
	for _, record := range records {
		if err := w.Write(record); err != nil {
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

func main() {
	//demo01_printnames()
	demo02_guitarneck()
	demo03_guitarneck_ToCSV()
}
