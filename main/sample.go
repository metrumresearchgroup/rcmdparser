package main

import (
	"bufio"
	"fmt"
	"github.com/metrumresearchgroup/rcmdparser/rcmdparser"
	"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"os"
)

// This function provides a sample of how the rcmdparser might be used.
// Assume that the package check functionality has been run for all of your packages,
// and the results have been saved to ../rcmdparser/testdata
func main() {
	var r io.Reader
	var byteData []byte

	wd, _ := os.Getwd()
	fmt.Sprintf("Working directory: %s", wd)
	file, _ := os.Open("rcmdparser/testdata/testwarningerror.Rcheck/00check.log")
	r = bufio.NewReader(file)
	byteData, _  = ioutil.ReadAll(r)
	output := rcmdparser.ParseTestLog(byteData)
	fmt.Println(fmt.Sprintf("========Output of Parser=========\n%s\n=========END=========", output))

	groupedOutput := rcmdparser.ParseCheckDir(afero.NewMemMapFs(), )

}
