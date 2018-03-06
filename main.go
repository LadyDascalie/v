package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/ladydascalie/v/validators"
)

const (
	tagname = "v"
	jsontag = "json"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tester := Tester{
		UserID: "hello   world", // len 11
		AA:     "hello world",   // should fail
		AAA:    "onee",          // should fail
		AAAA:   "aaa",           // should fail because not enough char
		B:      51,
		BB:     2,
		C: &Point{
			X: 12.12,   // should not fail
			Y: 100.0,   // should fail
			z: 123.123, // should not trigger validation at all
		},
		D:        []string{"a", "b", "c"},
		Nestable: &Nestable{},
	}

	log.Println(Struct(tester))
}

type Tester struct {
	UserID string `json:"user_id" v:"maxchar:10, in:hello|world"`
	AA     string `v:"between:0..10"`
	AAA    string `v:"in:one|two|three"`
	AAAA   string `v:"between:4..10"`

	B  int `v:"between:0..50"`
	BB int `v:"in:1|10|100"`

	C *Point `v:"required"`

	D []string `v:"required,maxchar:10"`

	Nestable *Nestable
}

type Point struct {
	X, Y, z float32 `v:"between:0..99.9"`
}

type Nestable struct {
	Nested bool //`v:""`
}

func Struct(s interface{}) error {
	// ensure we're ok even if passed a pointer
	v := reflect.Indirect(reflect.ValueOf(s))
	if v.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	t := v.Type() // get the struct type
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		if value.Kind() == reflect.Struct {
			Struct(value.Interface()) // recurse if this is an embeded struct
		}

		tags := field.Tag.Get(tagname)    // get all v tags
		vtags := strings.Split(tags, ",") // split them up
		jtag := field.Tag.Get(jsontag)    // get the json tag if any

		// range over the tags
		for _, vtag := range vtags {
			vtag = strings.TrimSpace(vtag)         // clean the tag
			if field.PkgPath == "" && vtag != "" { // guard against unexported fields
				if err := validate(vtag, value.Interface()); err != nil {
					log.Println(validationError(field.Name, jtag, err))
					// return validationError(vtag, jtag, err)
				}
			}
		}
	}
	return nil
}

func validationError(name, jtag string, err error) error {
	if jtag != "" {
		return fmt.Errorf("[validation] %s: %v", jtag, err)
	}
	return fmt.Errorf("[validation] %s: %v", name, err)
}

func validate(tag string, value interface{}) error {
	vtag := NewVTag(tag)
	if vtag == nil {
		return fmt.Errorf("v cannot parse struct tag <%v> please refer to the format rules", tag)
	}
	switch vtag.Name {
	case "maxchar":
		return validators.Maxchar(vtag.Args, value)
	case "between":
		return validators.Between(vtag.Args, value)
	case "in":
		return validators.In(vtag.Args, value)
	}
	return nil
}

type VTag struct {
	Name string
	Args string
}

func NewVTag(tag string) *VTag {
	var vtag VTag

	parts := strings.SplitN(tag, ":", -1)
	switch len(parts) {
	case 1:
		vtag.Name = parts[0]
	case 2:
		vtag.Name = parts[0]
		vtag.Args = parts[1]
	}
	return &vtag
}
