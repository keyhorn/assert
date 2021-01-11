package matcher

import "testing"

//Not negates the given matcher.
func Not(matcher *Matcher) *Matcher {
	m := new(Matcher)
	m.Describe = "not(" + matcher.Describe + ")"
	m.matches = func(t *testing.T, actual interface{}) bool {
		return !matcher.matches(t, actual)
	}
	return m
}
