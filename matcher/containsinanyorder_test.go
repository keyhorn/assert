package matcher

import "testing"

func TestContainsInAnyOrder1(t *testing.T) {
	if !ContainsInAnyOrder([]string {"b", "c", "a"}).Match([]string {"a", "b", "c"}) {
		t.Fail()
	}
	if !ContainsInAnyOrder([3]string {"b", "c", "a"}).Match([3]string {"a", "b", "c"}) {
		t.Fail()
	}
}

func TestContainsInAnyOrder2(t *testing.T) {
	if ContainsInAnyOrder([]string {"b", "c", "A"}).Match([]string {"a", "b", "c"}) {
		t.Fail()
	}
	if ContainsInAnyOrder([3]string {"b", "c", "A"}).Match([3]string {"a", "b", "c"}) {
		t.Fail()
	}
}

func TestContainsInAnyOrder3(t *testing.T) {
	if ContainsInAnyOrder([]string {"a", "b", "c"}).Match("a") {
		t.Fail()
	}
	if ContainsInAnyOrder([3]string {"a", "b", "c"}).Match("a") {
		t.Fail()
	}
}