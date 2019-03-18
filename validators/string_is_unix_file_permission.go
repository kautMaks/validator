package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsUnixFilePermissionError is a function that defines error message returned by StringIsUnixFilePermission validator.
// nolint: gochecknoglobals
var StringIsUnixFilePermissionError = func(v *StringIsUnixFilePermission) string {
	return fmt.Sprintf("'%s' is not a unix file permission", v.Field)
}

// StringIsUnixFilePermission is a validator object.
type StringIsUnixFilePermission struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a unix file permission
func (v *StringIsUnixFilePermission) Validate(e *validator.Errors) {

	if rxUnixFilePermission.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsUnixFilePermissionError(v))
}

// SetField sets validator field.
func (v *StringIsUnixFilePermission) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUnixFilePermission) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
