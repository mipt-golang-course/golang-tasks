//go:build !solution

package speller

import (
	"strings"
)

var underTwenty = []string{
	"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
	"seventeen", "eighteen", "nineteen",
}

var tens = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

var thousands = []string{
	"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion",
}

func Spell(n int64) string {
	if n == 0 {
		return "zero"
	}

	if n < 0 {
		return "minus " + Spell(-n)
	}

	parts := []string{}
	thousandCounter := 0

	for n > 0 {
		if n%1000 != 0 {
			part := hundredsToWords(n % 1000)
			if thousands[thousandCounter] != "" {
				part += " " + thousands[thousandCounter]
			}
			parts = append([]string{part}, parts...)
		}
		n /= 1000
		thousandCounter++
	}

	return strings.TrimSpace(strings.Join(parts, " "))
}

func hundredsToWords(n int64) string {
	if n == 0 {
		return ""
	}

	hundredPart := ""
	if n >= 100 {
		hundredPart = underTwenty[n/100] + " hundred"
		n %= 100
	}

	if n == 0 {
		return hundredPart
	}

	if n < 20 {
		if hundredPart != "" {
			return hundredPart + " " + underTwenty[n]
		}
		return underTwenty[n]
	}

	tensPart := tens[n/10]
	onesPart := underTwenty[n%10]

	if hundredPart != "" {
		if onesPart != "" {
			return hundredPart + " " + tensPart + "-" + onesPart
		}
		return hundredPart + " " + tensPart
	}

	if onesPart != "" {
		return tensPart + "-" + onesPart
	}
	return tensPart
}
