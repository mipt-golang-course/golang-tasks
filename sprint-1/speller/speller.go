//go:build !solution

package speller

import (
	"strings"
)

func Spell(num int64) string {
	if num == 0 {
		return "zero"
	}

	ones := []string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
		6: "six", 7: "seven", 8: "eight", 9: "nine"}
	teens := []string{10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen",
		16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen"}
	tens := []string{2: "twenty", 3: "thirty", 4: "forty", 5: "fifty",
		6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety"}

	var str strings.Builder

	if num < 0 {
		str.WriteString("minus")
		num = -num
	}

	convertLessThanHundred := func(n int) string {
		if n == 0 {
			return ""
		}

		var res strings.Builder
		switch {
		case n < 10:
			res.WriteString(ones[n])
		case n < 20:
			res.WriteString(teens[n])
		default:
			res.WriteString(tens[n/10])
			if n%10 != 0 {
				res.WriteString("-")
				res.WriteString(ones[n%10])
			}
		}
		return res.String()
	}

	convert := func(n int) string {
		var res strings.Builder
		if n >= 100 {
			res.WriteString(ones[n/100])
			res.WriteString(" hundred")
			if remainder := n % 100; remainder > 0 {
				res.WriteString(" ")
				res.WriteString(convertLessThanHundred(remainder))
			}
		} else {
			res.WriteString(convertLessThanHundred(n))
		}
		return res.String()
	}

	type scale struct {
		value int64
		name  string
	}

	scales := []scale{
		{1e9, "billion"},
		{1e6, "million"},
		{1e3, "thousand"},
		{1, ""},
	}

	for _, s := range scales {
		if num >= s.value {
			chunk := int(num / s.value)
			num %= s.value

			if str.Len() > 0 {
				str.WriteString(" ")
			}

			str.WriteString(convert(chunk))

			if s.name != "" {
				str.WriteString(" ")
				str.WriteString(s.name)
			}
		}
	}

	return str.String()
}
