//go:build !solution

package speller

func twenty(n int64) string {
	switch n {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	}
	return "unknown"
}

func dozens(n int64) string {
	switch n {
	case 10:
		return "ten"
	case 20:
		return "twenty"
	case 30:
		return "thirty"
	case 40:
		return "forty"
	case 50:
		return "fifty"
	case 60:
		return "sixty"
	case 70:
		return "seventy"
	case 80:
		return "eighty"
	case 90:
		return "ninety"
	}
	return "unknown"
}

func dozenNums(n int64) string {
	if n%10 == 0 {
		return dozens(n - n%10)
	}
	if n < 20 {
		return twenty(n)
	}
	return dozens(n-n%10) + "-" + twenty(n%10)
}

func nextNums(n int64) string {
	var order int64
	var word string

	switch {
	case n < 100:
		return Spell(n)
	case n < 1000:
		order = 100
		word = "hundred"
	case n < 1000000:
		order = 1000
		word = "thousand"
	case n < 1000000000:
		order = 1000000
		word = "million"
	default:
		order = 1000000000
		word = "billion"
	}

	if n%order == 0 {
		return twenty(n/order) + " " + word
	}
	return Spell(n/order) + " " + word + " " + nextNums(n%order)
}

func Spell(n int64) string {
	var result string

	if n < 0 {
		result = "minus "
		n *= -1
	}

	switch {
	case n < 20:
		result += twenty(n)
	case n < 100:
		result += dozenNums(n)
	default:
		result += nextNums(n)
	}

	return result
}
