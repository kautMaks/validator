package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasPrefixNoneOfError is a function that defines error message returned by StringHasPrefixNoneOf validator.
// nolint: gochecknoglobals
var StringHasPrefixNoneOfError = func(v *StringHasPrefixNoneOf) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' has prefix from %v", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' has prefix from contents of '%s'", v.Name, v.ComparedName)
}

// StringHasPrefixNoneOf is a validator object.
// Validate adds an error if the Field is prefixed by at least one string from ComparedField.
type StringHasPrefixNoneOf struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField []string
	Message       string
}

// Validate adds an error if the Field is prefixed by at least one string from ComparedField.
func (v *StringHasPrefixNoneOf) Validate(e *validator.Errors) {
	// if no excluding prefixes - string is valid
	if v.ComparedField == nil || len(v.ComparedField) == 0 {
		return
	}

	var hasPrefix = false

	for _, s := range v.ComparedField {
		if strings.HasPrefix(v.Field, s) {
			hasPrefix = true
		}
	}

	if !hasPrefix {
		return
	}

	e.Add(v.Name, StringHasPrefixNoneOfError(v))
}

// SetField sets validator field.
func (v *StringHasPrefixNoneOf) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasPrefixNoneOf) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
