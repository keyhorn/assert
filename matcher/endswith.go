package matcher

import (
	"fmt"
	"strings"
)

// EndsWith returns a matcher that compares two values that are string values,
// and returns true if the given string is ends with the expected string.
func EndsWith(expected string) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("ends with \"%s\"", expected)
	m.matches = func(actual interface{}) bool {
		if actualStr, ok := actual.(string); ok {
			return strings.HasSuffix(actualStr, expected)
		}
		return false
	}
	return m
}
