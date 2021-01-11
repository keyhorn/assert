package matcher

import (
	"testing"
)

func TestAllOf(t *testing.T) {
	if !AllOf(EqualTo("abc"), StartsWith("a"), EndsWith("c")).Match(t, "abc") {
		t.Fail()
	}
	if AllOf(EqualTo("abc"), StartsWith("a"), EndsWith("c")).Match(t, "aBc") {
		t.Fail()
	}
}
