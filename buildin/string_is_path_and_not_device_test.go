package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsPathAndNotDevice(t *testing.T) {
	r := require.New(t)

	// setup
	_silentdeleteMany(notExists, namedPipe)
	_createnamedpipe(namedPipe)

	// teardown
	defer _silentdeleteMany(namedPipe)

	var tests = []struct {
		field string
		valid bool
	}{
		{"/dev/tty", false},

		{notExists, false},
		{namedPipe, true},
		{"/dev/null", false},

		{"", false},
	}

	for index, test := range tests {
		v := &StringIsPathAndNotDevice{Name: "FileModes", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsPathAndNotDeviceError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
