//go:build !solution

package speller

func Spell(n int64) string {
	if n == 0 {
		return "zero"
	}
	if n < 0 {
		return "negative " + Spell(-n)
	}

	ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tens := []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	teens := []string{"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	result := ""
	billions := n / 1000000000
	if billions > 0 {
		result += Spell(billions) + " billion "
		n %= 1000000000
	}

	millions := n / 1000000
	if millions > 0 {
		result += Spell(millions) + " million "
		n %= 1000000
	}

	thousands := n / 1000
	if thousands > 0 {
		result += Spell(thousands) + " thousand "
		n %= 1000
	}

	hundreds := n / 100
	if hundreds > 0 {
		result += ones[hundreds] + " hundred "
		n %= 100
	}

	if n > 0 && result != "" {
		result += "and "
	}

	if n >= 20 {
		result += tens[n/10]
		n %= 10
		if n > 0 {
			result += "-" + ones[n]
		}
	} else if n >= 11 {
		result += teens[n-11]
	} else if n >= 1 {
		result += ones[n]
	}

	return result
}
