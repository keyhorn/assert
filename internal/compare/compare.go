package compare

import "reflect"

// EqualTo returns true if two values are equal.
func EqualTo(value1 interface{}, value2 interface{}) bool {
	return reflect.DeepEqual(value1, value2)
}

// LessThan returns true if value1 < value2.
func LessThan(value1 interface{}, value2 interface{}) bool {
	v1 := reflect.ValueOf(value1)
	v2 := reflect.ValueOf(value2)
	switch value1.(type) {
	case float32, float64:
		switch value2.(type) {
		case float32, float64:
			return v1.Float() < v2.Float()
		case int, int8, int16, int32, int64:
			return v1.Float() < v2.Float()
		case uint, uint8, uint16, uint32, uint64:
			return v1.Float() < v2.Float()
		default:
			return false
		}
	case int, int8, int16, int32, int64:
		switch value2.(type) {
		case float32, float64:
			return v1.Float() < v2.Float()
		case int, int8, int16, int32, int64:
			return v1.Int() < v2.Int()
		case uint, uint8, uint16, uint32, uint64:
			return v1.Int() < v2.Int()
		default:
			return false
		}
	case uint, uint8, uint16, uint32, uint64:
		switch value2.(type) {
		case float32, float64:
			return v1.Float() < v2.Float()
		case int, int8, int16, int32, int64:
			return v1.Int() < v2.Int()
		case uint, uint8, uint16, uint32, uint64:
			return v1.Uint() < v2.Uint()
		default:
			return false
		}
	case string:
		return v1.String() < v2.String()
	}
	return false
}

// GreaterThan returns true if value1 > value2.
func GreaterThan(value1 interface{}, value2 interface{}) bool {
	return LessThan(value2, value1)
}
