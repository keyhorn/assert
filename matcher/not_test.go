package matcher

import "testing"

func TestNot(t *testing.T) {
	if !Not(EqualTo("abcde")).Match("ABCDE") {
		t.Fail()
	}
	if !Not(Nil()).Match("abcde") {
		t.Fail()
	}
}
