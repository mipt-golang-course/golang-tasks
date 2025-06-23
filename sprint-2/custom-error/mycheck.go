//go:build !solution

package mycheck

import (
	"errors"
	"unicode"

	"github.com/mipt-golang-course/golang-tasks/sprint-1/varjoin"
)

var (
	errFoundNumbers  error = errors.New("found numbers")
	errLineIsTooLong error = errors.New("line is too long")
	errNoTwoSpaces   error = errors.New("no two spaces")
)

const (
	minLengthLine  = 20
	requiredSpaces = 2
)

type myStringsErr []error

func (errStr myStringsErr) Error() string {
	messages := make([]string, 0, len(errStr))
	for _, err := range errStr {
		messages = append(messages, err.Error())
	}

	return varjoin.Join(";", messages...)
}

func MyCheck(input string) error {
	var (
		err          myStringsErr
		actualLength int64
		actualSpaces int64
		hasNumbers   bool
	)

	for _, symb := range input {
		actualLength++
		if unicode.IsDigit(symb) {
			hasNumbers = true
		}
		if unicode.IsSpace(symb) {
			actualSpaces++
		}
	}

	if hasNumbers {
		err = append(err, errFoundNumbers)
	}
	if actualLength > minLengthLine {
		err = append(err, errLineIsTooLong)
	}
	if actualSpaces != requiredSpaces {
		err = append(err, errNoTwoSpaces)
	}

	if len(err) == 0 {
		return nil
	}
	return err
}
