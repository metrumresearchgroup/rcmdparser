package rcmdparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseShinyCheckLog(t *testing.T) {
	inputSlice, err := GetTestShinyCheckLog()
	if(err != nil) {
		t.Error(err)
	}
	actualResults := ParseCheckLog(inputSlice)
	// fmt.Println(actualResults)
	assert.Equal(t, 1, len(actualResults.Notes), "shiny log has one note.")
}

func TestParseShinyTestLog(t *testing.T) {
	inputSlice, err := GetTestShinyTestLog()
	if(err != nil) {
		t.Error(err)
	}

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