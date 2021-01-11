package matcher

import "testing"

func TestGreaterThan(t *testing.T) {
	if GreaterThan(100).Match(t, 99) {
		t.Fail()
	}
	if GreaterThan(100).Match(t, 100) {
		t.Fail()
	}
	if !GreaterThan(100).Match(t, 101) {
		t.Fail()
	}
}

func TestGreaterThanOrEqualtTo(t *testing.T) {
	if GreaterThanOrEqualTo(100).Match(t, 99) {
		t.Fail()
	}
	if !GreaterThanOrEqualTo(100).Match(t, 100) {
		t.Fail()
	}
	if !GreaterThanOrEqualTo(100).Match(t, 101) {
		t.Fail()
	}
}
