package matcher

import (
	"fmt"
	"strings"
)

// Contains returns a matcher that returns true if the argument is a string that contains a specific substring.
func Contains(expected string) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("ontains \"%s\"", expected)
	m.matches = func(actual interface{}) bool {
		if actualStr, ok := actual.(string); ok {
			return strings.Contains(actualStr, expected)
		}
		return false
	}
	return m
}
