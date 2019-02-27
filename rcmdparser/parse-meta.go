package rcmdparser

import (
	"bytes"
	"regexp"
)

// Parse consumes a log entry and updates the check metadata if relevant
func (cm *EnvirnomentInformation) Parse(entry []byte) {
	switch {
	case bytes.Contains(entry, []byte("log directory")):
		cm.LogDir = extractBetweenQuotes(entry)
	case bytes.Contains(entry, []byte("R version")):
		cm.Rversion = parseRVersion(entry)
	case bytes.Contains(entry, []byte("platform:")):
		cm.Platform = string(
			bytes.Replace(entry,
				[]byte("* using platform: "),
				[]byte(""), 1),
		)
	case bytes.Contains(entry, []byte("options")):
		cm.Options = extractBetweenQuotes(entry)
	case bytes.Contains(entry, []byte("this is package")):
		cm.Package, cm.PackageVersion = parsePackageInfo(entry)
	default:
	}
}

func extractBetweenQuotes(ent []byte) string {
	sb := bytes.Index(ent, []byte("‘"))
	eb := bytes.Index(ent, []byte("’"))
	if sb == -1 || eb == -1 {
		// didn't parse correctly, return whole entry
		return string(ent)
	}
	// when trying to just clip bytes eg sb+1:eb
	// was getting weird printing artifact, so the index
	// trimming to remote
	return string(bytes.TrimPrefix(ent[sb:eb], []byte("‘")))
}

func parsePackageInfo(ent []byte) (string, string) {
	// this is package ‘test1’ version ‘0.0.1’
	pi := bytes.Split(ent, []byte("version"))
	if len(pi) != 2 {
		// potentially wrong entry passed in
		return "", ""
	}
	return extractBetweenQuotes(pi[0]), extractBetweenQuotes(pi[1])
}

func parseRVersion(ent []byte) string {
	rv := regexp.MustCompile("\\d{1}\\.\\d{1}\\.\\d{1}")
	matches := rv.FindSubmatch(ent)
	if len(matches) == 0 {
		return ""
	}
	return string(matches[0])
}
