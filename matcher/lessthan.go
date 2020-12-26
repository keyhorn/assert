package matcher

import (
	"fmt"
	"reflect"
)

// LessThan matcher compares two values that are numeric or string values, and when
// called returns true if actual < expected. Strings are compared lexicographically with '<'.
// The matcher will always return false for unknown types.
func LessThan(expected interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("value less than <%v>", expected)
	m.matches = func(actual interface{}) bool {
		expectedValue := reflect.ValueOf(expected)
		actualValue := reflect.ValueOf(actual)
		switch expected.(type) {
		case float32, float64:
			switch actual.(type) {
			case float32, float64:
				return actualValue.Float() < expectedValue.Float()
			case int, int8, int16, int32, int64:
				return actualValue.Float() < expectedValue.Float()
			case uint, uint8, uint16, uint32, uint64:
				return actualValue.Float() < expectedValue.Float()
			default:
				return false
			}
		case int, int8, int16, int32, int64:
			switch actual.(type) {
			case float32, float64:
				return actualValue.Float() < expectedValue.Float()
			case int, int8, int16, int32, int64:
				return actualValue.Int() < expectedValue.Int()
			case uint, uint8, uint16, uint32, uint64:
				return actualValue.Int() < expectedValue.Int()
			default:
				return false
			}
		case uint, uint8, uint16, uint32, uint64:
			switch actual.(type) {
			case float32, float64:
				return actualValue.Float() < expectedValue.Float()
			case int, int8, int16, int32, int64:
				return actualValue.Int() < expectedValue.Int()
			case uint, uint8, uint16, uint32, uint64:
				return actualValue.Uint() < expectedValue.Uint()
			default:
				return false
			}
		case string:
			return actualValue.String() < expectedValue.String()
		}
		return false
	}
	return m
}

// LessThanOrEqualTo is a short hand matcher for anyOf(LessThan(x), equalTo(x))
func LessThanOrEqualTo(expected interface{}) *Matcher {
	m := new(Matcher)
	m.matches = func(actual interface{}) bool {
		anyOf := AnyOf(LessThan(expected), EqualTo(expected))
		m.Describe = anyOf.Describe
		return anyOf.matches(actual)
	}
	return m
}
