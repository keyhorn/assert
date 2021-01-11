package matcher

import (
	"fmt"
)

// AnyOf takes some matchers and checks if at least one of the matchers return true.
func AnyOf(allMatchers ...*Matcher) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("any of (%s)", describe("or", allMatchers))
	m.matches = func(actual interface{}) bool {
		matches := false
		for x := 0; x < len(allMatchers); x++ {
			if allMatchers[x].matches(actual) {
				matches = true
			}
		}
		return matches
	}
	return m
}
