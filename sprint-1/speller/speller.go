//go:build !solution

package speller

import "fmt"

func upToTen(n int64) string {
	if n < 0 || n > 9 {
		panic("invalid N passed: %d\n" + fmt.Sprintf("%d", n))
	}

	strs := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return strs[n]
}

func teens(n int64) string {
	if n < 10 || n > 19 {
		panic("invalid N passed: %d\n" + fmt.Sprintf("%d", n))
	}

	strs := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	return strs[n-10]
}

func upToHundred(n int64) string {
	if n > 99 {
		panic("invalid N passed: %d\n" + fmt.Sprintf("%d", n))
	}

	switch {
	case n < 10:
		return upToTen(n)
	case n >= 10 && n < 20:
		return teens(n)
	}

	strs := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	ret := strs[n/10-2]
	if n%10 != 0 {
		ret += "-" + upToTen(n%10)
	}

	return ret
}

func upToThousand(n int64) string {
	if n > 999 {
		panic("invalid N passed: %d\n" + fmt.Sprintf("%d", n))
	}

	if n < 100 {
		return upToHundred(n)
	}

	ret := upToTen(n/100) + " hundred"
	if n%100 != 0 {
		ret += " " + upToHundred(n%100)
	}

	return ret
}

func Spell(n int64) string {
	if n == 0 {
		return "zero"
	}

	was_neg := false

	if n < 0 {
		n = -n
		was_neg = true
	}

	powers := []string{"", "thousand", "million", "billion", "trillion"}

	ret := ""
	current_pow := 0

	for n != 0 {
		if n%1000 != 0 {
			new_str := upToThousand(n % 1000)
			if current_pow > 0 {
				new_str = new_str + " " + powers[current_pow]
			}

			if len(ret) > 0 {
				new_str = new_str + " " + ret
			}

			ret = new_str
		}

		n /= 1000
		current_pow += 1
	}

	if was_neg {
		ret = "minus " + ret
	}

	return ret
}
