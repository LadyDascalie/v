package validators

import (
	"math"
	"sync"
	"testing"
)

func TestGetFuncMap(t *testing.T) {
	fm := GetFuncMap()
	if fm != CustomFuncMap {
		t.Fail()
	}
}

func Test_customFuncMap_Set_Get(t *testing.T) {
	type args struct {
		tag       string
		validator Validator
	}
	tests := []struct {
		name string
		c    *customFuncMap
		args args
	}{
		{
			name: "",
			c: &customFuncMap{
				rw:         sync.RWMutex{},
				validators: make(map[string]Validator),
			},
			args: args{
				tag: "hello",
				validator: func(args string, value, structure interface{}) error {
					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Set(tt.args.tag, tt.args.validator)
			_, ok := tt.c.Get(tt.args.tag)
			if !ok {
				t.Fatalf("expected to find function %s in FuncMap", tt.args.tag)
			}
		})
	}
}

func TestRequired(t *testing.T) {
	var p *string
	var f func()
	var c chan int
	var s []int
	var m map[int]int

	type args struct {
		in0   string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "should not fail", args: args{"", ""}},
		{name: "should fail pointer", args: args{"", p}, wantErr: true},
		{name: "should fail func", args: args{"", f}, wantErr: true},
		{name: "should fail channel", args: args{"", c}, wantErr: true},
		{name: "should fail slice", args: args{"", s}, wantErr: true},
		{name: "should fail map", args: args{"", m}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Required(tt.args.in0, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Required() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIn(t *testing.T) {
	type args struct {
		args  string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should not fail (int)",
			args: args{
				args:  "0|5|10",
				value: 5,
			},
			wantErr: false,
		},
		{
			name: "should fail (int)",
			args: args{
				args:  "0|10",
				value: 11,
			},
			wantErr: true,
		},
		{
			name: "should not fail (string)",
			args: args{
				args:  "hello|world",
				value: "hello",
			},
			wantErr: false,
		},
		{
			name: "should fail (string)",
			args: args{
				args:  "hello",
				value: "hello world",
			},
			wantErr: true,
		},
		{
			name: "should not fail (string slice)",
			args: args{
				args:  "hello|world",
				value: []string{"hello", "world"},
			},
			wantErr: false,
		},
		{
			name: "should fail (string slice)",
			args: args{
				args:  "hello|world",
				value: []string{"hello world"},
			},
			wantErr: true,
		},
		{
			name: "should fail (slice of int)",
			args: args{
				args:  "1|2",
				value: []int{1, 3, 4},
			},
			wantErr: true,
		},
		{
			name: "invalid parameters for numeric in",
			args: args{
				args:  "hello|world",
				value: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := In(tt.args.args, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("In() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMaxchar(t *testing.T) {
	type args struct {
		args  string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid parameter",
			args: args{
				args:  "hello",
				value: "hello",
			},
			wantErr: true,
		},
		{
			name: "invalid type",
			args: args{
				args:  "10",
				value: 123,
			},
			wantErr: true,
		},
		{
			name: "should not fail",
			args: args{
				args:  "10",
				value: "hello",
			},
		},
		{
			name: "should fail",
			args: args{
				args:  "10",
				value: "hello world",
			},
			wantErr: true,
		},
		{
			name: "should not fail (slice of string)",
			args: args{
				args:  "10",
				value: []string{"hello", "world"},
			},
		},
		{
			name: "should fail (slice of string)",
			args: args{
				args:  "10",
				value: []string{"hello world"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Maxchar(tt.args.args, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Maxchar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBytesBetween(t *testing.T) {
	type args struct {
		args  string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "broken range",
			args: args{
				args:  "0..",
				value: "hello",
			},
			wantErr: true,
		},
		{
			name: "invalid parameter",
			args: args{
				args:  "0..10",
				value: []int{1, 2, 3},
			},
			wantErr: true,
		},
		{
			name: "should not fail",
			args: args{
				args:  "0..10",
				value: "hello",
			},
			wantErr: false,
		},
		{
			name: "should fail (too long)",
			args: args{
				args:  "0..10",
				value: "hello world",
			},
			wantErr: true,
		},
		{
			name: "should not fail (multi-byte characters)",
			args: args{
				args:  "0..3",
				value: "愛",
			},
			wantErr: false,
		},
		{
			name: "should not fail (multi-byte characters, closed range)",
			args: args{
				args:  "3..3",
				value: "愛",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BytesBetween(tt.args.args, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("BytesBetween() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	type args struct {
		args  string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "broken range",
			args: args{
				args:  "0..",
				value: 123,
			},
			wantErr: true,
		},
		{
			name: "invalid parameter",
			args: args{
				args:  "0..10",
				value: []int{1, 2, 3},
			},
			wantErr: true,
		},
		{
			name: "should not fail",
			args: args{
				args:  "0..10",
				value: "hello",
			},
			wantErr: false,
		},
		{
			name: "should fail (too long)",
			args: args{
				args:  "0..10",
				value: "hello world",
			},
			wantErr: true,
		},
		{
			name: "should not fail (int)",
			args: args{
				args:  "0..10",
				value: 5,
			},
			wantErr: false,
		},
		{
			name: "should fail (int too big)",
			args: args{
				args:  "0..10",
				value: 11,
			},
			wantErr: true,
		},
		{
			name: "should not fail (multi-byte characters)",
			args: args{
				args:  "0..1",
				value: "愛",
			},
			wantErr: false,
		},
		{
			name: "should not fail (multi-byte characters, closed range)",
			args: args{
				args:  "1..1",
				value: "愛",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Between(tt.args.args, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Between() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmptyString(t *testing.T) {
	type args struct {
		in0   string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should not fail",
			args: args{
				value: "",
			},
			wantErr: false,
		},
		{
			name: "should fail",
			args: args{
				value: "has text",
			},
			wantErr: true,
		},
		{
			name: "invalid type",
			args: args{
				in0:   "",
				value: 123,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EmptyString(tt.args.in0, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("EmptyString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsInt64(t *testing.T) {
	type args struct {
		in0   string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should not fail (can convert to int64)",
			args: args{
				in0:   "",
				value: "123",
			},
			wantErr: false,
		},
		{
			name: "should fail (cannot convert to int64)",
			args: args{
				in0:   "",
				value: "hello",
			},
			wantErr: true,
		},
		{
			name: "invalid type",
			args: args{
				in0:   "",
				value: []string{"123"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsInt64(tt.args.in0, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("IsInt64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsFloat64(t *testing.T) {
	type args struct {
		in0   string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should not fail (can convert to float64)",
			args: args{
				in0:   "",
				value: "123",
			},
			wantErr: false,
		},
		{
			name: "should fail (cannot convert to float64)",
			args: args{
				in0:   "",
				value: "hello",
			},
			wantErr: true,
		},
		{
			name: "invalid type",
			args: args{
				in0:   "",
				value: []string{"123"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsFloat64(tt.args.in0, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("IsFloat64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	type args struct {
		args  string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should fail (no matchers)",
			args: args{
				args:  "gibberish",
				value: "",
			},
			wantErr: true,
		},
		{
			name: "should fail (no matchers []byte)",
			args: args{
				args:  "gibberish",
				value: []byte(``),
			},
			wantErr: true,
		},
		{
			name: "should fail (no matchers []string)",
			args: args{
				args:  "gibberish",
				value: []string{},
			},
			wantErr: true,
		},
		{
			name: "should not fail (match email)",
			args: args{
				args:  "email",
				value: "someone@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "should not fail (match email []byte)",
			args: args{
				args:  "email",
				value: []byte(`someone@gmail.com`),
			},
			wantErr: false,
		},
		{
			name: "should fail (not email)",
			args: args{
				args:  "email",
				value: "",
			},
			wantErr: true,
		},
		{
			name: "should fail (not email []byte)",
			args: args{
				args:  "email",
				value: []byte(``),
			},
			wantErr: true,
		},
		{
			name: "should not fail (url)",
			args: args{
				args:  "url",
				value: "https://google.com",
			},
			wantErr: false,
		},
		{
			name: "should not fail ([]string email)",
			args: args{
				args:  "email",
				value: []string{"someone@gmail.com"},
			},
			wantErr: false,
		},
		{
			name: "should fail ([]string not email)",
			args: args{
				args:  "email",
				value: []string{""},
			},
			wantErr: true,
		},
		{
			name: "unexpected type",
			args: args{
				args:  "email",
				value: map[string]string{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Matches(tt.args.args, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Matches() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bounds(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantMin float64
		wantMax float64
		wantErr bool
	}{
		{
			name: "valid range",
			args: args{
				s: "0..10",
			},
			wantMin: 0.0,
			wantMax: 10.0,
			wantErr: false,
		},
		{
			name: "right range broken",
			args: args{
				s: "0..",
			},
			wantMin: 0,
			wantMax: 0,
			wantErr: true,
		},
		{
			name: "left range broken",
			args: args{
				s: "..10",
			},
			wantMin: 0,
			wantMax: 0,
			wantErr: true,
		},
		{
			name: "invalid range types",
			args: args{
				s: "hello..world",
			},
			wantMin: 0,
			wantMax: 0,
			wantErr: true,
		},
		{
			name: "left invalid",
			args: args{
				s: "hello..10",
			},
			wantMin: 0,
			wantMax: 0,
			wantErr: true,
		},
		{
			name: "right invalid",
			args: args{
				s: "0..hello",
			},
			wantMin: 0,
			wantMax: 0,
			wantErr: true,
		},
		{
			name: "left wildcard",
			args: args{
				s: "*..10",
			},
			wantMin: -math.MaxFloat64,
			wantMax: 10,
			wantErr: false,
		},
		{
			name: "right wilcard",
			args: args{
				s: "0..*",
			},
			wantMin: 0,
			wantMax: math.MaxFloat64,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax, err := bounds(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("bounds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMin != tt.wantMin {
				t.Errorf("bounds() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("bounds() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func Test_checkLength(t *testing.T) {
	type args struct {
		args  string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "broken range",
			args: args{
				args:  "0..",
				value: "hello",
			},
			wantErr: true,
		},
		{
			name: "should not fail",
			args: args{
				args:  "0..10",
				value: "hello",
			},
			wantErr: false,
		},
		{
			name: "should fail (too long)",
			args: args{
				args:  "0..10",
				value: "hello world",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkLength(tt.args.args, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("checkLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_strIn(t *testing.T) {
	type args struct {
		str    string
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				str:    "",
				values: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strIn(tt.args.str, tt.args.values); got != tt.want {
				t.Errorf("strIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_f64(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "float with precision",
			args: args{
				f: 64.12,
			},
			want: "64.12",
		},
		{
			name: "float without precision",
			args: args{
				f: 64.0,
			},
			want: "64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f64(tt.args.f); got != tt.want {
				t.Errorf("f64() = %v, want %v", got, tt.want)
			}
		})
	}
}
