package sanity

import "reflect"

const (
	Numeric = iota
	String
	Unknown
)

func Check(value interface{}) int {
	if IsNumeric(value) {
		return Numeric
	}
	if IsString(value) {
		return String
	}
	return Unknown
}

func Nullable(value interface{}) bool {
	switch reflect.ValueOf(value).Kind() {
	case reflect.Ptr, reflect.Slice:
		return true
	default:
		return false
	}
}

func IsNumeric(value interface{}) bool {
	switch value.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return true
	default:
		return false
	}
}

func IsString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}
