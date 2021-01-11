package matcher

import (
	"testing"

	"github.com/keyhorn/assert/internal/testutils"
)

func TestEmpty(t *testing.T) {
	var items = []struct {
		actual interface{}
		result bool
	}{
		{actual: nil, result: true},
		{actual: "", result: true},
		{actual: make([]int, 0), result: true},
		{actual: make([]int, 1), result: false},
		{actual: 1, result: false},
	}

	for _, item := range items {
		if Empty().Match(t, item.actual) != item.result {
			t.Errorf("[%v] Expected:<%v> Actual: <%v> Result: <%v>",
				"Empty",
				"empty",
				testutils.ToString(item.actual),
				testutils.ToString(item.result))
		}
	}
}
