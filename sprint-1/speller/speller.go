//go:build !solution

package speller

import "strings"

func Spell(num int64) string {
	if num == 0 {
		return "zero"
	}

	dictToTwenty := map[int]string{
		1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
		10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen"}

	dictTens := map[int]string{
		2: "twenty", 3: "thirty", 4: "forty", 5: "fifty", 6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety"}

	str := strings.Builder{}

	if num < 0 {
		str.WriteString("minus")
		num = -num
	}

	convertLessThanHundred := func(num int) string {
		res := strings.Builder{}
		if num/10 >= 2 {
			if num%10 == 0 {
				res.WriteString(dictTens[num/10])
			} else {
				res.WriteString(dictTens[num/10])
				res.WriteString("-")
				res.WriteString(dictToTwenty[num%10])
			}
		} else {
			res.WriteString(dictToTwenty[num])
		}
		return res.String()
	}

	convert := func(num int) string {
		res := strings.Builder{}
		if num >= 100 {
			if res.String() != "" {
				res.WriteString(" ")
			}
			res.WriteString(convertLessThanHundred(num / 100))
			res.WriteString(" hundred")
			num %= 100
		}
		if num > 0 {
			if res.String() != "" {
				res.WriteString(" ")
			}
			res.WriteString(convertLessThanHundred(num))
		}
		return res.String()
	}

	type scales_t struct {
		value int64
		name  string
	}

	scales := []scales_t{
		{1e9, "billion"},
		{1e6, "million"},
		{1e3, "thousand"},
		{1, ""},
	}

	for _, scale := range scales {
		if num >= scale.value {
			chunk := int(num / scale.value)
			num %= scale.value
			if str.String() != "" {
				str.WriteString(" ")
			}
			str.WriteString(convert(chunk))
			if scale.name != "" {
				str.WriteString(" ")
				str.WriteString(scale.name)
			}
		}
	}
	return str.String()
}
