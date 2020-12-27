package matcher

import (
	"testing"

	"github.com/keyhorn/assert/internal/testutils"
)

func TestEqualTo(t *testing.T) {
	var items = []struct {
		actual   interface{}
		expected interface{}
		result   bool
	}{
		{actual: int(1), expected: int(1), result: true},
		{actual: int8(1), expected: int8(1), result: true},
		{actual: int16(1), expected: int16(1), result: true},
		{actual: int32(1), expected: int32(1), result: true},
		{actual: int64(1), expected: int64(1), result: true},
		{actual: uint(1), expected: uint(1), result: true},
		{actual: uint8(1), expected: uint8(1), result: true},
		{actual: uint16(1), expected: uint16(1), result: true},
		{actual: uint32(1), expected: uint32(1), result: true},
		{actual: uint64(1), expected: uint64(1), result: true},
		{actual: float32(1.1), expected: float32(1.1), result: true},
		{actual: float64(1.1), expected: float64(1.1), result: true},
		{actual: true, expected: true, result: true},
		{actual: false, expected: false, result: true},
		{actual: true, expected: false, result: false},
		{actual: false, expected: true, result: false},
		{actual: "a", expected: "a", result: true},
		{actual: "a", expected: "A", result: false},
		{actual: "B", expected: "B", result: true},
		{actual: "B", expected: "b", result: false},
	}

	for _, item := range items {
		if EqualTo(item.expected).Match(item.actual) != item.result {
			t.Errorf("[%v] Expected:<%v> Actual: <%v> Result: <%v>",
				"EqualTo",
				testutils.ToString(item.expected),
				testutils.ToString(item.actual),
				testutils.ToString(item.result))
		}
	}
}
