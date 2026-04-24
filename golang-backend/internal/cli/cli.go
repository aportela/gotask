package cli

import (
	"flag"
)

type CommandLineParams struct {
	Verbose        bool
	InsertBulkData bool
}

func HandleFlags() (CommandLineParams, error) {
	params := CommandLineParams{}
	flag.BoolVar(&params.Verbose, "verbose", false, "Displays verbose output")
	flag.BoolVar(&params.InsertBulkData, "demo", false, "Insert demo/bulk data")

	flag.Parse()

	return params, nil
}
