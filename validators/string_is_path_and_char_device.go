package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndCharDeviceError is a function that defines error message returned by StringIsPathAndCharDevice validator.
// nolint: gochecknoglobals
var StringIsPathAndCharDeviceError = func(v *StringIsPathAndCharDevice) string {
	return fmt.Sprintf("'%s' is not an existing path or is an existing path without CharDevice mode", v.Field)
}

// StringIsPathAndCharDevice is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path without CharDevice mode.
type StringIsPathAndCharDevice struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not an existing path or is an existing path without CharDevice mode.
func (v *StringIsPathAndCharDevice) Validate(e *validator.Errors) {

	if Exists(v.Field) && isFileWithMode(v.Field, os.ModeCharDevice) {
		return
	}

	e.Add(v.Name, StringIsPathAndCharDeviceError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndCharDevice) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndCharDevice) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}