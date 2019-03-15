package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsValidUserOrGroupNameError is a function that defines error message returned by StringIsValidUserOrGroupName validator.
// nolint: gochecknoglobals
var StringIsValidUserOrGroupNameError = func(v *StringIsValidUserOrGroupName) string {
	return fmt.Sprintf("'%s' is not a valid user or group name", v.Name)
}

// StringIsValidUserOrGroupName is a validator object.
type StringIsValidUserOrGroupName struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a valid user or group name.
func (v *StringIsValidUserOrGroupName) Validate(e *validator.Errors) {

	if isValidUserOrGroupName(v.Field) {
		return
	}

	e.Add(v.Name, StringIsValidUserOrGroupNameError(v))
}

// SetField sets validator field.
func (v *StringIsValidUserOrGroupName) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsValidUserOrGroupName) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}

func isValidUserOrGroupName(name string) bool {

	if len(name) < 1 || len(name) > 32 {
		return false
	}

	if !rxUserGroupName.MatchString(name) {
		return false
	}

	return true
}