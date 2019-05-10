package rcmdparser

import (
	"fmt"
	"testing"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestParseTestLogs(t *testing.T) {
	type test struct {
		Input    string
		Expected TestResults
		Context  string
	}
	tests := []test{

		// todo: shouldn't this one expect Failed:1 ?
		test{
			Input: "./testdata/testwarningerror.Rcheck/tests/testthat.Rout.fail",
			Expected: TestResults{
				Ok:      0,
				Skipped: 0,
				Failed:  0,
			},
			Context: "testwarningerror tests",
		},

		test{
			Input: "./testdata/shiny.Rcheck/tests/test-all.Rout",
			Expected: TestResults{
				Ok:      743,
				Skipped: 0,
				Failed:  0,
			},
			Context: "shiny tests",
		},

		test{
			Input: "./testdata/test1.Rcheck/tests/testthat.Rout",
			Expected: TestResults{
				Ok:      1,
				Skipped: 0,
				Failed:  0,
			},
			Context: "test1 tests",
		},

		test{
			Input: "./testdata/testerror.Rcheck/tests/testthat.Rout.fail",
			Expected: TestResults{
				Ok:      1,
				Skipped: 0,
				Failed:  1,
			},
			Context: "testerror tests",
		},

		test{
			Input: "./testdata/testnote.Rcheck/tests/testthat.Rout",
			Expected: TestResults{
				Ok:      1,
				Skipped: 0,
				Failed:  0,
			},
			Context: "testnote tests",
		},

	}

	for _, tst := range tests {
		inputSlice := getFileContents(t, tst.Input)
		actual := parseTestLog(inputSlice)
		assert.Equal(t, tst.Expected.Ok, actual.Ok, fmt.Sprintf("%s, ok", tst.Context))
		assert.Equal(t, tst.Expected.Skipped, actual.Skipped, fmt.Sprintf("%s, skipped", tst.Context))
		assert.Equal(t, tst.Expected.Failed, actual.Failed, fmt.Sprintf("%s, failures", tst.Context))
	}

	// not doing single check as not setting the expected Output to
	// the string of results for simplicity
}

func TestParseCheckLogs(t *testing.T) {
	type test struct {
		Input    string
		Expected CheckLogEntries
		Context  string
	}
	tests := []test{

		test{
			Input: "./testdata/data.table.Rcheck_fail/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string{"checking tests ... ERROR\n  Running ‘autoprint.R’\n  Comparing ‘autoprint.Rout’ to ‘autoprint.Rout.save’ ... OK\n  Running ‘froll.R’\n  Running ‘knitr.R’\n  Comparing ‘knitr.Rout’ to ‘knitr.Rout.save’ ... OK\n  Running ‘main.R’ [46s/27s]\nRunning the tests in ‘tests/main.R’ failed.\nLast 13 lines of output:\n   9: 1253 0.667   485\n  10: 1739 0.636     5\n  \n  Error in eval(exprs[i], envir) : \n    1 error out of 7494 in 26.4s elapsed (42.3s cpu) on Fri Mar  1 16:04:12 2019. [endian==little, sizeof(long double)==16, sizeof(pointer)==8, TZ=America/New_York, locale='C/en_US.UTF-8/en_US.UTF-8/C/en_US.UTF-8/en_US.UTF-8', l10n_info()='MBCS=TRUE; UTF-8=TRUE; Latin-1=FALSE']. Search inst/tests/tests.Rraw for test number 1702.4.\n  Calls: test.data.table -> sys.source -> eval -> eval\n  In addition: Warning messages:\n  1: In fread(testDir(\"isoweek_test.csv\")) :\n    Stopped early on line 6. Expected 2 fields but found 1. Consider fill=TRUE and comment.char=. First discarded non-empty line: <<\"This is narry a date!\">>\n  2: In doTryCatch(return(expr), name, parentenv, handler) :\n    unable to load shared object '/Library/Frameworks/R.framework/Resources/modules//R_X11.so':\n    dlopen(/Library/Frameworks/R.framework/Resources/modules//R_X11.so, 6): Library not loaded: /opt/X11/lib/libSM.6.dylib\n    Referenced from: /Library/Frameworks/R.framework/Resources/modules//R_X11.so\n    Reason: image not found\n  Execution halted"},
				Warnings:    []string(nil),
				Notes:       []string{"checking package dependencies ... NOTE\nPackages suggested but not available for checking:\n  ‘bit64’ ‘R.utils’ ‘xts’ ‘nanotime’ ‘zoo’", "checking Rd cross-references ... NOTE\nPackage unavailable to check Rd xrefs: ‘bit64’"},
			},
			Context: "data.table.Rcheck_fail checklog",
		},

		test{
			Input: "./testdata/data.table.Rcheck_pass/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string(nil),
				Warnings:    []string(nil),
				Notes:       []string{"checking package dependencies ... NOTE\nPackages suggested but not available for checking:\n  ‘bit64’ ‘R.utils’ ‘xts’ ‘nanotime’ ‘zoo’", "checking Rd cross-references ... NOTE\nPackage unavailable to check Rd xrefs: ‘bit64’"},
			},
			Context: "data.table.Rcheck_pass checklog",
		},

		test{
			Input: "./testdata/RcppTOML.Rcheck_fail/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string(nil),
				Warnings:    []string(nil),
				Notes:       []string(nil),
			},
			Context: "RcppTOML.Rcheck_fail checklog",
		},

		test{
			Input: "./testdata/RcppTOML.Rcheck_pass/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string(nil),
				Warnings:    []string(nil),
				Notes:       []string(nil),
			},
			Context: "RcppTOML.Rcheck_pass checklog",
		},

		test{
			Input: "./testdata/releasy.Rcheck/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string{"checking whether package ‘releasy’ can be installed ... ERROR\nInstallation failed.\nSee ‘/Users/johncarlos/gitspace/releasy/checkplace/releasy.Rcheck/00install.out’ for details."},
				Warnings:    []string(nil),
				Notes:       []string(nil),
			},
			Context: "releasy checklog",
		},

		test{
			Input: "./testdata/shiny.Rcheck/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string(nil),
				Warnings:    []string(nil),
				Notes:       []string{"checking installed package size ... NOTE\n  installed size is 10.9Mb\n  sub-directories of 1Mb or more:\n    R     2.0Mb\n    www   7.9Mb"},
			},
			Context: "shiny checklog",
		},

		test{
			Input: "./testdata/test1.Rcheck/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string(nil),
				Warnings:    []string(nil),
				Notes:       []string(nil),
			},
			Context: "test1 checklog",
		},

		test{
			Input: "./testdata/testerror.Rcheck/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string{"checking tests ... ERROR\n  Running ‘testthat.R’\nRunning the tests in ‘tests/testthat.R’ failed.\nLast 13 lines of output:\n  > library(testthat)\n  > library(testerror)\n  > \n  > test_check(\"testerror\")\n  ── 1. Failure: median shows up as failure (@test-median.R#11)  ─────────────────\n  my_median(vals) not equal to 0.\n  1/1 mismatches\n  [1] 4 - 0 == 4\n  \n  ══ testthat results  ═══════════════════════════════════════════════════════════\n  OK: 1 SKIPPED: 0 FAILED: 1\n  1. Failure: median shows up as failure (@test-median.R#11) \n  \n  Error: testthat unit tests failed\n  Execution halted"},
				Warnings:    []string(nil),
				Notes:       []string(nil),
			},
			Context: "testerror checklog",
		},

		test{
			Input: "./testdata/testnote.Rcheck/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string(nil),
				Warnings:    []string(nil),
				Notes:       []string(nil),
			},
			Context: "testnote checklog",
		},

		test{
			Input: "./testdata/testwarning.Rcheck/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string(nil),
				Warnings:    []string{"checking for code/documentation mismatches ... WARNING\nCodoc mismatches from documentation object 'my_median':\nmy_median\n  Code: function(x, ...)\n  Docs: function(...)\n  Argument names in code not in docs:\n    x\n  Mismatches in argument names:\n    Position: 1 Code: x Docs: ..."},
				Notes:       []string(nil),
			},
			Context: "testwarning checklog",
		},

		test{
			Input: "./testdata/testwarningerror.Rcheck/00check.log",
			Expected: CheckLogEntries{
				Errors:      []string{"checking tests ... ERROR\n  Running ‘testthat.R’\nRunning the tests in ‘tests/testthat.R’ failed.\nComplete output:\n  > library(testthat)\n  > library(testdoc)\n  Error in library(testdoc) : there is no package called 'testdoc'\n  Execution halted"},
				Warnings:    []string{"checking for code/documentation mismatches ... WARNING\nCodoc mismatches from documentation object 'my_median':\nmy_median\n  Code: function(x, ...)\n  Docs: function(...)\n  Argument names in code not in docs:\n    x\n  Mismatches in argument names:\n    Position: 1 Code: x Docs: ...", "checking for unstated dependencies in ‘tests’ ... WARNING\n'library' or 'require' call not declared from: ‘testdoc’"},
				Notes:       []string(nil),
			},
			Context: "testwarningerror checklog",
		},
	}

	for _, tst := range tests {
		inputSlice := getFileContents(t, tst.Input)
		actual := parseCheckLog(inputSlice)
		assert.Equal(t, tst.Expected.Errors, actual.Errors, fmt.Sprintf("%s, errors", tst.Context))
		assert.Equal(t, tst.Expected.Warnings, actual.Warnings, fmt.Sprintf("%s, warnings", tst.Context))
		assert.Equal(t, tst.Expected.Notes, actual.Notes, fmt.Sprintf("%s, notes", tst.Context))
	}
}

