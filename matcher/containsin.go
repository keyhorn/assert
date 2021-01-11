package matcher

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/keyhorn/assert/internal/compare"
)

// ContainsIn returns a matcher that returns true if all expected values contains in actual list at specified order.
func ContainsIn(expected interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("contains in <%v>", expected)
	m.matches = func(t *testing.T, actual interface{}) bool {
		actualValue := reflect.ValueOf(actual)
		switch(actualValue.Kind()) {
		case reflect.Array, reflect.Slice:
			return containsIn(actual, expected)
		}
		return false
	}
	return m
}

func containsIn(actual, expected interface{}) bool {
	actualValue := reflect.ValueOf(actual)
	expectedValue := reflect.ValueOf(expected)
	for x := 0; x < actualValue.Len(); x++ {
		if !compare.EqualTo(actualValue.Index(x).Interface(), expectedValue.Index(x).Interface()) {
			return false
		}
	}
	return true
}
