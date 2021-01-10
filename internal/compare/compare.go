package compare

import (
	"reflect"
)

// EqualTo returns true if two values are equal.
func EqualTo(value1 interface{}, value2 interface{}) bool {
	return reflect.DeepEqual(value1, value2)
}

// LessThan returns true if value1 < value2.
func LessThan(value1 interface{}, value2 interface{}) bool {
	type1 := reflect.ValueOf(value1).Kind()
	type2 := reflect.ValueOf(value2).Kind()
	v1 := reflect.ValueOf(value1)
	v2 := reflect.ValueOf(value2)

	if isFloat(type1) || isFloat(type2) {
		return v1.Float() < v2.Float()
	} else if isInt(type1) || isInt(type2) {
		return v1.Int() < v2.Int()
	} else if isUint(type1) || isUint(type2) {
		return v1.Uint() < v2.Uint()
	} else if isString(type1) || isString(type2) {
		return v1.String() < v2.String()
	}

	return false
}

func isFloat(kind reflect.Kind) bool {
	return kind == reflect.Float32 || kind == reflect.Float64
}

func isInt(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32
}

func isUint(kind reflect.Kind) bool {
	return kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32
}

func isString(kind reflect.Kind) bool {
	return kind == reflect.String
}

// GreaterThan returns true if value1 > value2.
func GreaterThan(value1 interface{}, value2 interface{}) bool {
	return LessThan(value2, value1)
}
