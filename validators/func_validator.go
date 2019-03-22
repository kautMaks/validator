package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// FuncValidatorError is a function that defines error message returned by FuncValidator validator.
// nolint: gochecknoglobals
var FuncValidatorError = func(v *FuncValidator) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' failed custom func validation", v.Field)
}

// FuncValidator is a validator object.
type FuncValidator struct {
	Fn      func() bool
	Name    string
	Field   string
	Message string
}

// Validate is a validation method wrapper.
func (v *FuncValidator) Validate(e *validator.Errors) {
	// for backwards compatibility
	if strings.TrimSpace(v.Name) == "" {
		v.Name = v.Field
	}

	if v.Fn() {
		return
	}

	e.Add(v.Name, FuncValidatorError(v))
}
