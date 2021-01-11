package matcher

import (
	"fmt"
	"testing"

	"github.com/keyhorn/assert/internal/compare"
)

// EqualTo returns a matcher that Returns true if two values are equal.
func EqualTo(expected interface{}) *Matcher {
	m := new(Matcher)
	m.Describe = fmt.Sprintf("equal to <%v>", expected)
	m.matches = func(t *testing.T, actual interface{}) bool {
		return compare.EqualTo(expected, actual)
	}
	return m
}
