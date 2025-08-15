package main

import (
	"fmt"
)

// Example defines a data structure to describe a demonstration
// procedure and then execute this use case. Set the Execute attribute
// to the function to execute.
type Example struct {
	Name    string
	Execute func() error
	Comment string
}

func (u Example) String() string {
	return fmt.Sprintf("%-14s (%s)", u.Name, u.Comment)
}

var examples []Example

// NewExample creates a new Example program and registers the created
// program into the catalog of programs. After this registration, the
// example program can be obtain by name using the function GetExample
func NewExample(name string, comment string, function func() error) *Example {
	p := Example{name, function, comment}
	examples = append(examples, p)
	return &p
}

func GetExample(name string) (*Example, error) {
	for _, example := range examples {
		if example.Name == name {
			return &example, nil
		}
	}
	return nil, fmt.Errorf("no example program with name %s", name)
}

func ListExamples() {
	for _, example := range examples {
		fmt.Println(example)
	}
}
