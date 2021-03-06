package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasLowerCaseDive(t *testing.T) {
	r := require.New(t)

	field := []string{"Each", "Has", "Lowercase", ""} // 0 errors

	v := StringSliceDive{
		Validator: &StringHasLowerCase{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"TWO", "UPPERCASED", "but", "ONLy", "tWO", ""} // 2 errors

	v = StringSliceDive{
		Validator: &StringHasLowerCase{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(2, e.Count())
}
