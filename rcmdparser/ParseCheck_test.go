package rcmdparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseShinyOutput(t *testing.T) {
	inputSlice, err := GetTestShinyCheckLog()
	if(err != nil) {
		t.Error(err)
	}
	actualResults := ParseCheckLog(inputSlice)
	// fmt.Println(actualResults)
	assert.Equal(t, 1, len(actualResults.Notes), "shiny log has one note.")
}