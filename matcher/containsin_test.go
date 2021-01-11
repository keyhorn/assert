package matcher

import "testing"

func TestContainsIn1(t *testing.T) {
	if !ContainsIn([]string {"a", "b", "c"}).Match(t, []string {"a", "b", "c"}) {
		t.Fail()
	}
	if !ContainsIn([3]string {"a", "b", "c"}).Match(t, [3]string {"a", "b", "c"}) {
		t.Fail()
	}
}

func TestContainsIn2(t *testing.T) {
	if ContainsIn([]string {"A", "b", "c"}).Match(t, []string {"a", "b", "c"}) {
		t.Fail()
	}
	if ContainsIn([3]string {"A", "b", "c"}).Match(t, [3]string {"a", "b", "c"}) {
		t.Fail()
	}
}

func TestContainsIn3(t *testing.T) {
	if ContainsIn([]string {"a", "b", "c"}).Match(t, "a") {
		t.Fail()
	}
	if ContainsIn([3]string {"a", "b", "c"}).Match(t, "a") {
		t.Fail()
	}
}

func TestContainsIn4(t *testing.T) {
	if ContainsIn([]string {"b", "c", "a"}).Match(t, []string {"a", "b", "c"}) {
		t.Fail()
	}
	if ContainsIn([3]string {"b", "c", "a"}).Match(t, [3]string {"a", "b", "c"}) {
		t.Fail()
	}
}