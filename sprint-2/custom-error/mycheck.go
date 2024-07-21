//go:build !solution

package mycheck

import (
	"strings"
)

type MyError struct {
	mes []string
}

func (errors MyError) Error() string {
	if len(errors.mes) == 0 {
		return ""
	}

	accum := errors.mes[0]

	for i := range errors.mes[1:] {
		accum += ";" + errors.mes[i+1]
	}

	return accum
}

func MyCheck(input string) error {
	err := MyError{}

	if strings.ContainsAny(input, "0123456789") {
		err.mes = append(err.mes, "found numbers")
	}

	if len(input) >= 20 {
		err.mes = append(err.mes, "line is too long")
	}

	if strings.Count(input, " ") != 2 {
		err.mes = append(err.mes, "no two spaces")
	}

	if len(err.mes) == 0 {
		return nil
	}

	return err
}
