package v

import (
	"fmt"
	"strings"
)

// ErrorRequired is the type of error thrown when a required field is missing
type ErrorRequired struct {
	Field    string
	JSONName string
}

// Error satisfies the builtin Error interface
func (e ErrorRequired) Error() string {
	if e.JSONName != "" {
		return fmt.Sprintf("[validation] %s: required, please provide a value", e.JSONName)
	}
	return fmt.Sprintf("[validation] %s: required, please provide a value", e.Field)
}

// ErrorValidation is the type of error thrown on a validation error
type ErrorValidation struct {
	Name     string
	JSONName string
	Err      error
}

// Error satisfies the builtin Error interface
func (e ErrorValidation) Error() string {
	if e.JSONName != "" {
		return fmt.Sprintf("[validation] %s: %v", e.JSONName, e.Err)
	}
	return fmt.Sprintf("[validation] %s: %v", e.Name, e.Err)
}

type validationErorrs []error

// Error satisfies the builtin Error interface
func (v validationErorrs) Error() error {
	if len(v) == 0 {
		return nil
	}
	var messages []string
	for _, err := range v {
		messages = append(messages, err.Error())
	}
	return fmt.Errorf("%s", strings.Join(messages, " | "))
}
