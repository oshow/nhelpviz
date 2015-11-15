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
	return options, err
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
	return 0
}
