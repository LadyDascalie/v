package v

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/ladydascalie/v/validators"
)

const (
	tagname = "v"
	jsontag = "json"

	// common tags
	required = "required"
)

// Set a new validator into the custom func map
func Set(tag string, validator validators.Validator) {
	c := validators.GetFuncMap()
	c.Set(tag, validator)
}

// Get a validator from the custom func map
func Get(tag string) (validator validators.Validator, ok bool) {
	c := validators.GetFuncMap()
	return c.Get(tag)
}

// Struct takes in an interface, which must be a struct
// all validation is ran based on the provided tags.
func Struct(structure interface{}) error {
	// nothing to see here
	if structure == nil {
		return nil
	}

	// ensure we're ok even if passed a pointer
	v := reflect.Indirect(reflect.ValueOf(structure))
	if v.Kind() != reflect.Struct {
		return errors.New("only structs may be passed to this method")
	}

	t := v.Type() // get the struct type
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // prepare the field
		value := v.Field(i) // prepare the field value

		// retrieve the underlying value if possible
		if value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
			value = value.Elem()
		}

		// recurse if this is an embedded struct
		if value.Kind() == reflect.Struct && field.PkgPath == "" {
			// only exported fields should do this
			if err := Struct(value.Interface()); err != nil {
				return err
			}
		}

		// get all the v tags
		tags := field.Tag.Get(tagname)

		// get the json tags
		jtag := field.Tag.Get(jsontag)

		// split the v tags on comma
		vtags := strings.Split(tags, ",")

		// validation errors collection
		var vErrors validationErorrs

		// range over the tags
		for _, vtag := range vtags {
			if err := handleValidationTag(vtag, jtag, field, value, structure); err != nil {
				vErrors = append(vErrors, err)
			}
		}
		if len(vErrors) != 0 {
			return vErrors.Error()
		}
	}
	return nil
}

func handleValidationTag(vtag, jtag string, field reflect.StructField, value reflect.Value, structure interface{}) (err error) {
	// sanitize the tag. when multiple tags are used
	// some leading/trailing spaces may be left
	vtag = strings.TrimSpace(vtag)

	// guard against unexported fields
	// or simply missing or invalid tags
	if field.PkgPath != "" || vtag == "" {
		return
	}

	// is the field required but invalid?
	// this will trigger for instance on a *string
	// which has not been initialized.
	if !value.IsValid() && vtag == required {
		return ErrorRequired{
			Field:    field.Name,
			JSONName: jtag,
		}
	}
	// Our field is valid, and we can interface without panic
	// we are ready to send it to the validator methods
	if value.IsValid() && value.CanInterface() {
		if err = validate(vtag, value.Interface(), structure); err != nil {
			return ErrorValidation{
				Name:     field.Name,
				JSONName: jtag,
				Err:      err,
			}
		}
	}
	return
}

func validate(tag string, value, structure interface{}) error {
	vtag := newValidationTag(tag)
	if vtag == nil {
		return fmt.Errorf("v cannot parse struct tag <%v> please refer to the format rules", tag)
	}

	// first check for custom functions
	if vtag.Name == "func" {
		fn, ok := validators.CustomFuncMap.Get(vtag.Args)
		if ok {
			return fn(vtag.Args, value, structure)
		}
		return fmt.Errorf("custom validator %s did not match any available function", vtag.Args)
	}

	// run through the func map and see if there's a match
	for name, method := range validators.FuncMap {
		if vtag.Name == name {
			return method(vtag.Args, value)
		}
	}
	return fmt.Errorf("could not parse validation tag: %s", vtag.Name)
}

type validationTag struct {
	Name string
	Args string
}

func newValidationTag(tag string) *validationTag {
	var vtag validationTag

	parts := strings.SplitN(tag, ":", -1)
	switch len(parts) {
	case 1:
		vtag.Name = parts[0]
	case 2:
		vtag.Name = parts[0]
		vtag.Args = parts[1]
	default:
		return nil
	}
	return &vtag
}
