package helpers

import "regexp"

func RegexTest(pattern string, value string) bool {
	patternInstance := regexp.MustCompile(pattern)
	return patternInstance.MatchString(value)
}

func RegexTestValidChar(value string) bool {
	patternInstance := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return patternInstance.MatchString(value)
}
