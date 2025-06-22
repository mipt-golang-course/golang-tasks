//go:build !solution

package speller

import (
	"github.com/mipt-golang-course/golang-tasks/sprint-1/varjoin"
)

var ones = map[int64]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var teens = map[int64]string{
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tens = map[int64]string{
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

const maxNumber = 1_000_000_000_000

func Spell(n int64) string {
	if (n < -maxNumber+1) || (n > maxNumber-1) {
		return "invalid n"
	}

	if n == 0 {
		return "zero"
	}

	n, minus := SpellMinus(n)
	var parts []string
	if minus != "" {
		parts = append(parts, minus)
	}

	var (
		billions      = n / 1_000_000_000
		millions      = (n / 1_000_000) % 1_000
		thousands     = (n / 1_000) % 1_000
		belowThousand = n % 1_000
	)

	if billions > 0 {
		parts = append(parts, SpellBelowThousand(billions), "billion")
	}

	if millions > 0 {
		parts = append(parts, SpellBelowThousand(millions), "million")
	}

	if thousands > 0 {
		parts = append(parts, SpellBelowThousand(thousands), "thousand")
	}

	if belowThousand > 0 {
		parts = append(parts, SpellBelowThousand(belowThousand))
	}

	return varjoin.Join(" ", parts...)
}

func SpellMinus(n int64) (int64, string) {
	if n < 0 {
		return -n, "minus"
	}

	return n, ""
}

func SpellBelowThousand(n int64) string {
	var parts []string

	if n >= 100 {
		hundreds := n / 100
		parts = append(parts, ones[hundreds], "hundred")
		n %= 100
	}

	if n >= 20 {
		tensVal := (n / 10) * 10
		if n%10 != 0 {
			parts = append(parts, tens[tensVal]+"-"+ones[n%10])
		} else {
			parts = append(parts, tens[tensVal])
		}
		n = 0
	}

	if n >= 10 {
		parts = append(parts, teens[n])
		n = 0
	}

	if n > 0 {
		parts = append(parts, ones[n])
	}

	return varjoin.Join(" ", parts...)
}
