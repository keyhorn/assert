package matcher

import "testing"

func TestNot(t *testing.T) {
	if !Not(EqualTo("abcde")).Match(t, "ABCDE") {
		t.Fail()
	}
	if !Not(Nil()).Match(t, "abcde") {
		t.Fail()
	}
}
