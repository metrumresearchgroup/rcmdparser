package rcmdparser

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

// ParseCheckLog parses the check log
func parseCheckLog(e []byte) CheckLogEntries {
	splitOutput := bytes.Split(e, []byte("* "))
	var errors []string
	var notes []string
	var warnings []string
	var meta EnvironmentInformation
	for _, entry := range splitOutput {

		switch {

		case bytes.Contains(entry, []byte("... NOTE")):
			notes = append(notes, strings.TrimSpace(string(entry)))

		case bytes.Contains(entry, []byte("... ERROR")):
			errors = append(errors, strings.TrimSpace(string(entry)))

		case bytes.Contains(entry, []byte("... WARNING")):
			warnings = append(warnings, strings.TrimSpace(string(entry)))

		default:
			meta.Parse(entry)
		}
	}

	return CheckLogEntries{
		Environment: meta,
		Errors:      errors,
		Notes:       notes,
		Warnings:    warnings,
	}
}

func parseTestsFromCheckLog(rawLog []byte) TestResults {
	if len(rawLog) == 0 {
		return TestResults{}
	}
	var ok = 0
	var failed = 0
	var skipped = 0
	var unknown = 0

	splitOutput := bytes.Split(rawLog, []byte("* "))
	for _, entry := range splitOutput {
		if bytes.Contains(entry, []byte("checking tests ...")) {
			//I have to split on "\sRunning\s" to avoid splitting on the word "Running" when it reappaers in failure results,
			//specifically with regards to the data.table failed log.
			tests := bytes.Split(entry, []byte(" Running "))[1:] //Cut off the "checking tests" section.
			for _, test := range tests {
				// converting to string to take advantage of TrimSpace and HasSuffix
				str := strings.TrimSpace(string(test))
				if strings.HasSuffix(str, "... OK"){
					ok++
				} else if strings.Contains(str, "ERROR") || strings.Contains(str, "fail") || strings.Contains(str, "  Comparing"){
					failed++
				} else {
					unknown++
				}
			}
		}
	}

	//TODO: Can tests be skipped like this?
	return TestResults{
		Ok: ok,
		Failed: failed,
		Skipped: skipped,
		Unknown: unknown,
	}
}

// ParseTestLog parses the testthat log
func parseTestLog(e []byte) TestResults {
	if len(e) == 0 || !bytes.Contains(e, []byte("library(testthat)")) {
		return TestResults{}
	}
	contents := bytes.Split(e, []byte("library(testthat)"))[1]
	scanner := bufio.NewScanner(strings.NewReader(string(e)))
	scanner.Split(bufio.ScanLines)
	tr := TestResults{
		Output:    string(contents),
		Available: true,
	}
	re, _ := regexp.Compile(`OK: (\d+) SKIPPED: (\d+) FAILED: (\d+)`)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "OK:") {
			res := re.FindAllStringSubmatch(text, 1)[0]
			// string submatch should be:
			// 0 - full match
			// 1 - OK
			// 2 - SKIPPED
			// 3 - FAILED
			tr.Ok, _ = strconv.Atoi(res[1])
			tr.Skipped, _ = strconv.Atoi(res[2])
			tr.Failed, _ = strconv.Atoi(res[3])
			break
		}
	}
	return tr
}
