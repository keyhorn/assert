package matcher

import (
	"fmt"
	"reflect"
)

// HasValue returns a matcher that returns true if actual has a value with a conditions.
func HasValue(expected *Matcher) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("has value <%v>", expected.Describe)
	m.matches = func(actual interface{}) bool {
		actualMap := reflect.ValueOf(actual)
		switch(actualMap.Kind()) {
		case reflect.Map:
			iter := actualMap.MapRange()
			for iter.Next() {
				if expected.Match(iter.Value().Interface()) {
					return true
				}
			}
		}
		return false
	}
	return m
}
