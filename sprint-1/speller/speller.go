//go:build !solution

package speller

import "strings"

var smallNumbers = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

var tens = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

var thousands = []string{
	"", "thousand", "million", "billion",
}

func intPartToWord(n int64) string {
	if n == 0 {
		return ""
	}

	if n < 20 {
		return smallNumbers[n]
	}

	if n < 100 {
		if n%10 == 0 {
			return tens[n/10]
		}
		return tens[n/10] + "-" + smallNumbers[n%10]
	}

	if n < 1000 {
		if n%100 == 0 {
			return smallNumbers[n/100] + " hundred"
		}
		return smallNumbers[n/100] + " hundred " + intPartToWord(n%100)
	}

	return ""
}

func Spell(n int64) string {
	if n == 0 {
		return smallNumbers[0]
	}

	if n < 0 {
		return "minus " + Spell(-n)
	}

	parts := []string{}
	for i := 0; n > 0; i++ {
		if n%1000 != 0 {
			parts = append([]string{intPartToWord(n%1000) + " " + thousands[i]}, parts...)
		}
		n /= 1000
	}

	return strings.TrimSpace(strings.Join(parts, " "))
}
