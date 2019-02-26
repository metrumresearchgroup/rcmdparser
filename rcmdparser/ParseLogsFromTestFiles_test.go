package rcmdparser

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestParseShinyCheckLog(t *testing.T) {
	inputSlice := GetTestShinyCheckLog(t)

	actualResults := ParseCheckLog(inputSlice)
	// fmt.Println(actualResults)
	assert.Equal(t, 1, len(actualResults.Notes), "shiny log has one note.")
}

func TestParseShinyTestLog(t *testing.T) {
	inputSlice := GetTestShinyTestLog(t)

	expected := TestResults {
		Ok: 743,
		Skipped: 0,
		Failed: 0,
	}

	actual := ParseTestLog(inputSlice)

	//OK: 743 SKIPPED: 0 FAILED: 0
	assert.Equal(t, expected.Ok, actual.Ok )
	assert.Equal(t, expected.Skipped, actual.Skipped)
	assert.Equal(t, expected.Failed, actual.Failed)
}

func TestParseReleasyCheckLog(t *testing.T) {
	inputSlice := GetTestReleasyCheckLog(t)

	expected := LogEntries {
		Errors: []string{
			`checking whether package ‘releasy’ can be installed ... ERROR
Installation failed.
See ‘/Users/johncarlos/gitspace/releasy/checkplace/releasy.Rcheck/00install.out’ for details.`,
		},
	}

	actual := ParseCheckLog(inputSlice)

	assert.Equal(t, len(expected.Errors), len(actual.Errors))
	assert.True(t, reflect.DeepEqual(expected.Errors, actual.Errors))

}

// Doesn't buy us anything, all Releasy tests pass, a case which is already covered by Shiny.
//func TestParseReleasyTestLog(t * testing.T) {
//	inputSlice := GetTestReleasyTestLog(t)
//
//	expected := TestResults {
//
//	}
//}