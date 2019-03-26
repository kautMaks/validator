package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathError is a function that defines error message returned by StringIsPath validator.
// nolint: gochecknoglobals
var StringIsPathError = func(v *StringIsPath) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("path '%s' must exist", v.Field)
}

// StringIsPath is a validator object.
// Validate adds an error if the Field is a path that does not exist.
type StringIsPath struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a path that does not exist.
func (v *StringIsPath) Validate(e *validator.Errors) {
	if _, err := os.Stat(v.Field); !os.IsNotExist(err) {
		return
	}

	e.Add(v.Name, StringIsPathError(v))
}

// SetField sets validator field.
func (v *StringIsPath) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPath) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
