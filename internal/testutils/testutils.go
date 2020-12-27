package testutils

import "fmt"

// ToString returns the string representation of the value.
func ToString(value interface{}) string {
	var result string
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		result = fmt.Sprintf("%d", value)
	case float32, float64:
		result = fmt.Sprintf("%f", value)
	case complex64, complex128:
		result = fmt.Sprintf("%e", value)
	default:
		result = fmt.Sprintf("%v", value)
	}
	return result
}
