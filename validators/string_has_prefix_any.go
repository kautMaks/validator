package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasPrefixAnyError is a function that defines error message returned by StringHasPrefixAny validator.
// nolint: gochecknoglobals
var StringHasPrefixAnyError = func(v *StringHasPrefixAny) string {
	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' has no prefix from '%v'", v.Name, v.ComparedField)
	}

	return fmt.Sprintf("'%s' has no prefix from '%s'", v.Name, v.ComparedName)
}

// StringHasPrefixAny is a validator object.
type StringHasPrefixAny struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField []string
}

// Validate adds an error if the Field is not prefixed by at least one string from ComparedField.
func (v *StringHasPrefixAny) Validate(e *validator.Errors) {

	// if no required prefixes - string is valid
	if v.ComparedField == nil || len(v.ComparedField) == 0 {
		return
	}

	for _, s := range v.ComparedField {
		if strings.HasPrefix(v.Field, s) {
			return
		}
	}

	e.Add(v.Name, StringHasPrefixAnyError(v))
}

// SetField sets validator field.
func (v *StringHasPrefixAny) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasPrefixAny) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}