package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type options struct {
	InputFile  string `short:"i" long:"input" required:"true" description:"input .PROC file name"`
	OutputFile string `short:"o" long:"output" default:"STDOUT" description:"output dot file name"`
}

func getOptions() (*options, error) {
	options := new(options)
	_, err := flags.Parse(options)
	if err != nil {
		return nil, err
	}
	return options, nil
}

func main() {
	res := process()
	os.Exit(res)
}

func process() int {
	options, err := getOptions()
	if err != nil {
		return 1
	}

	fmt.Println(options)

	nhelp := parseNHELP(options.InputFile)
	fmt.Println(nhelp.routine)

	return 0
}

type Routine int

const (
	BRAIN1 Routine = iota
	BUILD1
	CHANGE
	DIVIDE
	GATHER
	GRAPH1
	INDEX1
	INPUT1
	MATCH1
	MATRIX
	MERGE1
	PRINT1
	PRINT2
	REPORT
	SEARCH
	SORT
	SUMUP1
	UPDATE
)

type Step struct {
	routine Routine // kind of routine like BUILD1, INPUT1, etc.
}

func parseNHELP(inputFile string) Step {
	// no implementation
	return Step{routine: BUILD1}
}
