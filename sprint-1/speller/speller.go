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

var checking = []struct {
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
		spell, _ := spelling(n)
		return spell
	}
}

func spelling(n int64) (string, bool) {
	var spell strings.Builder

	switch {
	case n < 10:
		head_part, ok := spell_ones[n]
		return head_part, ok

	case n < 20:
		head_part, ok := spell_teens[n]
		return head_part, ok

	case n < 100:

		head_part := spell_tens[n/10]
		tail_part, tail_ok := spell_ones[n%10]

		spell.WriteString(head_part)
		if tail_ok {
			spell.WriteString("-")
			spell.WriteString(tail_part)
			return spell.String(), true
		}
		return spell.String(), true

	default:
		for _, check := range checking {
			if n >= check.value {
				head_part, _ := spelling(n / check.value)
				tail_part, tail_ok := spelling(n % check.value)

				spell.WriteString(head_part)
				spell.WriteString(" ")
				spell.WriteString(check.spell)

				if tail_ok {
					spell.WriteString(" ")
					spell.WriteString(tail_part)
					return spell.String(), true
				}
				return spell.String(), true
			}
		}
	}
	return spell.String(), false
}
