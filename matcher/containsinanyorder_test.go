package matcher

import "testing"

func TestContainsInAnyOrder1(t *testing.T) {
	if !ContainsInAnyOrder([]string {"b", "c", "a"}).Match(t, []string {"a", "b", "c"}) {
		t.Fail()
	}
	if !ContainsInAnyOrder([3]string {"b", "c", "a"}).Match(t, [3]string {"a", "b", "c"}) {
		t.Fail()
	}
}

func TestContainsInAnyOrder2(t *testing.T) {
	if ContainsInAnyOrder([]string {"b", "c", "A"}).Match(t, []string {"a", "b", "c"}) {
		t.Fail()
	}
	if ContainsInAnyOrder([3]string {"b", "c", "A"}).Match(t, [3]string {"a", "b", "c"}) {
		t.Fail()
	}
}

func TestContainsInAnyOrder3(t *testing.T) {
	if ContainsInAnyOrder([]string {"a", "b", "c"}).Match(t, "a") {
		t.Fail()
	}
	if ContainsInAnyOrder([3]string {"a", "b", "c"}).Match(t, "a") {
		t.Fail()
	}
}