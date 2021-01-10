package matcher

import "testing"

func TestHasValue1(t *testing.T) {
	var items map[string]string = map[string]string{
		"go": "golang",
		"rb": "ruby",
		"js": "javascript",
	}

	if !HasValue(EqualTo("golang")).Match(items) {
		t.Fail()
	}
}

func TestHasValue2(t *testing.T) {
	var items map[string]string = map[string]string{
		"go": "golang",
		"rb": "ruby",
		"js": "javascript",
	}

	if HasValue(EqualTo("c++")).Match(items) {
		t.Fail()
	}
}

func TestHasValue3(t *testing.T) {
	if HasValue(EqualTo("js")).Match("js") {
		t.Fail()
	}
}