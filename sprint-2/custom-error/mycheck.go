package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Errors - пользовательский тип для слайса ошибок
type Errors []error

// Error - метод для вывода всех ошибок слайса через "; "
func (e Errors) Error() string {
	var msgs []string
	for _, err := range e {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, ";")
}

// MyCheck - функция для проверки строки на соответствие условиям
func MyCheck(input string) error {
	var errs Errors

	// Проверка на наличие цифр
	for _, r := range input {
		if unicode.IsDigit(r) {
			errs = append(errs, fmt.Errorf("found numbers"))
			break
		}
	}

	// Проверка на длину строки
	if len(input) >= 20 {
		errs = append(errs, fmt.Errorf("line is too long"))
	}

	// Проверка на наличие двух пробелов
	if strings.Count(input, " ") != 2 {
		errs = append(errs, fmt.Errorf("no two spaces"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
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

// Что за ошибки выводид линтер go и как их можно испраить?
//
// sprint-2/custom-error/mycheck.go:16:2: Consider pre-allocating `msgs` (prealloc)
// var msgs []string
// ^
// sprint-2/custom-error/mycheck.go:65:26: SA4023: this comparison is always true (staticcheck)
// if err = MyCheck(ret); err != nil {
//    ^
// sprint-2/custom-error/mycheck.go:24:6: SA4023(related information): github.com/mipt-golang-course/golang-tasks/sprint-2/custom-error.MyCheck never returns a nil interface value (staticcheck)
// func MyCheck(input string) error {
