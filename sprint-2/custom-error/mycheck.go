package mycheck

import (
    "errors"
    "strings"
	"unicode"
)

// 0) Статическое указание ошибок на этапе инициализации
var (
	ErrNumbers = errors.New("found numbers")
	ErrLength = errors.New("line is too long")
	ErrSpaces = errors.New("no two spaces")
)

// 1) вставьте определение типа для []error
type MyErr []error

// 2) определите метод Error для вашего типа, который будет выводить все ошибки слайса

func (myerrors MyErr) Error() string{
	var result string

	for iter, elem := range myerrors{
		result = result + elem.Error()
		if iter < len(myerrors)-1 {
			result = result + ";"
		}
	}

	return result
}

// 3) реализуйте функцию MyCheck

func MyCheck(input string) error {

	errorsSlice := make(MyErr, 0)

	var numbers = 0

    for _, i := range input {
		if unicode.IsDigit(i) {
			numbers = 1
			break
		}
    }

	if numbers ==1 {
		errorsSlice = append(errorsSlice, ErrNumbers)
	}
	if len(input) >= 20 {
		errorsSlice = append(errorsSlice, ErrLength)
	}
	if strings.Count(input, " ") != 2 {
		errorsSlice = append(errorsSlice, ErrSpaces)
	}

	if len(errorsSlice) > 0 {
		return errorsSlice
	}

	return nil
}