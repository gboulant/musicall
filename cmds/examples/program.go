package main

import (
	"fmt"
)

// Program defines a data structure to describe a demonstration use case and
// then execute this use case. Set the Execute attribute to the function to
// execute.
type Program struct {
	Name    string
	Execute func() error
	Comment string
}

func (u Program) String() string {
	return fmt.Sprintf("%-14s (%s)", u.Name, u.Comment)
}

var programs []Program

// NewProgram creates a new Program and registers the created program into the
// catalog of program (like AddPrograme)
func NewProgram(name string, comment string, function func() error) *Program {
	p := Program{name, function, comment}
	programs = append(programs, p)
	return &p
}

func GetProgram(name string) (*Program, error) {
	for _, program := range programs {
		if program.Name == name {
			return &program, nil
		}
	}
	return nil, fmt.Errorf("no program with name %s", name)
}
