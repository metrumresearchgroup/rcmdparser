package rcmdparser

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

// ParseCheckLog parses the check log
func ParseCheckLog(e []byte) CheckLogEntries {
	splitOutput := bytes.Split(e, []byte("* "))
	var errors []string
	var notes []string
	var warnings []string
	var meta EnvirnomentInformation
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

// ParseTestLog parses the testthat log
func ParseTestLog(e []byte) TestResults {
	if len(e) == 0 {
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
