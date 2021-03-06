package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsSymlinkAndTargetIsNotDir(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(notSymlink) // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := &StringIsSymlinkAndTargetIsNotDir{Name: "Name", Field: notSymlink}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // not a symlink is OK

	_ = os.Remove(symlink)
	err = os.Symlink(notSymlink, symlink) // symlink to file
	r.Nil(err)

	v = &StringIsSymlinkAndTargetIsNotDir{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // symlink to file is OK

	_ = os.Remove(symlink)
	err = os.Symlink("/tmp", symlink) // symlink to folder
	r.Nil(err)

	v = &StringIsSymlinkAndTargetIsNotDir{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count()) // symlink to folder is error
	r.Equal([]string{StringIsSymlinkAndTargetIsNotDirError(v)}, e.Get("Name"))

	err = os.Remove(symlink)
	r.Nil(err)

	err = os.Remove(notSymlink)
	r.Nil(err)

	v = &StringIsSymlinkAndTargetIsNotDir{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // does not exist is OK
}
