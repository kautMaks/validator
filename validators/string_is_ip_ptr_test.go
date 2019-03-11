package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIPptr(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"8.8.8.8", true},
		{"8.8.4.4", true},
		{"2001:4860:4860::8888", true},

		{"5.255.253.0", false},
		{"220.181.0.0", false},
		{" 5.255.253.0", false},
		{"http://www.google.com", false},
		{"220.181.0.0/33", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsIPptr{Name: "IP", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {
			r.Equalf([]string{StringIsIPptrError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
