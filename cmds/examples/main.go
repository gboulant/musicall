package main

import (
	"flag"
	"fmt"
	"os"
)

const defaultProgName string = "D01"

func main() {
	var listPrograms bool
	var programName string
	flag.BoolVar(&listPrograms, "l", false, "list of demo examples")
	flag.StringVar(&programName, "n", defaultProgName, "name of the demo to execute")
	flag.Parse()

	if listPrograms {
		for _, program := range programs {
			fmt.Println(program)
		}
		os.Exit(0)
	}

	program, err := GetProgram(programName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Executing demo %s ...\n", program.Name)
	if err := program.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
