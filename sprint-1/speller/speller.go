package speller

import (
	"strings"
)

var (
	ones = map[int64]string{
		1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
		10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen",
		17: "seventeen", 18: "eighteen", 19: "nineteen",
	}
	tens = map[int64]string{
		2: "twenty", 3: "thirty", 4: "forty", 5: "fifty", 6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety",
	}
	scales = map[int]string{
		1: "thousand", 2: "million", 3: "billion", 4: "trillion",
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
	switch {
	case n == 0:
		return ""
	case n < 20:
		return ones[n]
	case n < 100:
		return convertTens(n)
	default:
		remainder := n % 100
		if remainder == 0 {
			return ones[n/100] + " hundred"
		}
		return ones[n/100] + " hundred " + convertTens(remainder)
	}
}

func convertTens(n int64) string {
	switch {
	case n < 20:
		return ones[n]
	default:
		remainder := n % 10
		if remainder == 0 {
			return tens[n/10]
		}
		return tens[n/10] + "-" + ones[remainder]
	}
}
