package matcher

import "testing"

func TestHasEntry1(t *testing.T) {
	languages := map[string]string{
		"go": "golang", 
		"rb": "ruby", 
		"js": "javascript",
	}

	if !HasEntry(EqualTo("js"), EqualTo("javascript")).Match(t, languages) {
		t.Fail()
	}

}

func TestHasEntry2(t *testing.T) {
	languages := map[string]string{
		"go": "golang", 
		"rb": "ruby", 
		"js": "javascript",
	}

	if HasEntry(EqualTo("js"), EqualTo("typescript")).Match(t, languages) {
		t.Fail()
	}

}

func TestHasEntry3(t *testing.T) {
	languages := map[string]string{
		"go": "golang", 
		"rb": "ruby", 
		"js": "javascript",
	}

	if HasEntry(EqualTo("ts"), EqualTo("javascript")).Match(t, languages) {
		t.Fail()
	}

}

func TestHasEntry4(t *testing.T) {
	if HasKey(EqualTo("js")).Match(t, "js") {
		t.Fail()
	}
}