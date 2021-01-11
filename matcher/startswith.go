package matcher

import (
	"fmt"
	"strings"
)

// StartsWith returns a matcher that compares two values that are string values,
// and returns true if the given string is starts with the expected string.
func StartsWith(expected string) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("starts with \"%s\"", expected)
	m.matches = func(actual interface{}) bool {
		if actualStr, ok := actual.(string); ok {
			return strings.HasPrefix(actualStr, expected)
		}
		return false
	}
	return m
}
