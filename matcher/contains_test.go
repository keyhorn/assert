package matcher

import (
	"testing"

	"github.com/keyhorn/assert/internal/testutils"
)

func TestContains(t *testing.T) {
	var items = []struct {
		actual   interface{}
		expected string
		result   bool
	}{
		{actual: "hello world", expected: "wor", result: true},
		{actual: "hello world", expected: "wrd", result: false},
		{actual: 1, expected: "hello", result: false},
	}

	for _, item := range items {
		if Contains(item.expected).Match(item.actual) != item.result {
			t.Errorf("[%v] Expected:<%v> Actual: <%v> Result: <%v>",
				"EndsWith",
				testutils.ToString(item.expected),
				testutils.ToString(item.actual),
				testutils.ToString(item.result))
		}
	}
}
