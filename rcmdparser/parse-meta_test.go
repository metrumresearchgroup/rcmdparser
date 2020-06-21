package rcmdparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
* using log directory ‘/Users/devinpastoor/Downloads/output/shiny.Rcheck’
* using R version 3.5.2 (2018-12-20)
* using platform: x86_64-apple-darwin15.6.0 (64-bit)
* using session charset: UTF-8
* using options ‘--no-manual --no-build-vignettes’
* checking for file ‘shiny/DESCRIPTION’ ... OK
* checking extension type ... Package
* this is package ‘shiny’ version ‘1.2.0’
* checking package namespace information ... OK
 */

func Test_ParseMeta(t *testing.T) {
	type metaTest struct {
		Fixture  EnvironmentInformation
		Input    []byte
		Expected EnvironmentInformation
		Messsage string
	}
	tests := []metaTest{
		metaTest{
			Fixture: EnvironmentInformation{},
			Input:   []byte("* using log directory ‘/Users/devinpastoor/Downloads/output/shiny.Rcheck’"),
			Expected: EnvironmentInformation{
				LogDir: "/Users/devinpastoor/Downloads/output/shiny.Rcheck",
			},
			Messsage: "Meta data not equal: LogDir",
		},
		metaTest{
			Fixture: EnvironmentInformation{},
			Input:   []byte("using R version 3.5.2 (2018-12-20)"),
			Expected: EnvironmentInformation{
				Rversion: "3.5.2",
			},
			Messsage: "Meta data not equal: Rversion",
		},
		metaTest{
			Fixture: EnvironmentInformation{},
			Input:   []byte("* using platform: x86_64-apple-darwin15.6.0 (64-bit)"),
			Expected: EnvironmentInformation{
				Platform: "x86_64-apple-darwin15.6.0 (64-bit)",
			},
			Messsage: "Meta data not equal: Platform",
		},
		metaTest{
			Fixture: EnvironmentInformation{},
			Input:   []byte("* using options ‘--no-manual --no-build-vignettes’"),
			Expected: EnvironmentInformation{
				Options: "--no-manual --no-build-vignettes",
			},
			Messsage: "Meta data not equal: Options",
		},
		metaTest{
			Fixture: EnvironmentInformation{},
			Input:   []byte("* using options ‘--no-manual --no-build-vignettes’"),
			Expected: EnvironmentInformation{
				Options: "--no-manual --no-build-vignettes",
			},
			Messsage: "Meta data not equal: Options",
		},
		metaTest{
			Fixture: EnvironmentInformation{},
			Input:   []byte("* this is package ‘shiny’ version ‘1.2.0’"),
			Expected: EnvironmentInformation{
				Package:        "shiny",
				PackageVersion: "1.2.0",
			},
			Messsage: "Meta data not equal: Package and/or PackageVersion",
		},
	}

	for _, tst := range tests {
		actual := EnvironmentInformation{}
		actual.Parse(tst.Input)
		assert.Equal(t, tst.Expected, actual, tst.Messsage)
	}
}
