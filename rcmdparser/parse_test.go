package rcmdparser

import (
	"github.com/franela/goblin"
	"reflect"
	"testing"
)

func TestParseTestLog(t *testing.T) {

	var pttests = []struct {
		in       []byte
		expected TestResults
	}{
		{
			[]byte(`
				Stuff above

				> library(testthat)
				> library(test1)
				> 
				> test_check("test1")
				══ testthat results  ═══════════════════════════════════════════════════════════
				OK: 1 SKIPPED: 0 FAILED: 0
				> 
				> proc.time()
				   user  system elapsed 
				  0.866   0.075   0.874 
				`),
			TestResults{
				Ok:      1,
				Skipped: 0,
				Failed:  0,
				Output: `
				> library(test1)
				> 
				> test_check("test1")
				══ testthat results  ═══════════════════════════════════════════════════════════
				OK: 1 SKIPPED: 0 FAILED: 0
				> 
				> proc.time()
				   user  system elapsed 
				  0.866   0.075   0.874 
				`,
				Available: true,
			},
		},
		// test double digits
		{
			[]byte(`
				Stuff above

				> library(testthat)
				> library(test1)
				> 
				> test_check("test1")
				══ testthat results  ═══════════════════════════════════════════════════════════
				OK: 12 SKIPPED: 10 FAILED: 11 
				> 
				> proc.time()
				   user  system elapsed 
				  0.866   0.075   0.874 
				`),
			TestResults{
				Ok:      12,
				Skipped: 10,
				Failed:  11,
				Output: `
				> library(test1)
				> 
				> test_check("test1")
				══ testthat results  ═══════════════════════════════════════════════════════════
				OK: 12 SKIPPED: 10 FAILED: 11 
				> 
				> proc.time()
				   user  system elapsed 
				  0.866   0.075   0.874 
				`,
				Available: true,
			},
		},
		{
			nil,
			TestResults{
				Ok:        0,
				Skipped:   0,
				Failed:    0,
				Available: false,
			},
		},
	}
	for _, tt := range pttests {
		actual := ParseTestLog(tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("GOT: %v, WANT: %v", actual, tt.expected)
		}
	}
}

func TestParseCheckLog(t *testing.T) {

	var inputLog =
		`* checking foreign function calls ... OK
* checking R code for possible problems ... OK
* checking Rd files ... OK
* checking Rd metadata ... OK
* checking Rd cross-references ... OK
* checking for missing documentation entries ... OK
* checking for code/documentation mismatches ... WARNING
Codoc mismatches from documentation object 'my_median':
my_median
  Code: function(x, ...)
  Docs: function(...)
  Argument names in code not in docs:
    x
  Mismatches in argument names:
    Position: 1 Code: x Docs: ...

* checking Rd \usage sections ... OK
* checking Rd contents ... OK
* checking for unstated dependencies in examples ... OK
* checking examples ... NONE
* checking for unstated dependencies in ‘tests’ ... WARNING
'library' or 'require' call not declared from: ‘testdoc’
* checking tests ... ERROR
  Running ‘testthat.R’
Running the tests in ‘tests/testthat.R’ failed.
Complete output:
  > library(testthat)
  > library(testdoc)
  Error in library(testdoc) : there is no package called 'testdoc'
  Execution halted
* DONE
Status: 1 ERROR, 2 WARNINGs`

	expected := LogEntries{
		Errors: []string{

			`checking tests ... ERROR
  Running ‘testthat.R’
Running the tests in ‘tests/testthat.R’ failed.
Complete output:
  > library(testthat)
  > library(testdoc)
  Error in library(testdoc) : there is no package called 'testdoc'
  Execution halted`,
		},
		Meta:  CheckMeta{},
		Notes: []string{},
		Warnings: []string{

			`checking for code/documentation mismatches ... WARNING
Codoc mismatches from documentation object 'my_median':
my_median
  Code: function(x, ...)
  Docs: function(...)
  Argument names in code not in docs:
    x
  Mismatches in argument names:
    Position: 1 Code: x Docs: ...`,

			`checking for unstated dependencies in ‘tests’ ... WARNING
'library' or 'require' call not declared from: ‘testdoc’`,
		},
	}

	var actual = ParseCheckLog([]byte(inputLog))
	g := goblin.Goblin(t)
	g.Describe("Log with errors and warnings", func() {
		g.It("shouldn't drive me crazy like this0", func() {
			g.Assert(actual.Warnings[0]).Equal(expected.Warnings[0])
		})
		g.It("shouldn't drive me crazy like this1", func() {
			g.Assert(actual.Warnings[1]).Equal(expected.Warnings[1])
		})
		g.It("should parse errors correctly", func() {
			g.Assert(actual.Errors).Equal(expected.Errors)
		})
		g.It("should parse warnings correctly", func() {
			g.Assert(actual.Warnings).Equal(expected.Warnings)
		})
	})
}