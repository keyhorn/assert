package matcher

import (
	"reflect"
)

// Empty returns a matcher that returns true if the actual is "empty".
func Empty() *Matcher {
	m := new(Matcher)
	m.Describe = "is empty"
	m.matches = func(actual interface{}) bool {
		if actual == nil {
			return true
		}
		if actualValue, ok := actual.(string); ok {
			return actualValue == ""
		}
		actualValue := reflect.ValueOf(actual)
		switch actualValue.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			if reflect.ValueOf(actual).Len() == 0 {
				return true
			}
		}
		return false
	}
	return m
}
