//go:build !solution

package testequal

import (
	"bytes"
	rfl "reflect"
)

func checkEqual(expected, actual interface{}) bool {
	if expected == nil && actual == nil {
		return true
	}

	if expected == nil || actual == nil {
		return false
	}

	expVal := rfl.ValueOf(expected)
	actVal := rfl.ValueOf(actual)

	if expVal.Kind() != actVal.Kind() {
		return false
	}

	switch expVal.Kind() {
	case rfl.Int, rfl.Int8, rfl.Int16, rfl.Int32, rfl.Int64:
		return expVal.Int() == actVal.Int()
	case rfl.Uint, rfl.Uint8, rfl.Uint16, rfl.Uint32, rfl.Uint64:
		return expVal.Uint() == actVal.Uint()
	case rfl.String:
		return expVal.String() == actVal.String()
	case rfl.Slice:
		{
			if (expVal.IsNil() != actVal.IsNil()) || (expVal.Len() != actVal.Len()) {
				return false
			}

			expElemType := expVal.Type().Elem()
			actElemType := actVal.Type().Elem()

			if expElemType.Kind() != actElemType.Kind() {
				return false
			}

			switch expElemType.Kind() {
			case rfl.Uint8:
				{
					return bytes.Equal(expVal.Bytes(), actVal.Bytes())
				}
			case rfl.Int:
				{
					for i := 0; i < expVal.Len(); i++ {
						if expVal.Index(i).Int() != actVal.Index(i).Int() {
							return false
						}
					}

					return true
				}
			default:
				return false
			}
		}

	case rfl.Map:
		{
			if (expVal.IsNil() != actVal.IsNil()) || (expVal.Len() != actVal.Len()) {
				return false
			}

			expKeyType := expVal.Type().Key().Kind()
			actKeyType := actVal.Type().Key().Kind()

			if (expKeyType != rfl.String) || (expKeyType != actKeyType) {
				return false
			}

			expElemType := expVal.Type().Elem().Kind()
			actElemType := actVal.Type().Elem().Kind()

			if (expElemType != rfl.String) || (expElemType != actElemType) {
				return false
			}

			for _, key := range expVal.MapKeys() {
				if expVal.MapIndex(key).String() != actVal.MapIndex(key).String() {
					return false
				}
			}

			return true
		}

	default:
		return false
	}
}

func log(t T, msgAndArgs ...interface{}) {
	t.Helper()

	if len(msgAndArgs) == 0 {
		t.Errorf("")
	} else if len(msgAndArgs) == 1 {
		if msgFormat, ok := msgAndArgs[0].(string); ok {
			t.Errorf(msgFormat)
		} else {
			t.Errorf("%v", msgAndArgs[0])
		}
	} else {
		if msgFormat, ok := msgAndArgs[0].(string); ok {
			t.Errorf(msgFormat, msgAndArgs[1:]...)
		} else {
			t.Errorf("%v", msgAndArgs[0])
		}
	}
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if checkEqual(expected, actual) {
		return true
	}

	log(t, msgAndArgs...)
	return false
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !checkEqual(expected, actual) {
		return true
	}

	log(t, msgAndArgs...)
	return false
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if checkEqual(expected, actual) {
		return
	}

	log(t, msgAndArgs...)
	t.FailNow()
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !checkEqual(expected, actual) {
		return
	}

	log(t, msgAndArgs...)
	t.FailNow()
}
