package matcher

import (
	"fmt"
	"reflect"
	"testing"
)

// HasValue returns a matcher that returns true if actual has a value with a conditions.
func HasValue(expected *Matcher) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("has value <%v>", expected.Describe)
	m.matches = func(t *testing.T, actual interface{}) bool {
		return hasValue(t, actual, expected)
	}
	return m
}

func hasValue(t *testing.T, actual interface{}, expected *Matcher) bool {
	actualMap := reflect.ValueOf(actual)
	switch(actualMap.Kind()) {
	case reflect.Map:
		iter := actualMap.MapRange()
		for iter.Next() {
			if expected.Match(t, iter.Value().Interface()) {
				return true
			}
		}
	}
	return false
}