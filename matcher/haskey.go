package matcher

import (
	"fmt"
	"reflect"
)

// HasKey returns a matcher that returns true if actual has a key with a conditions.
func HasKey(expected *Matcher) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("has key <%v>", expected.Describe)
	m.matches = func(actual interface{}) bool {
		return hasKey(actual, expected)
	}
	return m
}

func hasKey(actual interface{}, expected *Matcher) bool {
	actualMap := reflect.ValueOf(actual)
	switch(actualMap.Kind()) {
	case reflect.Map:
		mapKeys := actualMap.MapKeys()
		for x := 0; x < len(mapKeys); x++ {
			if expected.Match(mapKeys[x].Interface()) {
				return true
			}
		}
	}
	return false
}
