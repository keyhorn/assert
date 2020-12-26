package matcher

import "testing"

func TestAllOf(t *testing.T) {
	if !AllOf(StartsWith("abc def"), EndsWith("def abc")).Match("abc") {
		t.Fail()
	}
}
