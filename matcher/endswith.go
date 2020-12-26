package matcher

import (
	"fmt"
	"strings"
)

//EndsWith returns a matcher that matches if the given string is suffixed with the expected string.
func EndsWith(expected string) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("value with suffix %s", expected)
	m.matches = func(actual interface{}) bool {
		if actualStr, ok := actual.(string); ok {
			return strings.HasSuffix(actualStr, expected)
		}
		return false
	}
	return m
}
