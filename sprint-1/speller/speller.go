//go:build !solution

package speller

import "strings"

const maxNumber = 1_000_000_000_000

func Spell(n int64) string {
	if n < -maxNumber+1 || n > maxNumber-1 {
		return "invalid n"
	}

	if n == 0 {
		return "zero"
	}

	var out string

	if n < 0 {
		n = -n
		out = "minus"
	}

	billions := (n / 1_000_000_000)
	millions := (n / 1_000_000) % 1_000
	thousands := (n / 1_000) % 1_000
	rem := n % 1_000

	if billions > 0 {
		out = join(" ", out, spellThousand(billions), "billion")
	}

	if millions > 0 {
		out = join(" ", out, spellThousand(millions), "million")
	}

	if thousands > 0 {
		out = join(" ", out, spellThousand(thousands), "thousand")
	}

	if rem > 0 {
		out = join(" ", out, spellThousand(rem))
	}

	return out
}

func spellThousand(n int64) string {
	var out string

	q := n / 100
	r := n % 100

	if q > 0 {
		out = join(" ", spellDigit(q), "hundred")
	}

	if n%100 > 0 {
		out = join(" ", out, spellHundred(r))
	}

	return out
}

func spellHundred(n int64) string {
	var out string

	switch {
	case n < 10:
		return spellDigit(n)
	case 10 <= n && n <= 19:
		return spellTenth(n)
	case n < 30:
		out = "twenty"
	case n < 40:
		out = "thirty"
	case n < 50:
		out = "forty"
	case n < 60:
		out = "fifty"
	case n < 70:
		out = "sixty"
	case n < 80:
		out = "seventy"
	case n < 90:
		out = "eighty"
	case n < 100:
		out = "ninety"
	default:
		return ""
	}

	digit := n % 10

	if digit != 0 {
		out = join("-", out, spellDigit(digit))
	}

	return out
}

func spellTenth(n int64) string {
	switch n {
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
	default:
		return ""
	}
}

func spellDigit(n int64) string {
	switch n {
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
	default:
		return ""
	}
}

func join(sep string, args ...string) string {
	switch argc := len(args); argc {
	case 0:
		return ""
	case 1:
		return args[0]
	default:
		sepLen := len(sep)
		totalLen := (argc - 1) * sepLen

		for _, s := range args {
			totalLen += len(s)
		}

		var output strings.Builder
		output.Grow(totalLen)
		output.WriteString(args[0])

		if sepLen > 0 {
			for i := 1; i < argc; i++ {
				if len(args[i]) == 0 {
					continue
				}

				if i == 1 && len(args[0]) == 0 {
					output.WriteString(args[i])
					continue
				}

				output.WriteString(sep)
				output.WriteString(args[i])
			}
		} else {
			for i := 1; i < argc; i++ {
				output.WriteString(args[i])
			}
		}

		return output.String()
	}
}