func TestParseTestsFromCheckLog(t *testing.T) {
	type test struct {
		Input    string
		Expected TestResults
		Context  string
	}
	tests := []test{
		test{
			Input: "./testdata/RcppTOML.Rcheck_fail/00check.log",
			Expected: TestResults{
				Ok:      7,
				Skipped: 0,
				Failed:  1,
			},
			Context: "testRcppTOML.Rcheck_fail checklog",
		},
		test{
			Input: "./testdata/RcppTOML.Rcheck_pass/00check.log",
			Expected: TestResults{
				Ok:      8,
				Skipped: 0,
				Failed:  0,
			},
			Context: "testRcppTOML.Rcheck_pass checklog",
		},
		test{
			Input: "./testdata/data.table.Rcheck_pass/00check.log",
			Expected: TestResults{
				Ok:      2,
				Skipped: 0,
				Failed:  0,
			},
			Context: "data_table.Rcheck_pass checklog",
		},
		test{
			Input: "./testdata/data.table.Rcheck_fail/00check.log",
			Expected: TestResults{
				Ok:      2,
				Skipped: 0,
				Failed:  1,
			},
			Context: "data_table.Rcheck_fail checklog",
		},
	}

	for _, tst := range tests {
		inputSlice := getFileContents(t, tst.Input)
		actual := parseTestsFromCheckLog(inputSlice)
		assert.Equal(t, tst.Expected.Ok, actual.Ok, fmt.Sprintf("Not equal <ok> %s", tst.Context))
		assert.Equal(t, tst.Expected.Skipped, actual.Skipped, fmt.Sprintf("Not equal <skipped> %s", tst.Context))
		assert.Equal(t, tst.Expected.Failed, actual.Failed, fmt.Sprintf("Not equal <fails>: %s", tst.Context))
	}
}

func TestParseTestLog(t *testing.T) {
	type test struct {
		Input    []byte
		Expected TestResults
		Message  string
	}
	tests := []test{
		test{
			Input:    []byte(""),
			Expected: TestResults{},
			Message:  "Not equal: Zero length slice test",
		},
		test{
			Input: []byte("library(testthat)"),
			Expected: TestResults{
				Available: true,
			},
			Message: "Not equal: Contains <library(testthat)> test",
		},
		test{
			Input:    []byte("library)"),
			Expected: TestResults{},
			Message:  "Not equal: Does not contain <library(testthat)> test",
		},
	}
	for _, tst := range tests {
		actual := parseTestLog(tst.Input)
		assert.Equal(t, tst.Expected, actual, tst.Message)
	}
}

func getFileContents(t *testing.T, filepath string) []byte {
	var fileSystem = new(afero.OsFs)

	content, err := afero.ReadFile(fileSystem, filepath)
	if err != nil {
		//return nil, errors.New(fmt.Sprintf("Error: Could not read file %s", filepath))
		t.Error(err)
	}
	return content
}
