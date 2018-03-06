package validators

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/ladydascalie/v/convert"
	"github.com/ladydascalie/v/sanity"
)

// FuncMap defines where all the validator live.
// This is also where you need to add any custom validator that you need.
var FuncMap = map[string]func(args string, value interface{}) error{
	"maxchar":  Maxchar,
	"in":       In,
	"between":  Between,
	"required": Required,
}

// Required checks that the nullable type is in not nil
func Required(args string, value interface{}) error {
	if sanity.IsNullable(value) && reflect.ValueOf(value).IsNil() {
		return fmt.Errorf("required, please provide a value")
	}
	return nil
}

// In checks if the provided value is contained within the provided arguments
func In(args string, value interface{}) error {
	accepted := strings.Split(args, "|")

	switch {
	case sanity.IsString(value):
		nv := value.(string)
		if !strIn(nv, accepted) {
			return fmt.Errorf("accepted values are: [%s], but got: %s", strings.Join(accepted, ", "), nv)
		}
		return nil
	case sanity.IsStringSlice(value):
		nv := value.([]string)
		for _, item := range nv {
			if !strIn(item, accepted) {
				return fmt.Errorf("accepted values are: [%s], but got: %s", strings.Join(accepted, ", "), nv)
			}
		}
		return nil
	case sanity.IsNumeric(value):
		nv, err := convert.ToFloat64(value)
		if err != nil {
			return err
		}
		values, err := stringSliceToFloatSlice(accepted)
		if err != nil {
			return err
		}
		if !floatIn(nv, values) {
			return fmt.Errorf("accepted values are: [%s], but got: %s", strings.Join(accepted, ", "), f64(nv))
		}
		return nil
	}

	return nil
}

// Maxchar checks if the value's length in characters is constrained by the argument's value
func Maxchar(args string, value interface{}) error {
	max, err := strconv.Atoi(args)
	if err != nil {
		return err
	}
	switch v := value.(type) {
	case string:
		count := utf8.RuneCountInString(v)
		if count > max {
			return fmt.Errorf("expected maximum %d characters, got: %d", max, count)
		}
	case []string:
		var count int
		for _, item := range v {
			count = utf8.RuneCountInString(item)
			if count > max {
				return fmt.Errorf("items have an expected maximum %d characters, got: %d on value: %s", max, count, item)
			}
		}
	default:
		return fmt.Errorf("expected value of type string, but got %T", v)
	}
	return nil
}

// Between checks if the provided value is constrained by the argument's bounds
func Between(args string, value interface{}) error {
	if _, ok := value.(string); ok {
		return length(args, value)
	}
	min, max, err := bounds(args)
	if err != nil {
		return err
	}
	nv, err := convert.ToFloat64(value)
	if err != nil {
		return err
	}

	if nv < min || nv > max {
		return fmt.Errorf("expected a value between %s and %s, but got %s", f64(min), f64(max), f64(nv))
	}
	return nil
}

/*--------+
| helpers |
+--------*/

func bounds(s string) (min, max float64, err error) {
	parts := strings.Split(s, "..")
	if len(parts) != 2 {
		return min, max, fmt.Errorf("invalid range statement: %v", s)
	}
	min, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return
	}
	max, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return
	}
	return
}

func length(args string, value interface{}) error {
	str := value.(string) // already checked by between
	min, max, err := bounds(args)
	if err != nil {
		return err
	}
	count := float64(utf8.RuneCountInString(str))
	if count < min || count > max {
		return fmt.Errorf("expected a value between %s and %s, but got %s", f64(min), f64(max), f64(count))
	}
	return nil
}

func strIn(str string, values []string) bool {
	for _, v := range values {
		if v == str {
			return true
		}
	}
	return false
}

func floatIn(f64 float64, floats []float64) bool {
	for _, f := range floats {
		if f == f64 {
			return true
		}
	}
	return false
}

func stringSliceToFloatSlice(values []string) (floats []float64, err error) {
	var f float64
	for _, v := range values {
		f, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("value %v cannot is not numeric", v)
		}
		floats = append(floats, f)
	}
	return
}

// f64 formats a float value for human reading.
// ex: 100.0 -> 100 or 100.123000 -> 100.123
func f64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}