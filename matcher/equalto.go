package matcher

import "fmt"

// EqualTo checks if two values are equal.
func EqualTo(expected interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("value equal to <%v>", expected)
	m.matches = func(actual interface{}) bool {
		return expected == actual
	}
	return m
}
