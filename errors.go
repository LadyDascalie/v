package v

import "fmt"

type ErrorRequired struct {
	Field    string
	JSONName string
}

func (e ErrorRequired) Error() string {
	if e.JSONName != "" {
		return fmt.Sprintf("[validation] %s: required, please provide a value", e.JSONName)
	}
	return fmt.Sprintf("[validation] %s: required, please provide a value", e.Field)
}

type ErrorValidation struct {
	Name     string
	JSONName string
	Err      error
}

func (e ErrorValidation) Error() string {
	if e.JSONName != "" {
		return fmt.Sprintf("[validation] %s: %v", e.JSONName, e.Err)
	}
	return fmt.Sprintf("[validation] %s: %v", e.Name, e.Err)
}
