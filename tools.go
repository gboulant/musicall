package musicall

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// LogError must be used to log error on the terminal. Indeed, for a
// more clear code, some functions that manipulate the music data don't
// return an error, even if an error occurs (for example if trying to
// get the note of name "Ré" while it is registered with the name "Re",
// without accent). Instead, we call the LogError function to print the
// error message. By default, the print uses log.Fatal that interupts the
// process (and then you know explicitly the error and can fix it).
var LogError = log.Fatalf

//var LogError = fmt.Printf

// --------------------------------------------------------------------
// The Example structure can be used to defined a set of demonstrative
// examples in a executable program (see the usage in the executable
// programs of the folder cmds, e.g. the program cmds/examples).

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

func StartExampleApp(defaultExampleName string) {
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
