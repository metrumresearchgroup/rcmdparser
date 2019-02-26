package rcmdparser

import (
	"errors"
	"fmt"
	"github.com/spf13/afero"
)

func GetFileAsByteSlice(filepath string) ([]byte, error) {
	//var fileSystem = new(afero.MemMapFs)
	var fileSystem = new(afero.OsFs)

	//wd, _ := os.Getwd()
	//fmt.Println(fmt.Sprintf("Working directory: %s", wd ))
	// WD is: /Users/<itme>/go/src/github.com/metrumresearchgroup/rcmdparser/rcmdparser

	content, err := afero.ReadFile(fileSystem, filepath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: Could not read file %s", filepath))
	}

	return content, nil

}

func GetTestShinyCheckLog() ([]byte, error) {
	content, err := GetFileAsByteSlice("./testdata/shiny.Rcheck/00check.log")
	return content, err
}

func GetTestShinyInstallLog() ([]byte, error) {
	content, err := GetFileAsByteSlice("./testdata/shiny.Rcheck/00install.out")
	return content, err
}

func GetTestShinyTestLog() ([]byte, error) {
	content, err := GetFileAsByteSlice("./testdata/shiny.Rcheck/tests/test-all.Rout")
	return content, err
}