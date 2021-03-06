package buildin

import (
	"fmt"
	"unicode/utf8"

	"github.com/s3rj1k/validator"
)

// StringLengthInRangeError is a function that defines error message returned by StringLengthInRange validator.
// nolint: gochecknoglobals
var StringLengthInRangeError = func(v *StringLengthInRange) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	min := v.Min
	max := v.Max

	strLength := utf8.RuneCountInString(v.Field)
	if max == 0 {
		max = strLength
	}

	return fmt.Sprintf("'%s' not in range(%d, %d)", v.Field, min, max)
}

// StringLengthInRange is a validator object.
// Validate adds an error if the Field length is not in range between Min and Max (inclusive).
// If only Min provided - Max=length of string. If only Max provided - Min=0.
// It is possible to provide either both or one of the Min/Max values.
type StringLengthInRange struct {
	Name    string
	Field   string
	Min     int
	Max     int
	Message string
}

// Validate adds an error if the Field length is not in range between Min and Max (inclusive).
// If only Min provided - Max=length of string. If only Max provided - Min=0.
// It is possible to provide either both or one of the Min/Max values.
func (v *StringLengthInRange) Validate(e *validator.Errors) {
	strLength := utf8.RuneCountInString(v.Field)

	min := v.Min
	max := v.Max

	if max == 0 {
		max = strLength
	}

	if strLength >= min && strLength <= max {
		return
	}

	e.Add(v.Name, StringLengthInRangeError(v))
}

// SetField sets validator field.
func (v *StringLengthInRange) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringLengthInRange) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
