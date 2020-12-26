package matcher

import "testing"

func TestKey(t *testing.T) {
	var items map[string]string = map[string]string{
		"go": "golang",
		"rb": "ruby",
		"js": "javascript",
	}

	if !Key("go").Match(items) {
		t.Fail()
	}
}

func TestAllKeys1(t *testing.T) {
	var items map[string]string = map[string]string{
		"go": "golang",
		"rb": "ruby",
		"js": "javascript",
	}
	var keys = [3]string{"go", "rb", "js"}

	if !AllKeys(keys).Match(items) {
		t.Fail()
	}
}

func TestAllKeys2(t *testing.T) {
	var items map[string]string = map[string]string{
		"go": "golang",
		"rb": "ruby",
		"js": "javascript",
	}
	var keys = []string{"go", "rb", "js"}

	if !AllKeys(keys).Match(items) {
		t.Fail()
	}
}
