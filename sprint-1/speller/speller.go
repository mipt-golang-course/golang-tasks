package speller

var oneNumbers = []string{
	"zero", "one", "two", "three", "four", "five", "six",
	"seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen",
	"sixteen", "seventeen", "eighteen", "nineteen",
}

var tenNumbers = []string{
	"", "", "twenty", "thirty", "forty", "fifty",
	"sixty", "seventy", "eighty", "ninety",
}

type bigNumber struct {
	value int64
	name  string
}

var bigNumbers = []bigNumber{
	{1000000000000, "trillion"},
	{1000000000, "billion"},
	{1000000, "million"},
	{1000, "thousand"},
}

func Spell(n int64) string {
	if n < 0 {
		return "minus " + Spell(-n)
	}

	switch {
	case n == 0:
		return "zero"

	case n < 20:
		return oneNumbers[n]

	case n < 100:
		tensPart := tenNumbers[n/10]
		if n%10 == 0 {
			return tensPart
		}
		return tensPart + "-" + oneNumbers[n%10]

	case n < 1000:
		h := n / 100
		r := n % 100
		if r == 0 {
			return oneNumbers[h] + " hundred"
		}
		return oneNumbers[h] + " hundred " + Spell(r)

	default:
		for _, v := range bigNumbers {
			if n >= v.value {
				h := n / v.value
				r := n % v.value
				if r == 0 {
					return Spell(h) + " " + v.name
				}
				return Spell(h) + " " + v.name + " " + Spell(r)
			}
		}
	}

	return ""
}
