package v

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/ladydascalie/v/validators"
)

func init() {
	validators.CustomFuncMap.Set("custom_function", func(args string, value, structure interface{}) error {
		tester, ok := structure.(*Tester)
		if !ok {
			return fmt.Errorf("expected type: %T, but got type: %T", &Tester{}, structure)
		}
		log.Printf("%#+v\n", tester)

		slice, ok := value.([]string)
		if !ok {
			errors.New("cannot handle non slice")
		}
		for _, item := range slice {
			switch item {
			case "a", "b", "c":
				// all good
			default:
				return fmt.Errorf("item %s did not match any of the expected values", item)
			}
		}
		return nil
	})
}

func makeTestableStruct() Tester {
	str := "hello world"
	tester := Tester{
		WrongTag:                  "a",
		MaxCharIn:                 str,    // should fail
		StringBetween:             str,    // should fail
		StringPointerBetween:      &str,   // should fail if commented out
		StringPointerBetween2:     &str,   // should not fail
		StringIn:                  "onee", // should fail
		StringBetweenNonZeroRange: "aaa",  // should fail because not enough char
		IntBetween:                51,
		IntIn:                     2,
		RequiredField: &SubStruct{
			F32Between:  12.12,   // should not fail
			F32Between2: 100.0,   // should fail
			f32Between3: 123.123, // should not trigger validation at all
		},
		RequiredSliceOfString: &[]string{"a", "b", "hello world"},
		CallOutsideFunc:       &[]string{"a", "b", "c", "d"}, // d is not wanted
	}
	return tester
}

type Tester struct {
	WrongTag                  string  `v:"gibberish"`
	MaxCharIn                 string  `json:"max_char_in" v:"maxchar:10, in:hello|world"`
	StringBetween             string  `v:"between:0..10"`
	StringPointerBetween      *string `json:"string_pointer_between" v:"required,between:0..10"`
	StringPointerBetween2     *string `json:"string_pointer_between_2" v:"required,between:0..10"`
	StringIn                  string  `v:"in:one|two|three"`
	StringBetweenNonZeroRange string  `v:"between:4..10"`

	IntBetween int `v:"between:0..50"`
	IntIn      int `v:"in:1|10|100"`

	RequiredField *SubStruct `v:"required"`

	RequiredSliceOfString *[]string `v:"required,maxchar:10,in:a|b|c"`
	CallOutsideFunc       *[]string `v:"func:custom_function"`
}

type SubStruct struct {
	F32Between  float32 `v:"between:0..99.9"`
	F32Between2 float32 `v:"between:0..99.9"`
	f32Between3 float32 `v:"between:0..99.9"`
}

func TestStruct(t *testing.T) {
	tester := makeTestableStruct()
	log.Println(Struct(&tester))
}
