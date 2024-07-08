//go:build !solution

package speller

import (
	"strings"
)

var nameNumbers = [20]string{
	"zero", "one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen",
	"fifteen", "sixteen", "seventeen", "eighteen",
	"nineteen"}

var tyNumbers = map[int64]string{
	2: "twenty", 3: "thirty",
	4: "forty", 5: "fifty", 6: "sixty",
	7: "seventy", 8: "eighty",
	9: "ninety"}

var bigNumbers = [5]string{
	"zero", "hundred",
	"thousand", "million", "billion"}

var numbers = [5]int64{0, 100, 1000, 1000000, 1000000000}

func Floordiv(a, b int64) int64 {
	var x int64
	if a < 0 {
		x = b - 1
	}
	return (a - x) / b
}

func DigitOfNumber(n int64) int64 {
	var cnt int64 = 1
	n = Floordiv(n, 10)

	for n > 0 {
		n = Floordiv(n, 10)
		cnt++
	}
	return cnt
}

func Units(n int64) string {
	var result string
	var hundreds int64
	if n > 99 {
		hundreds = n / 100
		result = nameNumbers[hundreds] + " hundred "
		n = n % 100
	}
	switch {
	case hundreds == 0 && n < 20:
		result = result + nameNumbers[n]
	case n%10 == 0:
		result = result + tyNumbers[n/10]
	default:
		result = result + tyNumbers[n/10] + "-" + nameNumbers[n%10]
	}
	return result
}

func Spell(n int64) string {
	var result string
	var iterN int64
	if n == 0 {
		result = nameNumbers[0]
	} else if n < 1000 {
		result = Units(n)
	} else {
		var i int64
		class := DigitOfNumber(n) / 3
		if class < 2 {
			class = 2
		}
		iterClass := class
		for i = 1; i <= class; i++ {
			if n > 999 {
				iterN = n / numbers[iterClass]
			} else {
				iterN = n
			}
			if iterN == 0 {
				break
			}
			result += strings.TrimRight(Units(iterN), " ") + " " + bigNumbers[iterClass] + " "
			n = n % numbers[iterClass]
			iterClass--
		}
	}
	return strings.TrimRight(result, " ")
}
