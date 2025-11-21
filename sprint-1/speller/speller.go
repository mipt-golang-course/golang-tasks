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

var checking_after_hundreds = []struct {
	value int64
	spell string
}{
	{
		value: 1_000_000_000,
		spell: "billion",
	},
	{
		value: 1_000_000,
		spell: "million",
	},
	{
		value: 1_000,
		spell: "thousand",
	},
	{
		value: 100,
		spell: "hundred",
	},
}

func Spell(n int64) string {
	switch {
	case n < 0:
		return "minus " + Spell(-n)
	case n == 0:
		return "zero"
	default:
		var bstr strings.Builder
		spell, _ := spelling(n, bstr)
		return spell.String()
	}
}

func spelling(n int64, bstr strings.Builder) (strings.Builder, bool) {

	switch {
	case n < 10:
		head_part, ok := spell_ones[n]
		bstr.WriteString(head_part)
		return bstr, ok

	case n < 20:
		head_part, ok := spell_teens[n]
		bstr.WriteString(head_part)
		return bstr, ok

	case n < 100:

		head_part := spell_tens[n/10]
		tail_part, tail_ok := spell_ones[n%10]

		bstr.WriteString(head_part)
		if tail_ok {
			bstr.WriteString("-")
			bstr.WriteString(tail_part)
			return bstr, true
		}
		return bstr, true

	default:
		for _, check := range checking_after_hundreds {
			if n >= check.value {
				head_part, _ := spelling(n/check.value, bstr)
				tail_part, tail_ok := spelling(n%check.value, bstr)

				bstr.WriteString(head_part.String())
				bstr.WriteString(" ")
				bstr.WriteString(check.spell)

				if tail_ok {
					bstr.WriteString(" ")
					bstr.WriteString(tail_part.String())
					return bstr, true
				}
				return bstr, true
			}
		}
	}
	return bstr, false
}
