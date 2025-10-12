package main

import "fmt"

/*
ASCII codes:

From 32  to 47  : special characters (space, !, ", ..., /)
From 48  to 57  : numbers characters (0, 1, 3, ...9)
From 58  to 64  : special characters (:, ;, ..., @)
From 65  to 90  : upper case letters (A, B, C, ..., Z)
From 91  to 96  : special characters ([, \, ..., `)
From 97  to 122 : lower case letters (a, b, c, ..., z)
From 123 to 126 : special characters ({, |, ..., DEL)
*/
var numbercodes []int
var lowercasecodes []int
var uppercasecodes []int

var specialchars []rune = []rune{',', ';', ':', '!', '.'}
var specialcodes []int

const spacecode = 32

func init() {
	lowercasecodes = sequenceOfInt(97, 122)
	uppercasecodes = sequenceOfInt(65, 90)
	numbercodes = sequenceOfInt(48, 57)
	specialcodes = make([]int, len(specialchars))
	for i, char := range specialchars {
		specialcodes[i] = char2code(char)
	}
}

func info() {
	fmt.Println(lowercasecodes)
	fmt.Println(uppercasecodes)
	fmt.Println(numbercodes)
	fmt.Println(specialcodes)
}

func char2code(character rune) int {
	return int(character)
}

func code2char(code int) rune {
	return rune(code)
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
