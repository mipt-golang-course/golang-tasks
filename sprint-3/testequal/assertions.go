package testequal

import (
	"bytes"
	"maps"
	"slices"
)

func checkEquality(expected, actual interface{}) bool {
	switch exp := expected.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return exp == actual

	case string:
		return checkStringEquality(exp, actual)

	case map[string]string:
		return checkMapStringEquality(exp, actual)

	case []int:
		return checkSliceIntEquality(exp, actual)

	case []byte:
		return checkSliceByteEquality(exp, actual)

	default:
		return false
	}
}

func checkStringEquality(expected string, actual interface{}) bool {
	if act, ok := actual.(string); ok && expected == act {
		return true
	}
	return false
}

func checkMapStringEquality(expected map[string]string, actual interface{}) bool {
	if act, ok := actual.(map[string]string); ok {
		return maps.Equal(expected, act)
	}
	return false
}

func checkSliceIntEquality(expected []int, actual interface{}) bool {
	if act, ok := actual.([]int); ok {
		return slices.Equal(expected, act)
	}
	return false
}

func checkSliceByteEquality(expected []byte, actual interface{}) bool {
	if act, ok := actual.([]byte); ok {
		return bytes.Equal(expected, act)
	}
	return false
}

func displayMessage(t T, msgAndArgs ...interface{}) {
	t.Helper()

	switch len(msgAndArgs) {
	case 0:
		t.Errorf("")
	case 1:
		t.Errorf(msgAndArgs[0].(string))
	default:
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if checkEquality(expected, actual) {
		return true
	}

	t.Helper()
	displayMessage(t, msgAndArgs...)

	return false
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if !checkEquality(expected, actual) {
		return true
	}

	t.Helper()
	displayMessage(t, msgAndArgs...)

	return false
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if checkEquality(expected, actual) {
		return
	}

	t.Helper()
	displayMessage(t, msgAndArgs...)

	t.FailNow()
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if !checkEquality(expected, actual) {
		return
	}

	t.Helper()
	displayMessage(t, msgAndArgs...)

	t.FailNow()
}
