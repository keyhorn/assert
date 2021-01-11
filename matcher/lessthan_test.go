package matcher

import "testing"

func TestLessThan(t *testing.T) {
	if !LessThan(100).Match(99) {
		t.Fail()
	}
	if LessThan(100).Match(100) {
		t.Fail()
	}
	if LessThan(100).Match(101) {
		t.Fail()
	}
}

func TestLessThanOrEqualtTo(t *testing.T) {
	if !LessThanOrEqualTo(100).Match(99) {
		t.Fail()
	}
	if !LessThanOrEqualTo(100).Match(100) {
		t.Fail()
	}
	if LessThanOrEqualTo(100).Match(101) {
		t.Fail()
	}
}
