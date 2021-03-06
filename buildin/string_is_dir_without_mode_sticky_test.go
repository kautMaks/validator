package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsDirWithoutModeSticky(t *testing.T) {
	r := require.New(t)

	// setup
	_silentdeleteMany(notExists, notDir, notStickyDir, stickyDir)
	_createdirs(notStickyDir, stickyDir)
	_setfilemode(stickyDir, os.ModeSticky)
	_createfile(notDir)

	// teardown
	defer _silentdeleteMany(notStickyDir, stickyDir, notDir)

	var tests = []struct {
		field string
		valid bool
	}{
		{notStickyDir, true},

		{notExists, false},
		{notDir, false},
		{stickyDir, false},

		{"", false},
	}

	for index, test := range tests {
		v := &StringIsDirWithoutModeSticky{Name: "StickyDir", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsDirWithoutModeStickyError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
