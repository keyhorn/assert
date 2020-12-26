package matcher

import "testing"

func TestAnyOf(t *testing.T) {
	if !AnyOf(EqualTo("abc"), EqualTo("def")).Match("abc") {
		t.Fail()
	}
}
