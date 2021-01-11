package matcher

import (
	"testing"
)

func TestAllOf(t *testing.T) {
	if !AllOf(EqualTo("abc"), StartsWith("a"), EndsWith("c")).Match("abc") {
		t.Fail()
	}
	if AllOf(EqualTo("abc"), StartsWith("a"), EndsWith("c")).Match("aBc") {
		t.Fail()
	}
}
