package rcmdparser

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func getFileContents(t *testing.T, filepath string) []byte {
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
func TestParseShinyCheckLog(t *testing.T) {

	inputSlice := getFileContents(t, "testdata/shiny.Rcheck/00check.log")
	actualResults := ParseCheckLog(inputSlice)
	// fmt.Println(actualResults)
	assert.Equal(t, 1, len(actualResults.Notes), "shiny log has one note.")
}

func TestParseTestLogs(t *testing.T) {
	type test struct {
		Input    string
		Expected TestResults
		Context  string
	}
	tests := []test{
		test{
			Input: "./testdata/shiny.Rcheck/tests/test-all.Rout",
			Expected: TestResults{
				Ok:      743,
				Skipped: 10,
				Failed:  0,
			},
			Context: "shiny tests",
		},
	}

	for _, tst := range tests {
		inputSlice := getFileContents(t, tst.Input)
		actual := ParseTestLog(inputSlice)
		assert.Equal(t, tst.Expected.Ok, actual.Ok, fmt.Sprintf("%s, ok", tst.Context))
		assert.Equal(t, tst.Expected.Skipped, actual.Skipped, "shiny test log skipped")
		assert.Equal(t, tst.Expected.Failed, actual.Failed, "shiny test log failures")
	}

	// not doing single check as not setting the expected Output to
	// the string of results for simplicity
}

// func TestParseReleasyCheckLog(t *testing.T) {
// 	inputSlice := GetTestReleasyCheckLog(t)

// 	expected := CheckLogEntries{
// 		Errors: []string{
// 			`checking whether package ‘releasy’ can be installed ... ERROR
// Installation failed.
// See ‘/Users/johncarlos/gitspace/releasy/checkplace/releasy.Rcheck/00install.out’ for details.`,
// 		},
// 	}

// 	actual := ParseCheckLog(inputSlice)

// 	assert.Equal(t, len(expected.Errors), len(actual.Errors))
// 	assert.True(t, reflect.DeepEqual(expected.Errors, actual.Errors))

// }

// //Expect 8 ok, none others
// func TestParseRcppTOMLSuccessCheckLog(t *testing.T) {

// 	inputSlice := GetTestRcppTOMLPassCheckLog(t)

// 	fixture := CheckOutputInfo{
// 		CheckOutputRaw: inputSlice,
// 	}

// 	actual := fixture.Parse()

// 	assert.Equal(t, 8, actual.Tests.Ok)
// 	assert.Equal(t, 0, actual.Tests.Failed)
// 	assert.Equal(t, 0, actual.Tests.Skipped)
// 	assert.Equal(t, 1, actual.Tests.Unknown)

// }

// func TestParseRcppTOMLFailCheckLog(t *testing.T) {
// 	inputSlice := GetTestRcppTOMLFailCheckLog(t)

// 	fixture := CheckOutputInfo{
// 		CheckOutputRaw: inputSlice,
// 	}

// 	actual := fixture.Parse()

// 	assert.Equal(t, 7, actual.Tests.Ok)
// 	assert.Equal(t, 0, actual.Tests.Failed)
// 	assert.Equal(t, 0, actual.Tests.Skipped)
// 	assert.Equal(t, 2, actual.Tests.Unknown)
// }

// func TestParseDataDotTablePassCheckLog(t *testing.T) {
// 	inputSlice := GetTestDataDotTablePassCheckLog(t)

// 	fixture := CheckOutputInfo{
// 		CheckOutputRaw: inputSlice,
// 	}

// 	actual := fixture.Parse()

// 	assert.Equal(t, 2, actual.Tests.Ok)
// 	assert.Equal(t, 0, actual.Tests.Failed)
// 	assert.Equal(t, 0, actual.Tests.Skipped)
// 	assert.Equal(t, 2, actual.Tests.Unknown)
// }

// func TestParseDataDotTableFailCheckLog(t *testing.T) {
// 	inputSlice := GetTestDataDotTableFailCheckLog(t)

// 	fixture := CheckOutputInfo{
// 		CheckOutputRaw: inputSlice,
// 	}

// 	actual := fixture.Parse()

// 	assert.Equal(t, 2, actual.Tests.Ok)
// 	assert.Equal(t, 1, actual.Tests.Failed)
// 	assert.Equal(t, 0, actual.Tests.Skipped)
// 	assert.Equal(t, 1, actual.Tests.Unknown)
// }

// // Doesn't buy us anything, all Releasy tests pass, a case which is already covered by Shiny.
// //func TestParseReleasyTestLog(t * testing.T) {
// //	inputSlice := GetTestReleasyTestLog(t)
// //
// //	expected := TestResults {
// //
// //	}
// //}
