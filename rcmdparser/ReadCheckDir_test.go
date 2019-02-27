package rcmdparser

import (
	"reflect"
	"testing"

	"github.com/dpastoor/goutils"

	"github.com/spf13/afero"
)

func TestReadCheckDir(t *testing.T) {
	testFS := afero.NewMemMapFs()
	// tests dir but no testthat
	_ = testFS.MkdirAll("noTestThat/tests", 0755)
	_ = goutils.WriteLinesFS(testFS, []string{"log"}, "noTestThat/00check.log")
	_ = goutils.WriteLinesFS(testFS, []string{"install"}, "noTestThat/00install.out")

	// with testthat
	_ = testFS.MkdirAll("WithTestThat/tests", 0755)
	_ = goutils.WriteLinesFS(testFS, []string{"log"}, "WithTestThat/00check.log")
	_ = goutils.WriteLinesFS(testFS, []string{"install"}, "WithTestThat/00install.out")
	_ = goutils.WriteLinesFS(testFS, []string{"tests"}, "WithTestThat/tests/testthat.Rout")

	// Failed UsesTestthat
	_ = testFS.MkdirAll("FailedTest/tests", 0755)
	_ = goutils.WriteLinesFS(testFS, []string{"log"}, "FailedTest/00check.log")
	_ = goutils.WriteLinesFS(testFS, []string{"install"}, "FailedTest/00install.out")
	_ = goutils.WriteLinesFS(testFS, []string{"failed-tests"}, "FailedTest/tests/testthat.Rout.fail")
	
	var cdtests = []struct {
		in       string
		expected CheckOutputInfo
	}{
		{
			"noTestThat",
			CheckOutputInfo{
				TestInfo{true, false, nil},
				[]byte("log\n"),
				[]byte("install\n"),
			},
		},
		{
			"WithTestThat",
			CheckOutputInfo{
				TestInfo{true, true, []byte("tests\n")},
				[]byte("log\n"),
				[]byte("install\n"),
			},
		},
		{
			"FailedTest",
			CheckOutputInfo{
				TestInfo{true, true, []byte("failed-tests\n")},
				[]byte("log\n"),
				[]byte("install\n"),
			},
		},
	}
	for _, tt := range cdtests {
		actual, _ := ParseCheckDir(testFS, tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("GOT: %v, WANT: %v", actual, tt.expected)
		}
	}
}
