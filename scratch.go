package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/metrumresearchgroup/rcmdparser/rcmdparser"
	"github.com/spf13/afero"
)

// This function provides a sample of how the rcmdparser might be used.
// Assume that the package check functionality has been run for all of your packages,
// and the results have been saved to ../rcmdparser/testdata
func main() {

	fs := afero.NewOsFs()
	output, err := rcmdparser.NewCheck(fs, "rcmdparser/testdata/testwarning.Rcheck")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	prettyPrint(output)
}

func prettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
