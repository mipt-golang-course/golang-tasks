package mycheck_test

import (
	"testing"

	my_check "github.com/mipt-golang-course/golang-tasks/sprint-2/custom-error"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input       string
	resultError string
}

func TestMyCheck(t *testing.T) {
	for _, tc := range []testCase{
		{input: "too_long line abcdabcdabcdabcdabcd", resultError: "line is too long"},
		{input: "found numbers 123", resultError: "found numbers"},
		{input: "not_two space", resultError: "no two spaces"},
		{input: "all ok man", resultError: ""},
		{input: "line too long, found numbers 123 and not 2 spaces", resultError: "found numbers;line is too long;no two spaces"},
		{input: "line too long, not two spaces", resultError: "line is too long;no two spaces"},
	} {
		t.Run(tc.input, func(t *testing.T) {
			err := my_check.MyCheck(tc.input)

			if tc.resultError == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tc.resultError)
			}
		})
	}
}
