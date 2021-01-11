package matcher

import (
	"testing"

	"github.com/keyhorn/assert/internal/testutils"
)

func TestNil(t *testing.T) {
	var items = []struct {
		actual interface{}
		result bool
	}{
		{actual: nil, result: true},
		{actual: "", result: false},
		{actual: make([]int, 1), result: false},
	}

	for _, item := range items {
		if Nil().Match(item.actual) != item.result {
			t.Errorf("[%v] Expected:<%v> Actual: <%v> Result: <%v>",
				"Nil",
				"nil",
				testutils.ToString(item.actual),
				testutils.ToString(item.result))
		}
	}
}
