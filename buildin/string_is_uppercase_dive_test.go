package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsUpperCaseDive(t *testing.T) {
	r := require.New(t)

	field := []string{"NEED", "ONLY", "UPPER CASES", "  ", ""}

	v := StringSliceDive{
		Validator: &StringIsUpperCase{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"These", "HAVE", "some", "Lowers", "", " "}

	v = StringSliceDive{
		Validator: &StringIsUpperCase{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(3, e.Count())
}
