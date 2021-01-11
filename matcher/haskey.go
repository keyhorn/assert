package matcher

import (
	"fmt"
	"reflect"
	"testing"
)

// HasKey returns a matcher that returns true if actual has a key with a conditions.
func HasKey(expected *Matcher) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("has key <%v>", expected.Describe)
	m.matches = func(t *testing.T, actual interface{}) bool {
		return hasKey(t, actual, expected)
	}
	return m
}

func hasKey(t *testing.T, actual interface{}, expected *Matcher) bool {
	actualMap := reflect.ValueOf(actual)
	switch(actualMap.Kind()) {
	case reflect.Map:
		mapKeys := actualMap.MapKeys()
		for x := 0; x < len(mapKeys); x++ {
			if expected.Match(t, mapKeys[x].Interface()) {
				return true
			}
		}
	}
	return false
}
