//go:build !solution

package mycheck

import (
	"errors"
	"fmt"
	"strings"
)

type errsSlice []error

const maxStringLen = 20

func (errs errsSlice) Error() string {
	if len(errs) == 0 {
		return ""
	}

	var errMsg strings.Builder

	errMsg.WriteString(fmt.Sprintf("%v", errs[0]))

	for i := range errs[1:] {
		errMsg.WriteString(";")
		errMsg.WriteString(fmt.Sprintf("%v", errs[i+1]))
	}

	return errMsg.String()
}

func MyCheck(input string) error {
	var errs errsSlice

	if strings.ContainsAny(input, "0123456789") {
		errs = append(errs, errors.New("found numbers"))
	}

	if len(input) >= maxStringLen {
		errs = append(errs, errors.New("line is too long"))
	}

	if strings.Count(input, " ") != 2 {
		errs = append(errs, errors.New("no two spaces"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
