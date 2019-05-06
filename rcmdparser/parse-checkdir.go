package rcmdparser

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/dpastoor/goutils"

	"github.com/spf13/afero"
)

// ParseCheckDir reads the relevant files in a check directory
// should take form of
// * 00check.log
// * 00install.out
// * (maybe) tests/testthat.Rout
// * (maybe) tests/testthat.Rout.fail
//
// cd - CheckOutputRaw directory
func parseCheckDir(fs afero.Fs, cd string) (CheckOutputInfo, error) {
	ok, _ := goutils.DirExists(fs, cd)
	if !ok {
		return CheckOutputInfo{}, fmt.Errorf("dir does not exist: %s", cd)
	}

	checkFilePath := filepath.Join(cd, "00check.log")
	installFilePath := filepath.Join(cd, "00install.out")

	check, err := afero.ReadFile(fs, checkFilePath)
	if err != nil {
		// if the checkfile doesn't exist, most likely something more
		// drastic went wrong
		return CheckOutputInfo{}, err
	}

	//checkLogEntries := ParseCheckLog(check)

	install, err := afero.ReadFile(fs, installFilePath)
	if err != nil {
		// if the checkfile doesn't exist, most likely something more
		// drastic went wrong, like missing system dependency
		return CheckOutputInfo{}, err
	}

	var test TestInfo
	hasTests, _ := goutils.DirExists(fs, filepath.Join(cd, "tests"))

	if hasTests {

		// regular tests
		// TODO(devin): implement tests for non-testthat package
		test.HasTests = true
		testthatFilePath := filepath.Join(cd, "tests", "testthat.Rout")
		testthatFileFailPath := filepath.Join(cd, "tests", "testthat.Rout.fail")

		if exists(fs, testthatFilePath) {
			testFile, _ := afero.ReadFile(fs, testthatFilePath)
			test.UsesTestthat = true
			test.ResultsFile = testFile
		} else if exists(fs, testthatFileFailPath) {
			testFile, _ := afero.ReadFile(fs, testthatFileFailPath)
			test.UsesTestthat = true
			test.ResultsFile = testFile
		}
	}

	return CheckOutputInfo{
		TestInfo:         test,
		CheckOutputRaw:   check,
		InstallOutputRaw: install,
		//CheckParsed: CheckResults {
		//	Checks: checkLogEntries,
		//},
	}, nil
}

func exists(fs afero.Fs, path string) bool {
	ok, _ := goutils.Exists(fs, path)
	return ok
}
func parseEntries(b []byte) [][]byte {
	splitFile := bytes.Split(b, []byte("* "))
	return splitFile
}
