package validation

import (
	"fmt"
	"net/http"
	"time"
)

// ValidationError holds the error message and corresponding HTTP status code
type ValidationError struct {
	Field   string
	Message string
	Code    int
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}

// Rule is a function type that returns an error if validation fails
type Rule func() error

type Validator struct {
	rules []Rule
}

func (v *Validator) Add(rule Rule) {
	v.rules = append(v.rules, rule)
}

func (v *Validator) Execute() []error {
	var errors []error
	for _, rule := range v.rules {
		if err := rule(); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func ValidateRequired(fieldName, value string) Rule {
	return func() error {
		if value == "" {
			return ValidationError{
				Field:   fieldName,
				Message: "is required",
				Code:    http.StatusBadRequest,
			}
		}
		return nil
	}
}

func ValidateDate(fieldName string, value time.Time) Rule {
	return func() error {
		if value.IsZero() {
			return ValidationError{
				Field:   fieldName,
				Message: "is required and must be a valid date",
				Code:    http.StatusBadRequest,
			}
		}
		return nil
	}
}
