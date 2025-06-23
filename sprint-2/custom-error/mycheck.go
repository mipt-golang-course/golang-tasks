//go:build !solution

package mycheck

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type MyError []error

func (err *MyError) AddError(newErr string) {
	*err = append(*err, errors.New(newErr))
}

func (err MyError) Error() error {
	switch len(err) {
	case 0:
		return nil
	case 1:
		return err[0]
	default:
		return fmt.Errorf("%w;%w", err[0], err[1:].Error())
	}
}

func MyCheck(input string) error {
	var errs MyError

	var spaceCounter int
	var foundNumbers, isLong bool

	if len(input) >= 20 {
		isLong = true
	}

	for _, val := range input {
		if unicode.IsDigit(val) {
			foundNumbers = true
		} else if val == ' ' {
			spaceCounter++
		}
	}

	if foundNumbers {
		errs.AddError("found numbers")
	}
	if isLong {
		errs.AddError("line is too long")
	}
	if spaceCounter != 2 {
		errs.AddError("no two spaces")
	}

	return errs.Error()
}

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
