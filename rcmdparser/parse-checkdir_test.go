package rcmdparser

import (
	"fmt"
	"github.com/dpastoor/goutils"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadCheckDir(t *testing.T) {
	osFs := afero.NewOsFs()
	testFS := afero.NewMemMapFs()
	// tests dir but no testthat
	_ = testFS.MkdirAll("noTestThat/tests", 0755)
	_ = goutils.WriteLinesFS(testFS, []string{"log"}, "noTestThat/00check.log")
	_ = goutils.WriteLinesFS(testFS, []string{"install"}, "noTestThat/00install.out")

	// with testthat
	_ = testFS.MkdirAll("WithTestThat/tests", 0755)
	_ = goutils.WriteLinesFS(testFS, []string{"log"}, "WithTestThat/00check.log")
	_ = goutils.WriteLinesFS(testFS, []string{"install"}, "WithTestThat/00install.out")
	_ = goutils.WriteLinesFS(testFS, []string{"library(testthat)"}, "WithTestThat/tests/testthat.Rout")

	// Failed UsesTestthat
	_ = testFS.MkdirAll("FailedTest/tests", 0755)
	_ = goutils.WriteLinesFS(testFS, []string{"log"}, "FailedTest/00check.log")
	_ = goutils.WriteLinesFS(testFS, []string{"install"}, "FailedTest/00install.out")
	_ = goutils.WriteLinesFS(testFS, []string{"library(testthat)"}, "FailedTest/tests/testthat.Rout.fail")
	
	var cdtests = []struct {
		fs       afero.Fs      
		in       string
		expected CheckOutputInfo
	}{
		{
			testFS,
			"noTestThat",
			CheckOutputInfo{
				TestInfo{true, false, nil},
				[]byte("log\n"),
				[]byte("install\n"),
			},
		},
		{
			testFS,
			"WithTestThat",
			CheckOutputInfo{
				TestInfo{true, true, []byte("library(testthat)\n")},
				[]byte("log\n"),
				[]byte("install\n"),
			},
		},
		{
			testFS,
			"FailedTest",
			CheckOutputInfo{
				TestInfo{true, true, []byte("library(testthat)\n")},
				[]byte("log\n"),
				[]byte("install\n"),
			},
		},
		{
			osFs,
			"./testdata/shiny.Rcheck",
			CheckOutputInfo{
				TestInfo{true, true, nil},
				nil,
				nil,
			},
		},
		{
			osFs,
			"./testdata/releasy.Rcheck",
			CheckOutputInfo{
				TestInfo{false, false, nil},
				nil,
				nil,
			},
		},
		{
			osFs,
			"./testdata/testerror.Rcheck",
			CheckOutputInfo{
				TestInfo{true, true, nil},
				nil,
				nil,
			},
		},
	}

	for _, tt := range cdtests {
		actual, _ := parseCheckDir(tt.fs, tt.in)
		if len(tt.expected.CheckOutputRaw) <= 0 {
			actual.CheckOutputRaw = tt.expected.CheckOutputRaw
		}
		if len(tt.expected.InstallOutputRaw) <= 0 {
			actual.InstallOutputRaw = tt.expected.InstallOutputRaw
		}
		if len(tt.expected.TestInfo.ResultsFile) <= 0  {
			actual.TestInfo.ResultsFile = tt.expected.TestInfo.ResultsFile
		}
		assert.Equal(t, actual, tt.expected, fmt.Sprintf("Not equal: %s", tt.in))
	}
}
