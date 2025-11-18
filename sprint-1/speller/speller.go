//go:build !solution

package speller

import (
	"strings"
)

var spell_ones = map[int64]string{
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

var spell_teens = map[int64]string{
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

var spell_tens = map[int64]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

func Spell(n int64) string {
	if n < 0 {
		return "minus " + Spell(-n)
	}

	switch {

	case n == 0:
		return "zero"

	case n < 100:
		return spelling_to_hundred(n)

	case n >= 100:
		return spelling_after_hundred(n)

	default:
		return ""
	}
}

func spelling_to_hundred(n int64) string {
	switch {
	case n < 10:
		{
			return spell_ones[n]
		}
	case n < 20:
		{
			return spell_teens[n]
		}
	case n < 100:
		{
			v := spell_tens[n/10] + "-" + spell_ones[n%10]
			return strings.TrimRight(v, "-")
		}
	default:
		return ""
	}
}

func spelling_after_hundred(n int64) string {
	switch {
	case n < 100:
		v := spelling_to_hundred(n)
		return strings.TrimSpace(v)
	case n < 1000:
		v := spelling_to_hundred(n/100) + " hundred " + spelling_to_hundred(n%100)
		return strings.TrimSpace(v)

	case n < 1000000:
		v := spelling_after_hundred(n/1000) + " thousand " + spelling_after_hundred(n%1000)
		return strings.TrimSpace(v)

	case n < 1000000000:
		v := spelling_after_hundred(n/1000000) + " million " + spelling_after_hundred(n%1000000)
		return strings.TrimSpace(v)

	case n < 1000000000000:
		v := spelling_after_hundred(n/1000000000) + " billion " + spelling_after_hundred(n%1000000000)
		return strings.TrimSpace(v)

	default:
		return ""
	}
}