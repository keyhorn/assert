package matcher

import (
	"fmt"
	"testing"

	"github.com/keyhorn/assert/internal/compare"
)

// LessThan returns a matcher that compares two values that are numeric or string values,
// and returns true if actual < expected.
func LessThan(expected interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("value less than <%v>", expected)
	m.matches = func(t *testing.T, actual interface{}) bool {
		return compare.LessThan(actual, expected)
	}
	return m
}

// LessThanOrEqualTo is a short hand matcher for AnyOf(LessThan(x), equalTo(x))
func LessThanOrEqualTo(expected interface{}) *Matcher {
	m := new(Matcher)
	m.matches = func(t *testing.T, actual interface{}) bool {
		anyOf := AnyOf(LessThan(expected), EqualTo(expected))
		m.Describe = anyOf.Describe
		return anyOf.matches(t, actual)
	}
	return m
}
