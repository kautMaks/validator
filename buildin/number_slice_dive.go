package buildin

import (
	"github.com/s3rj1k/validator"
)

// NumberValidator is an interface for string validator objects.
type NumberValidator interface {
	Validate(*validator.Errors)
	SetField(interface{})
	SetNameIndex(int)
	GetName() string
}

// NumberSliceDive is a validator object.
// Validate applies Validator to each value in the Field.
type NumberSliceDive struct {
	Validator NumberValidator
	Field     interface{}
}

// Validate applies Validator to each value in the Field.
func (v *NumberSliceDive) Validate(e *validator.Errors) {
	slice := v.Field

	if slice == nil {
		slice = []int8{0}
	}

	switch field := slice.(type) {
	case []int8:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []int16:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []int32:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []int:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []int64:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []uintptr:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []uint8:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []uint16:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []uint32:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []uint:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	case []uint64:
		for i, val := range field {
			v.Validator.SetField(val)
			v.Validator.SetNameIndex(i)
			v.Validator.Validate(e)
		}
	default:
		e.Add(v.Validator.GetName(), ErrBadNumType.Error())

		return
	}
}
