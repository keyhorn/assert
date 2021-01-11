package matcher

import (
	"fmt"
	"reflect"
	"testing"
)

// HasEntry returns a matcher that returns true if actual has entry with a conditions.
func HasEntry(key *Matcher, value *Matcher) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("has key <%v> and value <%v>", key.Describe, value.Describe)
	m.matches = func(t *testing.T, actual interface{}) bool {
		return hasEntry(t, actual, key, value)
	}
	return m
}

func hasEntry(t *testing.T, actual interface{}, expectedKey *Matcher, expectedValue *Matcher) bool {
	actualMap := reflect.ValueOf(actual)
	switch(actualMap.Kind()) {
	case reflect.Map:
		mapKeys := actualMap.MapKeys()
		for x := 0; x < len(mapKeys); x++ {
			if expectedKey.Match(t, mapKeys[x].Interface()) {
				if expectedValue.Match(t, actualMap.MapIndex(mapKeys[x]).Interface()) {
					return true
				}
			}
		}
	}
	return false
}