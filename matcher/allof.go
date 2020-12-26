package matcher

import "fmt"

// AllOf takes some matchers and checks if at least one of the matchers return true.
func AllOf(allMatchers ...*Matcher) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("all of (%s)", describe("and", allMatchers))
	m.matches = func(actual interface{}) bool {
		matches := true
		for x := 0; x < len(allMatchers); x++ {
			if allMatchers[x].matches(actual) {
				matches = false
			}
		}
		return matches
	}
	return m
}
