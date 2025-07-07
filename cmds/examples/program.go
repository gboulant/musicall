package main

import (
	"fmt"
)

type Program struct {
	Name    string
	Execute func() error
	Comment string
}

func (u Program) String() string {
	return fmt.Sprintf("%-14s (%s)", u.Name, u.Comment)
}

var programs []Program

func AddProgram(name string, execute func() error, comment string) {
	programs = append(programs, Program{name, execute, comment})
}

func GetProgram(name string) (Program, error) {
	for _, program := range programs {
		if program.Name == name {
			return program, nil
		}
	}
	return Program{}, fmt.Errorf("no program with name %s", name)
}
