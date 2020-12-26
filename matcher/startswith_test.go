package matcher

import "testing"

func TestStartsWith(t *testing.T) {
	if !StartsWith("hello").Match("hello world") {
		t.Fail()
	}
}
