# `v` - It validates.

[![Go Report Card](https://goreportcard.com/badge/github.com/ladydascalie/v)](https://goreportcard.com/report/github.com/ladydascalie/v)

## Usage:
```go
package main

import (
	"log"

	"github.com/ladydascalie/v"
)

type Person struct {
	FirstName   string `v:"maxchar:255"`
	LastName    string `v:"maxchar:255"`
	PhoneNumber string `v:"between:0..9"`
	Age         int    `v:"between:21..*"`
	// wilcard syntax means math.MaxFloat64 will be used here.
	// if the wilcard was on the left side, this would have been
	// -math.MaxFloat64
}

func main() {
	p1 := Person{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "123456789",
		Age:         16,
	}

	if err := v.Struct(p1); err != nil {
		log.Println(err)
	}
}

// Output: Age: expected a value between 21 and 1.7976931348623157e+308, but got 16
```


## The `FuncMap`:

These are the built-in validators provided by `v`

```go
var FuncMap = map[string]func(args string, value interface{}) error{
	"required":      Required,
	"maxchar":       Maxchar,
	"in":            In,
	"between":       Between,
	"bytes_between": BytesBetween,
	"empty_string":  EmptyString,
	"is_int64":      IsInt64,
	"is_float64":    IsFloat64,
	"matches":       Matches,
}
```

### Custom Validators

You may add custom validators to `v`, an `init` method is a very good time to do this:

```go
func init() {
	v.Set("custom_function", func(args string, value, structure interface{}) error {
		// do some validation
		return nil
	})
}
```

Then simply set the validation tag like so:

```go
type A struct {
	Name string `v:"func:custom_function"`
}
```

That's all it takes.

### About RegExp

The `RegExp` that `matches` provides are taken from [govalidator](https://github.com/asaskevich/govalidator). Here is the complete list of them:

```
- email
- credit_card
- isbn10
- isbn13
- uuid3
- uuid4
- uuid5
- uuid
- alpha
- alphanum
- numeric
- int
- float
- hexadecimal
- hex_color
- rgb_color
- ascii
- printable_ascii
- multi_byte
- full_width
- half_width
- base64
- data_uri
- latitude
- longitude
- dns_name
- url
- ssn
- win_path
- unix_path
- semver
- has_lowercase
- has_uppercase
```
