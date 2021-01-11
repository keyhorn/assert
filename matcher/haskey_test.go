package matcher

import "testing"

func TestHasKey1(t *testing.T) {
	var items map[string]string = map[string]string{
		"go": "golang",
		"rb": "ruby",
		"js": "javascript",
	}

	if !HasKey(EqualTo("go")).Match(t, items) {
		t.Fail()
	}
}

func TestHasKey2(t *testing.T) {
	var items map[string]string = map[string]string{
		"go": "golang",
		"rb": "ruby",
		"js": "javascript",
	}

	if HasKey(EqualTo("c++")).Match(t, items) {
		t.Fail()
	}
}

func TestHasKey3(t *testing.T) {
	if HasKey(EqualTo("js")).Match(t, "js") {
		t.Fail()
	}
}