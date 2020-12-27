package matcher

import (
	"fmt"

	"github.com/keyhorn/assert/internal/compare"
)

// LessThan matcher compares two values that are numeric or string values, and when
// called returns true if actual < expected. Strings are compared lexicographically with '<'.
// The matcher will always return false for unknown types.
func LessThan(expected interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("value less than <%v>", expected)
	m.matches = func(actual interface{}) bool {
		return compare.LessThan(actual, expected)
	}
	return m
}

// LessThanOrEqualTo is a short hand matcher for anyOf(LessThan(x), equalTo(x))
func LessThanOrEqualTo(expected interface{}) *Matcher {
	m := new(Matcher)
	m.matches = func(actual interface{}) bool {
		anyOf := AnyOf(LessThan(expected), EqualTo(expected))
		m.Describe = anyOf.Describe
		return anyOf.matches(actual)
	}
	return m
}
