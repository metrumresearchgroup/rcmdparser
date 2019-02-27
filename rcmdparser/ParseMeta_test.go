package rcmdparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
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


func TestParse_LogDirectory(t *testing.T) {
	fixture := EnvirnomentInformation{}
	inputSlice := []byte("* using log directory ‘/Users/devinpastoor/Downloads/output/shiny.Rcheck’")
	fixture.Parse(inputSlice)
	expected := "/Users/devinpastoor/Downloads/output/shiny.Rcheck"
	assert.Equal(t, expected, fixture.LogDir)

}

func TestParse_RVersion(t *testing.T) {
	fixture := EnvirnomentInformation{}
	inputSlice := []byte("using R version 3.5.2 (2018-12-20)")
	fixture.Parse(inputSlice)
	expected := "3.5.2"
	assert.Equal(t, expected, fixture.Rversion)
}

func TestParse_Platform(t *testing.T) {
	fixture := EnvirnomentInformation{}
	inputSlice := []byte("* using platform: x86_64-apple-darwin15.6.0 (64-bit)")
	fixture.Parse(inputSlice)
	expected := "x86_64-apple-darwin15.6.0 (64-bit)"
	assert.Equal(t, expected, fixture.Platform)
}

func TestParse_Options(t *testing.T) {
	fixture := EnvirnomentInformation{}
	inputSlice := []byte("* using options ‘--no-manual --no-build-vignettes’")
	fixture.Parse(inputSlice)
	expected := "--no-manual --no-build-vignettes"
	assert.Equal(t, expected, fixture.Options)
}

func TestParse_PackageAndVersion(t *testing.T) {
	fixture := EnvirnomentInformation{}
	inputSlice := []byte("* this is package ‘shiny’ version ‘1.2.0’")
	fixture.Parse(inputSlice)
	expectedPackage := "shiny"
	expectedVersion := "1.2.0"
	assert.Equal(t, expectedPackage, fixture.Package)
	assert.Equal(t, expectedVersion, fixture.PackageVersion)
}