package matcher

import "testing"

func TestEmpty(t *testing.T) {
	if !Empty().Match(nil) {
		t.Fail()
	}
	if !Empty().Match("") {
		t.Fail()
	}
}
