package matcher

import (
	"fmt"
	"strings"
)

//StartsWith returns a matcher that matches if the given string is prefixed with the expected string.
func StartsWith(expected string) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("value with prefix %s", expected)
	m.matches = func(actual interface{}) bool {
		if actualStr, ok := actual.(string); ok {
			return strings.HasPrefix(actualStr, expected)
		}
		return false
	}
	return m
}
