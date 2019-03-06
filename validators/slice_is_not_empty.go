package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// SliceIsNotEmptyError is a function that defines error message returned by SliceIsNotEmpty validator.
// nolint: gochecknoglobals
var SliceIsNotEmptyError = func(v *SliceIsNotEmpty) string {
	return fmt.Sprintf("%s slice is empty", v.Name)
}

// SliceIsNotEmpty is a validator object.
type SliceIsNotEmpty struct {
	Name  string
	Field interface{}
}

// Validate adds an error if the slice in Field is empty.
// Standard check adds an error if the slice is nil or of unsupported type.
// Supports all Go basic types except for bool https://tour.golang.org/basics/11
// nolint: gocyclo
func (v *SliceIsNotEmpty) Validate(e *validator.Errors) {

	if v.Field == nil {
		e.Add(v.Name, ErrNilValue.Error())

		return
	}

	switch field := v.Field.(type) {
	case []string:
		if len(field) > 0 {
			return
		}
	case []int8:
		if len(field) > 0 {
			return
		}
	case []int16:
		if len(field) > 0 {
			return
		}
	case []int32:
		if len(field) > 0 {
			return
		}
	case []int:
		if len(field) > 0 {
			return
		}
	case []int64:
		if len(field) > 0 {
			return
		}
	case []uintptr:
		if len(field) > 0 {
			return
		}
	case []uint8:
		if len(field) > 0 {
			return
		}
	case []uint16:
		if len(field) > 0 {
			return
		}
	case []uint32:
		if len(field) > 0 {
			return
		}
	case []uint:
		if len(field) > 0 {
			return
		}
	case []uint64:
		if len(field) > 0 {
			return
		}
	case []float32:
		if len(field) > 0 {
			return
		}
	case []float64:
		if len(field) > 0 {
			return
		}
	case []complex64:
		if len(field) > 0 {
			return
		}
	case []complex128:
		if len(field) > 0 {
			return
		}
	default:
		e.Add(v.Name, ErrBadSliceType.Error())
		return
	}

	e.Add(v.Name, SliceIsNotEmptyError(v))
}
