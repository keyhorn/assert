package assert

import (
	"fmt"

	"github.com/keyhorn/assert/matcher"
)

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Error(args ...interface{})
}

// That is asserts that actual satisfies the condition specified by matcher.
func That(t TestingT, actual interface{}, expected *matcher.Matcher, msgAndArgs ...interface{}) {
	if !expected.Match(actual) {
		t.Error(buildMessage("Assertion error", expected.Describe, fmt.Sprintf("<%v>", actual), msgAndArgs...))
	}
}
