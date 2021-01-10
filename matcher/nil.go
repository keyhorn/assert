package matcher

import "reflect"

// Nil returns a matcher that returns true if the actual value is nil.
func Nil() *Matcher {
	m := new(Matcher)
	m.Describe = "is <nil>"
	m.matches = func(actual interface{}) bool {
		if actual == nil {
			return true
		}
		switch reflect.TypeOf(actual).Kind() {
		case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
			return reflect.ValueOf(actual).IsNil()
		}
		return false
	}
	return m
}
