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
		Fixture  EnvirnomentInformation
		Input    []byte
		Expected EnvirnomentInformation
		Messsage string
	}
	tests := []metaTest{
		metaTest{
			Fixture: EnvirnomentInformation{},
			Input:   []byte("* using log directory ‘/Users/devinpastoor/Downloads/output/shiny.Rcheck’"),
			Expected: EnvirnomentInformation{
				LogDir: "/Users/devinpastoor/Downloads/output/shiny.Rcheck",
			},
			Messsage: "Meta data not equal: LogDir",
		},
		metaTest{
			Fixture: EnvirnomentInformation{},
			Input:   []byte("using R version 3.5.2 (2018-12-20)"),
			Expected: EnvirnomentInformation{
				Rversion: "3.5.2",
			},
			Messsage: "Meta data not equal: Rversion",
		},
		metaTest{
			Fixture: EnvirnomentInformation{},
			Input:   []byte("* using platform: x86_64-apple-darwin15.6.0 (64-bit)"),
			Expected: EnvirnomentInformation{
				Platform: "x86_64-apple-darwin15.6.0 (64-bit)",
			},
			Messsage: "Meta data not equal: Platform",
		},
		metaTest{
			Fixture: EnvirnomentInformation{},
			Input:   []byte("* using options ‘--no-manual --no-build-vignettes’"),
			Expected: EnvirnomentInformation{
				Options: "--no-manual --no-build-vignettes",
			},
			Messsage: "Meta data not equal: Options",
		},
		metaTest{
			Fixture: EnvirnomentInformation{},
			Input:   []byte("* using options ‘--no-manual --no-build-vignettes’"),
			Expected: EnvirnomentInformation{
				Options: "--no-manual --no-build-vignettes",
			},
			Messsage: "Meta data not equal: Options",
		},
		metaTest{
			Fixture: EnvirnomentInformation{},
			Input:   []byte("* this is package ‘shiny’ version ‘1.2.0’"),
			Expected: EnvirnomentInformation{
				Package:        "shiny",
				PackageVersion: "1.2.0",
			},
			Messsage: "Meta data not equal: Package and/or PackageVersion",
		},
	}

	for _, tst := range tests {
		actual := EnvirnomentInformation{}
		actual.Parse(tst.Input)
		assert.Equal(t, tst.Expected, actual, tst.Messsage)
	}
}
