package matcher

import (
	"fmt"
	"reflect"
)

// Key is a matcher that checks if actual has a key == expected.
func Key(expected interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("map has key '%s'", expected)
	m.matches = func(actual interface{}) bool {
		return hasKey(actual, expected)
	}
	return m
}

//AllKeys is a matcher that checks if map actual has all keys == expecteds.
func AllKeys(expected ...interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("map has keys '%s'", expected)
	m.matches = func(actual interface{}) bool {
		keyValuesToMatch := reflect.ValueOf(correctExpectedValue(expected...))
		for x := 0; x < keyValuesToMatch.Len(); x++ {
			if !hasKey(actual, keyValuesToMatch.Index(x).Interface()) {
				return false
			}
		}
		return true
	}
	return m
}

func hasKey(actual interface{}, expected interface{}) bool {
	mapKeys := reflect.ValueOf(actual).MapKeys()
	for x := 0; x < len(mapKeys); x++ {
		if mapKeys[x].Interface() == expected {
			return true
		}
	}
	return false
}

func correctExpectedValue(expected ...interface{}) interface{} {
	kind := reflect.ValueOf(expected[0]).Kind()
	if kind == reflect.Slice {
		return expected[0]
	}
	return expected[:][0]
}
