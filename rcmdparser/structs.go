package rcmdparser

// FILE LEVEL PARSE
// CheckResults is a struct of the R CMD check results
type CheckResults struct {
	Checks CheckLogEntries
	Tests  TestResults
}

// CheckLogEntries are the parsed results from the check log
type CheckLogEntries struct {
	Environment EnvirnomentInformation
	Errors      []string
	Warnings    []string
	Notes       []string
}

// EnvirnomentInformation stores metadata about the RCMDCHECK
type EnvirnomentInformation struct {
	LogDir         string
	Rversion       string
	Platform       string
	Options        string
	Package        string
	PackageVersion string
}

// TestResults is the results from testthat
type TestResults struct {
	Ok        int
	Skipped   int
	Failed    int
	Output    string
	Available bool
}

/////////////////

//DIRECTORY LEVEL PARSE
// CheckOutputInfo represents key elements of a R CMD check output directory
type CheckOutputInfo struct {
	TestInfo         TestInfo
	CheckOutputRaw   []byte
	InstallOutputRaw []byte
	//CheckParsed      CheckResults
}

// TestInfo represents output of tests and whether it uses testthat
type TestInfo struct {
	HasTests     bool
	UsesTestthat bool
	ResultsFile  []byte
}










