package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsNullDive(t *testing.T) {
	r := require.New(t)

	field := []string{"", "", ""}

	v := StringSliceDive{
		Validator: &StringIsNull{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"", " 12.5", "   ", "a", " ", ""}

	v = StringSliceDive{
		Validator: &StringIsNull{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(4, e.Count())
}
