package mycheck

import (
	"errors"
	"strings"
)

type ErrorList []error

func (el ErrorList) Error() string {
	var msgs []string
	for _, err := range el {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, ";")
}

func MyCheck(input string) error {
	var errList ErrorList
	hasNumbers := false
	for _, char := range input {
		if char >= '0' && char <= '9' {
			hasNumbers = true
			break
		}
	}
	if hasNumbers {
		errList = append(errList, errors.New("found numbers"))
	}
	if len(input) >= 20 {
		errList = append(errList, errors.New("line is too long"))
	}
	spaceCount := 0
	for _, char := range input {
		if char == ' ' {
			spaceCount++
		}
	}
	if spaceCount != 2 {
		errList = append(errList, errors.New("no two spaces"))
	}
	if len(errList) > 0 {
		return errList
	}
	return nil
}
