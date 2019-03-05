package validators

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringsArePathsNotInTheSameDirError is a function that defines error message returned by StringsArePathsNotInTheSameDir validator.
// nolint: gochecknoglobals
var StringsArePathsNotInTheSameDirError = func(v *StringsArePathsNotInTheSameDir) string {
	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' path is in the same dir with '%s'", v.Name, v.ComparedField)
	}

	return fmt.Sprintf("'%s' path is in the same dir with '%s'", v.Name, v.ComparedName)
}

// StringsArePathsNotInTheSameDir is a validator object
type StringsArePathsNotInTheSameDir struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField string
}

// Validate adds an error if paths share same path tree to last path element.
// Supplied paths are converted to absolute paths before comparison.
func (v *StringsArePathsNotInTheSameDir) Validate(e *validator.Errors) {

	absFieldPath, _ := filepath.Abs(v.Field)
	absComparedFieldPath, _ := filepath.Abs(v.ComparedField)

	if (filepath.Dir(absFieldPath) != filepath.Dir(absComparedFieldPath)) &&
		filepath.IsAbs(absFieldPath) &&
		filepath.IsAbs(absComparedFieldPath) {
		return
	}

	e.Add(v.Name, StringsArePathsNotInTheSameDirError(v))
}

// SetField sets validator field.
func (v *StringsArePathsNotInTheSameDir) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringsArePathsNotInTheSameDir) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
