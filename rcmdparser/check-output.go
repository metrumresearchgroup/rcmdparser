package rcmdparser

import (
	"github.com/spf13/afero"
)

// NewCheck creates a new CheckOutputRaw Object
// fs and check directory
func NewCheck(fs afero.Fs, cd string) (CheckResults, error) {
	cr, err := ParseCheckDir(fs, cd)
	if err != nil {
		return CheckResults{}, err
	}
	return cr.Parse(), nil
}

// Parse output to LogResults
func (c CheckOutputInfo) Parse() CheckResults {
	lr := CheckResults{
		Checks: ParseCheckLog(c.CheckOutputRaw),
	}
	if c.TestInfo.UsesTestthat {
		lr.Tests = ParseTestLog(c.TestInfo.ResultsFile)
	} else {
		//Assume that the test information is in the check log and pull it from there.
		//TODO: This case may not be as generalizable.
		lr.Tests = ParseTestsFromCheckLog(c.CheckOutputRaw)
	}
	return lr
}
