//go:build !solution

package speller

import "strings"

var (
	ones = []string{
		"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
		"seventeen", "eighteen", "nineteen",
	}
	tens = []string{
		"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}
	scales = []string{
		"", "thousand", "million", "billion", "trillion",
	}
)

func Spell(n int64) string {
	if n == 0 {
		return "zero"
	}

	if n < 0 {
		return "minus " + Spell(-n)
	}

	parts := []string{}
	scale := 0

	for n > 0 {
		if n%1000 != 0 {
			parts = append([]string{convertHundreds(n%1000) + " " + scales[scale]}, parts...)
		}
		n /= 1000
		scale++
	}

	return strings.TrimSpace(strings.Join(parts, " "))
}

func convertHundreds(n int64) string {
	if n == 0 {
		return ""
	} else if n < 20 {
		return ones[n]
	} else if n < 100 {
		if n%10 == 0 {
			return tens[n/10]
		}
		return tens[n/10] + "-" + ones[n%10]
	} else {
		if n%100 == 0 {
			return ones[n/100] + " hundred"
		}
		return ones[n/100] + " hundred " + convertTens(n%100)
	}
}

func convertTens(n int64) string {
	if n < 20 {
		return ones[n]
	}
	if n%10 == 0 {
		return tens[n/10]
	}
	return tens[n/10] + "-" + ones[n%10]
}
