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

// demo01_printnames print all notes for all couple (string, fret)
func demo01_printnames() error {
	var fretNumber guitar.FretNumber
	for _, stringNumber := range stringNumbers {
		for fretNumber = range 8 {
			note := guitar.Note{StringNum: stringNumber, FretNum: fretNumber}
			fmt.Printf("Note(s=%d, f=%d): %s\n", stringNumber, fretNumber, note.Name())
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
			note := guitar.Note{StringNum: stringNumber, FretNum: guitar.FretNumber(fretNumber)}
			line += fmt.Sprintf("%-6s", note.Name())
		}
		fmt.Println(line)
	}
	return nil
}

// makerecords creates a dataset of records prepared to be saved in CSV formated
// file. The data in a record cell is a string created from a Note using the
// function note2string.
func makerecords(nbfrets int, note2string func(n guitar.Note) string) (records [][]string) {
	records = make([][]string, len(stringNumbers)+1)

	// Header line
	records[0] = make([]string, nbfrets+1)
	records[0][0] = "Sn"
	for fretNumber := range nbfrets {
		records[0][fretNumber+1] = fmt.Sprintf("F%.2d", fretNumber)
	}

	// Lines of data
	for _, stringNumber := range stringNumbers {
		records[stringNumber] = make([]string, nbfrets+1)
		records[stringNumber][0] = fmt.Sprintf("S%d", stringNumber)
		for fretNumber := range nbfrets {
			note := guitar.Note{StringNum: stringNumber, FretNum: guitar.FretNumber(fretNumber)}
			records[stringNumber][fretNumber+1] = note2string(note)
		}
	}
	return records
}

// saverecords saves the dataset of records into a CSV formated file
func saverecords(records [][]string, csvpath string) error {
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

// demo03_guitarneck_NameToCSV creates a CSV formated file containing the names
// of the notes for each guitar strings (lines) and frets (columns). The header
// is:
//
// Sn; F0; F1; ...; F16
//
// where the column Sn contains the number of the guitar string (1 for the Mi3
// and 6 for the Mi1).
func demo03_guitarneck_NameToCSV() error {
	nbfrets := 16

	// -------------------------------------------
	// STEP 01: creating the records
	note2string := func(n guitar.Note) string {
		return n.Name()
	}
	records := makerecords(nbfrets, note2string)

	// -------------------------------------------
	// STEP 02: save the records in CSF formated file
	csvpath := "output.guitarneck.names.csv"
	return saverecords(records, csvpath)
}

// demo04_guitarneck_FreqToCSV creates a CSV formated file containing the
// frequencies of the notes for each guitar strings (lines) and frets (columns).
// The header
func demo04_guitarneck_FreqToCSV() error {
	nbfrets := 16

	// -------------------------------------------
	// STEP 01: creating the records
	note2string := func(n guitar.Note) string {
		return fmt.Sprintf("%.1f", n.Frequency())
	}
	records := makerecords(nbfrets, note2string)

	// -------------------------------------------
	// STEP 02: save the records in CSF formated file
	csvpath := "output.guitarneck.frequencies.csv"
	return saverecords(records, csvpath)
}

func main() {
	demo01_printnames()
	demo02_guitarneck()
	demo03_guitarneck_NameToCSV()
	demo04_guitarneck_FreqToCSV()
}
