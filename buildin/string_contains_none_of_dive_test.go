package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringContainsNoneOfDive(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field          []string
		comparedField  []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"foo", "bar", "golang "}, []string{"no", "chance"}, true, nil},
		{[]string{" bar ", "bar", ""}, []string{"  "}, true, nil},

		{[]string{"foo", "bar", "golang ", ""}, []string{"a"}, false, []int{1, 2}},
		{[]string{"foo", " bar  ", "bar", ""}, []string{""}, false, []int{0, 1, 2, 3}},
		{[]string{"foo", "bar", " golang ", " "}, []string{" "}, false, []int{2, 3}},
	}

	for index, test := range tests {
		v := StringSliceDive{
			Validator: &StringContainsNoneOf{Name: "StringContains", ComparedField: test.comparedField},
			Field:     test.field,
		}

		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("StringContains[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
