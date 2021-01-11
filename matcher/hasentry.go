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
		actualMap := reflect.ValueOf(actual)
		switch(actualMap.Kind()) {
		case reflect.Map:
			mapKeys := actualMap.MapKeys()
			for x := 0; x < len(mapKeys); x++ {
				if key.Match(t, mapKeys[x].Interface()) && value.Match(t, actualMap.MapIndex(mapKeys[x]).Interface()) {
						return true
				}
			}
		}
		return false
	}
	return m
}
