package matcher

import "testing"

func TestLessThan(t *testing.T) {
	if !LessThan(100).Match(t, 99) {
		t.Fail()
	}
	if LessThan(100).Match(t, 100) {
		t.Fail()
	}
	if LessThan(100).Match(t, 101) {
		t.Fail()
	}
}

func TestLessThanOrEqualtTo(t *testing.T) {
	if !LessThanOrEqualTo(100).Match(t, 99) {
		t.Fail()
	}
	if !LessThanOrEqualTo(100).Match(t, 100) {
		t.Fail()
	}
	if LessThanOrEqualTo(100).Match(t, 101) {
		t.Fail()
	}
}
