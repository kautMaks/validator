package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsHTTPURL(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		url   string
		valid bool
	}{
		{"", false},
		{"http://", false},
		{"https://", false},
		{"http", false},
		{"google.com", false},
		{"http://www.google.com", true},
		{"http://google.com", true},
		{"google.com", false},
		{"https://www.google.cOM", true},
		{"ht123tps://www.google.cOM", false},
		{"https://golang.Org", true},
		{"https://invalid#$%#$@.Org", false},
	}

	for _, test := range tests {
		v := &StringIsHTTPURL{Name: "URL", Field: test.url}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(test.valid, !e.HasAny(), test.url, e.Error())
	}

	v := &StringIsHTTPURL{Name: "URL", Field: "http://"}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsHTTPURLError(v)}, e.Get("URL"))
}
