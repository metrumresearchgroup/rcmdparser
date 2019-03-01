package tests

import (
	"github.com/metrumresearchgroup/rcmdparser/rcmdparser"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestParseShinyCheckLog(t *testing.T) {
	inputSlice := GetTestShinyCheckLog(t)

	actualResults := rcmdparser.ParseCheckLog(inputSlice)
	// fmt.Println(actualResults)
	assert.Equal(t, 1, len(actualResults.Notes), "shiny log has one note.")
}

func TestParseShinyTestLog(t *testing.T) {
	inputSlice := GetTestShinyTestLog(t)

	expected := rcmdparser.TestResults{
		Ok: 743,
		Skipped: 0,
		Failed: 0,
	}

	actual := rcmdparser.ParseTestLog(inputSlice)

	//OK: 743 SKIPPED: 0 FAILED: 0
	assert.Equal(t, expected.Ok, actual.Ok )
	assert.Equal(t, expected.Skipped, actual.Skipped)
	assert.Equal(t, expected.Failed, actual.Failed)
}

func TestParseReleasyCheckLog(t *testing.T) {
	inputSlice := GetTestReleasyCheckLog(t)

	expected := rcmdparser.CheckLogEntries{
		Errors: []string{
			`checking whether package ‘releasy’ can be installed ... ERROR
Installation failed.
See ‘/Users/johncarlos/gitspace/releasy/checkplace/releasy.Rcheck/00install.out’ for details.`,
		},
	}

	actual := rcmdparser.ParseCheckLog(inputSlice)

	assert.Equal(t, len(expected.Errors), len(actual.Errors))
	assert.True(t, reflect.DeepEqual(expected.Errors, actual.Errors))

}

//Expect 8 ok, none others
func TestParseRcppTOMLSuccessCheckLog(t *testing.T) {

	inputSlice := GetTestRcppTOMLPassCheckLog(t)

	fixture := rcmdparser.CheckOutputInfo{
		CheckOutputRaw: inputSlice,
	}

	actual := fixture.Parse()

	assert.Equal(t, 8, actual.Tests.Ok)
	assert.Equal(t, 0, actual.Tests.Failed)
	assert.Equal(t, 0, actual.Tests.Skipped)
	assert.Equal(t, 1, actual.Tests.Unknown)

}

func TestParseRcppTOMLFailCheckLog(t *testing.T) {
	inputSlice := GetTestRcppTOMLFailCheckLog(t)

	fixture := rcmdparser.CheckOutputInfo{
		CheckOutputRaw: inputSlice,
	}

	actual := fixture.Parse()

	assert.Equal(t, 7, actual.Tests.Ok)
	assert.Equal(t, 0, actual.Tests.Failed)
	assert.Equal(t, 0, actual.Tests.Skipped)
	assert.Equal(t, 2, actual.Tests.Unknown)
}

func TestParseDataDotTablePassCheckLog(t *testing.T) {
	inputSlice := GetTestDataDotTablePassCheckLog(t)

	fixture := rcmdparser.CheckOutputInfo{
		CheckOutputRaw: inputSlice,
	}

	actual := fixture.Parse()

	assert.Equal(t, 2, actual.Tests.Ok)
	assert.Equal(t, 0, actual.Tests.Failed)
	assert.Equal(t, 0, actual.Tests.Skipped)
	assert.Equal(t, 2, actual.Tests.Unknown)
}

func TestParseDataDotTableFailCheckLog(t *testing.T) {
	inputSlice := GetTestDataDotTableFailCheckLog(t)

	fixture := rcmdparser.CheckOutputInfo{
		CheckOutputRaw: inputSlice,
	}

	actual := fixture.Parse()

	assert.Equal(t, 2, actual.Tests.Ok)
	assert.Equal(t, 1, actual.Tests.Failed)
	assert.Equal(t, 0, actual.Tests.Skipped)
	assert.Equal(t, 1, actual.Tests.Unknown)
}

// Doesn't buy us anything, all Releasy tests pass, a case which is already covered by Shiny.
//func TestParseReleasyTestLog(t * testing.T) {
//	inputSlice := GetTestReleasyTestLog(t)
//
//	expected := TestResults {
//
//	}
//}