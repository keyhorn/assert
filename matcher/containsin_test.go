package matcher

import "testing"

func TestContainsIn1(t *testing.T) {
	if !ContainsIn([]string {"a", "b", "c"}).Match([]string {"a", "b", "c"}) {
		t.Fail()
	}
	if !ContainsIn([3]string {"a", "b", "c"}).Match([3]string {"a", "b", "c"}) {
		t.Fail()
	}
}

func TestContainsIn2(t *testing.T) {
	if ContainsIn([]string {"A", "b", "c"}).Match([]string {"a", "b", "c"}) {
		t.Fail()
	}
	if ContainsIn([3]string {"A", "b", "c"}).Match([3]string {"a", "b", "c"}) {
		t.Fail()
	}
}

func TestContainsIn3(t *testing.T) {
	if ContainsIn([]string {"a", "b", "c"}).Match("a") {
		t.Fail()
	}
	if ContainsIn([3]string {"a", "b", "c"}).Match("a") {
		t.Fail()
	}
}

func TestContainsIn4(t *testing.T) {
	if ContainsIn([]string {"b", "c", "a"}).Match([]string {"a", "b", "c"}) {
		t.Fail()
	}
	if ContainsIn([3]string {"b", "c", "a"}).Match([3]string {"a", "b", "c"}) {
		t.Fail()
	}
}