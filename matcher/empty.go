package matcher

import "reflect"

// Empty matches if the actual is "empty".
func Empty() *Matcher {
	m := new(Matcher)
	m.Describe = "empty value"
	m.matches = func(actual interface{}) bool {
		if actual == nil {
			return true
		}
		if actualValue, ok := actual.(string); ok {
			return actualValue == ""
		}
		if reflect.ValueOf(actual).Len() == 0 {
			return true
		}
		return false
	}
	return m
}
