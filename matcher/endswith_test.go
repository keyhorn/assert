package matcher

import "testing"

func TestEndsWith(t *testing.T) {
	if !EndsWith("world").Match("hello world") {
		t.Fail()
	}
}
