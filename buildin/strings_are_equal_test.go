package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringsAreEqual(t *testing.T) {
	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "test", true},
		{"test_fail", "test_true", false},
		{"test with space", " test with space ", false},
		{" test with space second", " test with space second       ", false},
	}

	for _, testCase := range cases {
		v := &StringsAreEqual{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	for _, testCase := range cases {
		v := &StringsAreEqual{Name: "strings1", Field: testCase.str1, ComparedField: testCase.str2, ComparedName: "strings2"}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)

		if !testCase.expected {
			r.Contains(e.Get("strings1"), StringsAreEqualError(v))
		}
	}

	v := &StringsAreEqual{Name: "strings", Field: "test_fail", ComparedField: "test"}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringsAreEqualError(v)}, e.Get("strings"))
}

func Test_StringsAreIEqual(t *testing.T) {
	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "TesT", true},
		{"test_fail", "Test_truE", false},
		{"test with space", " Test with spacE ", false},
		{" test with space second", " Test with space seconD       ", false},
	}

	for _, testCase := range cases {
		v := &StringsAreEqual{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2, CaseInsensitive: true}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	for _, testCase := range cases {
		v := &StringsAreEqual{Name: "strings1", Field: testCase.str1, ComparedField: testCase.str2, ComparedName: "strings2", CaseInsensitive: true}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)

		if !testCase.expected {
			r.Contains(e.Get("strings1"), StringsAreEqualError(v))
		}
	}

	v := &StringsAreEqual{Name: "strings", Field: "test_fail", ComparedField: "test", CaseInsensitive: true}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringsAreEqualError(v)}, e.Get("strings"))
}
