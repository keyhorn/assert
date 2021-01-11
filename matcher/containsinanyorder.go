package matcher

import (
	"fmt"
	"reflect"

	"github.com/keyhorn/assert/internal/compare"
)

// ContainsInAnyOrder returns a matcher that returns true if all expected values contains in actual list at any order.
func ContainsInAnyOrder(expected interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("contains in <%v>", expected)
	m.matches = func(actual interface{}) bool {
		actualValue := reflect.ValueOf(actual)
		switch(actualValue.Kind()) {
		case reflect.Array, reflect.Slice:
			return containsInAnyOrder(actual, expected)
		}
		return false
	}
	return m
}

func containsInAnyOrder(actual, expected interface{}) bool {
	actualValue := reflect.ValueOf(actual)
	expectedValue := reflect.ValueOf(expected)
	for x := 0; x < actualValue.Len(); x++ {
		match := false
		for y := 0; y < expectedValue.Len(); y++ {
			if compare.EqualTo(actualValue.Index(x).Interface(), expectedValue.Index(y).Interface()) {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}
	return true
}
