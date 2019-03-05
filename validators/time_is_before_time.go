package validators

import (
	"fmt"
	"time"

	"github.com/s3rj1k/validator"
)

// TimeIsBeforeTimeError is a function that defines error message returned by TimeIsBeforeTime validator.
// nolint: gochecknoglobals
var TimeIsBeforeTimeError = func(v *TimeIsBeforeTime) string {
	return fmt.Sprintf("%s must be before %s", v.Name, v.ComparedName)
}

// TimeIsBeforeTime is a validator object.
type TimeIsBeforeTime struct {
	Name          string
	Field         time.Time
	ComparedName  string
	ComparedField time.Time
}

// Validate adds an error if the Field time is not before the ComparedField time.
func (v *TimeIsBeforeTime) Validate(e *validator.Errors) {
	if v.Field.UnixNano() < v.ComparedField.UnixNano() {
		return
	}

	e.Add(v.Name, TimeIsBeforeTimeError(v))
}
