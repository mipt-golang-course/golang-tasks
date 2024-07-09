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

var bigNumbers = [6]string{"billion", "million", "thousand", "hundred"}

var numbers = [4]int64{1000000000, 1000000, 1000, 100}

func Units(n int64) string {
	var result string
	var hundreds int64
	if n > 99 {
		hundreds = n / 100
		result = nameNumbers[hundreds] + " hundred "
		n = n % 100
	}
	switch {
	case n > 0 && n < 20:
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
	var sing bool
	if n < 0 {
		sing = true
		n = n * -1
	}
	if n == 0 {
		result = nameNumbers[0]
	} else if n < 1000 {
		result = Units(n)
	} else {
		var i int64
		for i = 0; i < 4; i++ {
			if n > 999 {
				iterN = n / numbers[i]
			} else {
				iterN = n
			}
			if iterN == 0 {
				continue
			}
			result += strings.TrimRight(Units(iterN), " ") + " "
			if i != 3 {
				result += bigNumbers[i] + " "
			}
			n = n % numbers[i]
		}
	}
	if sing {
		result = "minus " + result
	}
	return strings.TrimRight(result, " ")
}
