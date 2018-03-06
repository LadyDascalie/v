package sanity

import "reflect"

const (
	Numeric = iota
	String
	Unknown
)

// IsNullable checks if a given value is nullable
func IsNullable(value interface{}) bool {
	switch reflect.ValueOf(value).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	default:
		return false
	}
}

// IsNumeric checks if a given value is numeric
func IsNumeric(value interface{}) bool {
	switch value.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return true
	default:
		return false
	}
}

// IsString checks if a given value is a string
func IsString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

// IsStringSlice checks it a given value is a slice of strings
func IsStringSlice(value interface{}) bool {
	_, ok := value.([]string)
	return ok
}
