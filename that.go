package assert

import (
	"testing"

	"github.com/keyhorn/assert/matcher"
)

// That is asserts that actual satisfies the condition specified by matcher.
func That(t *testing.T, actual interface{}, expected *matcher.Matcher, msgAndArgs ...interface{}) {
	if !expected.Match(actual) {
		t.Fatal(buildMessage("Assertion error", expected.Describe, actual, msgAndArgs...))
	}
}
