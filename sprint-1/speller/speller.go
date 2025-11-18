//go:build !solution

package speller

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
			v, ok := spell_ones[n%10]
			if ok {
				v = "-" + v
			}
			return spell_tens[n/10] + v
		}
	case n < 1000:
		{
			return spell_ones[n/100] + " hundred " + Spell(n%100)
		}
	default:
		return ""
	}
}
