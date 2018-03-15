package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToFloat64 tries it's best to convert anything into a float64
func ToFloat64(value interface{}) (f float64, err error) {
	switch v := value.(type) {
	case float64:
		return v, nil
		// in all those cases we convert to float64 first
	case int, int8, int16, int32, int64, float32:
		return reflect.ValueOf(v).Convert(reflect.TypeOf(f)).Interface().(float64), nil
	case uint, uint8, uint16, uint32, uint64: // run the uint cases last as they are the rarest.
		return reflect.ValueOf(v).Convert(reflect.TypeOf(f)).Interface().(float64), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case []byte:
		return strconv.ParseFloat(string(v), 64)
	default:
		return f, fmt.Errorf("%[1]v (%[1]T) cannot be converted to float64", value)
	}
}

// ToInt64 tries it's best to convert anything into an int64
func ToInt64(value interface{}) (f int64, err error) {
	switch v := value.(type) {
	case int64:
		return v, nil
	case int, int8, int16, int32, float32, float64:
		return reflect.ValueOf(v).Convert(reflect.TypeOf(f)).Interface().(int64), nil
	case uint, uint8, uint16, uint32, uint64: // run the uint cases last as they are the rarest.
		return reflect.ValueOf(v).Convert(reflect.TypeOf(f)).Interface().(int64), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	case []byte:
		return strconv.ParseInt(string(v), 10, 64)
	default:
		return f, fmt.Errorf("%[1]v (%[1]T) cannot be converted to int64", value)
	}
}
