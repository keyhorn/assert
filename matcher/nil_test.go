package matcher

import "testing"

func TestNil(t *testing.T) {
	if !Nil().Match(nil) {
		t.Fail()
	}
}
