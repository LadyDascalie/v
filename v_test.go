package v

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

		"github.com/ladydascalie/v/validators"
)

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	validators.CustomFuncMap.Set("custom_function", func(args string, value, structure interface{}) error {
		return errors.New("custom_function was called")
	})
	os.Exit(m.Run())
}

func TestSetGet(t *testing.T) {
	Set("test", func(args string, value, structure interface{}) error {
		return nil
	})
	_, ok := Get("test")
	if !ok {
		t.Fail()
	}
}

func TestStruct(t *testing.T) {
	tests := []struct {
		name      string
		structure interface{}
		wantErr   bool
	}{
		{
			name:      "test nil",
			structure: nil,
			wantErr:   false,
		},
		{
			name: "test unexported",
			structure: struct {
				field string
			}{},
			wantErr: false,
		},
		{
			name:      "test non struct",
			structure: []int{},
			wantErr:   true,
		},
		{
			name: "required without json name",
			structure: struct {
				Field *string `v:"required"`
			}{},
			wantErr: true,
		},
		{
			name: "test required",
			structure: struct {
				NonNullableField string   `v:"required" json:"non_nullable_field"`
				Ptr              *string  `v:"required" json:"ptr"`
				Ptr2             *string  `v:"required"`
				Channel          chan int `v:"required" json:"channel"`
				Func             func()   `v:"required" json:"func"`
				Slice            []int    `v:"required" json:"slice"`
			}{
				NonNullableField: "",
				Ptr:              nil,
				Channel:          nil,
				Func:             nil,
				Slice:            nil,
			},
			wantErr: true,
		},
		{
			name: "test maxchar",
			structure: struct {
				TooLong   string `v:"maxchar:10"` // should error
				Ok        string `v:"maxchar:11"` // should not error
				WrongType int    `v:"maxchar:10"` // shoult not error
			}{
				TooLong: "Hello World",
				Ok:      "Hello World",
			},
			wantErr: true,
		},
		{
			name: "match email",
			structure: struct {
				Email string `v:"matches:email"`
			}{
				Email: "someone@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "test custom function",
			structure: struct {
				Field string `v:"func:custom_function"`
			}{},
			wantErr: true,
		},
		{
			name: "missing custom function",
			structure: struct {
				Field string `v:"func:custom_function_oops" json:"field"`
			}{},
			wantErr: true,
		},
		{
			name: "with sub struct",
			structure: struct {
				S struct {
					Field string `v:"func:custom_function"`
				}
			}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Struct(tt.structure); (err != nil) != tt.wantErr {
				t.Errorf("Struct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Embedding(t *testing.T) {
	type A struct {
		StrField string `v:"between:0..10"`
	}
	type B struct {
		A
		IntField int `v:"between:0..10"`
	}
	err := Struct(B{
		A: A{
			StrField: "hello world!", // this is not ok
		},
		IntField: 10, // this is ok
	})
	if err == nil {
		t.Error("This should fail")
	}
}

func Test_validationErorrs_Error(t *testing.T) {
	tests := []struct {
		name    string
		v       validationErorrs
		wantErr bool
	}{
		{
			name:    "slice has errors",
			v:       []error{fmt.Errorf("some error")},
			wantErr: true,
		},
		{
			name:    "slice is empty",
			v:       nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.Error(); (err != nil) != tt.wantErr {
				t.Errorf("validationErorrs.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		tag       string
		value     interface{}
		structure interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid tag",
			args: args{
				tag:   "invalid",
				value: "hello",
				structure: struct {
					Field string
				}{},
			},
			wantErr: true,
		},
		{
			name: "invalid tag",
			args: args{
				tag:   "invalid:invalid:invalid",
				value: "hello",
				structure: struct {
					Field string
				}{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate(tt.args.tag, tt.args.value, tt.args.structure); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
