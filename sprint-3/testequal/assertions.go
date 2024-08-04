//go:build !solution

package testequal

import (
	"bytes"
	"maps"
	"slices"
)

func actualCheck(expected, actual interface{}) bool {
	switch exp := expected.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return expected == actual

	case string:
		if act, ok := actual.(string); ok && exp == act {
			return true
		}

	case map[string]string:
		if act, ok := actual.(map[string]string); ok {
			if exp == nil && act == nil {
				return true
			}

			if exp != nil && act != nil && maps.Equal(exp, act) {
				return true
			}

			return false
		}

	case []int:
		if act, ok := actual.([]int); ok {
			if exp == nil && act == nil {
				return true
			}

			if exp != nil && act != nil && slices.Equal(exp, act) {
				return true
			}

			return false
		}

	case []byte:
		if act, ok := actual.([]byte); ok {
			if exp == nil && act == nil {
				return true
			}

			if exp != nil && act != nil && bytes.Equal(exp, act) {
				return true
			}

			return false
		}
	}

	return false
}

func printMsgAndArgs(t T, msgAndArgs ...interface{}) {
	t.Helper()

	if len(msgAndArgs) == 0 {
		t.Errorf("")
	} else if len(msgAndArgs) == 1 {
		t.Errorf(msgAndArgs[0].(string))
	} else {
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if actualCheck(expected, actual) {
		return true
	}

	t.Helper()
	printMsgAndArgs(t, msgAndArgs...)

	return false
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if !actualCheck(expected, actual) {
		return true
	}

	t.Helper()
	printMsgAndArgs(t, msgAndArgs...)

	return false
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if actualCheck(expected, actual) {
		return
	}

	t.Helper()
	printMsgAndArgs(t, msgAndArgs...)

	t.FailNow()
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if !actualCheck(expected, actual) {
		return
	}

	t.Helper()
	printMsgAndArgs(t, msgAndArgs...)

	t.FailNow()
}
