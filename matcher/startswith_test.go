package matcher

import (
	"testing"

	"github.com/keyhorn/assert/internal/testutils"
)

func TestStartsWith(t *testing.T) {
	var items = []struct {
		actual   interface{}
		expected string
		result   bool
	}{
		{actual: "hello world", expected: "hello", result: true},
		{actual: "hello world", expected: "world", result: false},
		{actual: 1, expected: "hello", result: false},
	}

	for _, item := range items {
		if StartsWith(item.expected).Match(item.actual) != item.result {
			t.Errorf("[%v] Expected:<%v> Actual: <%v> Result: <%v>",
				"StartsWith",
				testutils.ToString(item.expected),
				testutils.ToString(item.actual),
				testutils.ToString(item.result))
		}
	}
}
