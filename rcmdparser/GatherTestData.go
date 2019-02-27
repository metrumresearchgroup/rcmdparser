package rcmdparser

import (
	"github.com/spf13/afero"
	"testing"
)

func GetFileAsByteSlice(t *testing.T, filepath string) []byte {
	//var fileSystem = new(afero.MemMapFs)
	var fileSystem = new(afero.OsFs)

	//wd, _ := os.Getwd()
	//fmt.Println(fmt.Sprintf("Working directory: %s", wd ))
	// WD is: /Users/<itme>/go/src/github.com/metrumresearchgroup/rcmdparser/rcmdparser

	content, err := afero.ReadFile(fileSystem, filepath)
	if err != nil {
		//return nil, errors.New(fmt.Sprintf("Error: Could not read file %s", filepath))
		t.Error(err)
	}

	return content

}

func GetTestShinyCheckLog(t *testing.T) []byte {
	content := GetFileAsByteSlice(t,"./testdata/shiny.Rcheck/00check.log")
	return content
}

func GetTestShinyInstallLog(t *testing.T) []byte {
	content := GetFileAsByteSlice(t, "./testdata/shiny.Rcheck/00install.out")
	return content
}

func GetTestShinyTestLog(t *testing.T) []byte {
	content := GetFileAsByteSlice(t,"./testdata/shiny.Rcheck/tests/test-all.Rout")
	return content
}

func GetTestReleasyCheckLog(t *testing.T) []byte {
	content := GetFileAsByteSlice(t,"./testdata/releasy.Rcheck/00check.log")
	return content
}

func GetTestReleasyTestLog(t *testing.T) []byte {
	return GetFileAsByteSlice(t, "./testdata/releasy.Rcheck/tests/test-all.Rout")
}

func GetTestRcppTOMLPassCheckLog(t *testing.T) []byte {
	return GetFileAsByteSlice(t, "./testdata/RcppTOML.Rcheck_pass/00check.log")
}

func GetTestRcppTOMLFailCheckLog(t *testing.T) []byte {
	return GetFileAsByteSlice(t, "./testdata/RcppTOML.Rcheck_fail/00check.log")
}

//Add more readable getters here.