package matcher

import (
	"fmt"

	"github.com/keyhorn/assert/internal/compare"
)

// GreaterThan matcher compares two values that are numeric or string values, and when
// called returns true if actual > expected. Strings are compared lexicographically with '>'.
// The matcher will always return false for unknown types.
func GreaterThan(expected interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("value greater than <%v>", expected)
	m.matches = func(actual interface{}) bool {
		return compare.GreaterThan(actual, expected)
	}
	return m
}

// GreaterThanOrEqualTo is a short hand matcher for anyOf(GreaterThan(x), EqualTo(x))
func GreaterThanOrEqualTo(expected interface{}) *Matcher {
	m := new(Matcher)
	m.matches = func(actual interface{}) bool {
		anyOf := AnyOf(GreaterThan(expected), EqualTo(expected))
		m.Describe = anyOf.Describe
		return anyOf.matches(actual)
	}
	return m
}
