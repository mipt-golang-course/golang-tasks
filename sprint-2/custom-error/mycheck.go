package mycheck

import (
	"fmt"
	"strings"
	"unicode"
)

// Errors - пользовательский тип для слайса ошибок
type Errors []error

// Error - метод для вывода всех ошибок слайса через "; "
func (error Errors) Error() string {
	msgs := make([]string, 0, len(error))
	for _, err := range error {
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

// func main() {
// 	for {
// 		fmt.Printf("Укажите строку (q для выхода): ")
// 		reader := bufio.NewReader(os.Stdin)
// 		ret, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		ret = strings.TrimRight(ret, "\n")
// 		if ret == `q` {
// 			break
// 		}
// 		if err = MyCheck(ret); err != nil {
// 			fmt.Println(err)
// 		} else {
// 			fmt.Println(`Строка прошла проверку`)
// 		}
// 	}
// }
