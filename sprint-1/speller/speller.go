//go:build !solution

package speller

import (
	"strconv"
)

var sMap = map[uint64]string{
	0:          "zero",
	1:          "one",
	2:          "two",
	3:          "three",
	4:          "four",
	5:          "five",
	6:          "six",
	7:          "seven",
	8:          "eight",
	9:          "nine",
	10:         "ten",
	11:         "eleven",
	12:         "twelve",
	13:         "thirteen",
	14:         "fourteen",
	15:         "fifteen",
	16:         "sixteen",
	17:         "seventeen",
	18:         "eightteen",
	19:         "nineteen",
	20:         "twenty",
	30:         "thirty",
	40:         "forty",
	50:         "fifty",
	60:         "sixty",
	70:         "seventy",
	80:         "eighty",
	90:         "ninety",
	100:        "hundred",
	1000:       "thousand",
	1000000:    "million",
	1000000000: "billion",
}

// количество разрядов
func digits(n uint64) int {
	var i uint64 = 1
	var k int = 0

	for {
		if (n / i) > 0 {
			k++
			i = i * 10
		} else {
			break
		}
	}
	return int(k)
}

func Spell(n int64) string {

	var d uint64
	var dig int

	result := ""
	sep := "-"
	prefix := ""

	if n < 0 {
		prefix = "minus "
		d = uint64(n - 2*n)
		dig = digits(uint64(d))
	} else {
		d = uint64(n)
		dig = digits(d)
	}


	spellOnes := func(d uint64) string {
		return sMap[d]
	}

	spellTens := func(d uint64) string {
		var res string
		rem := d % 10
		quo := d / 10

		if rem == 0 {
			res = sMap[quo*10]
		} else {
			res = sMap[quo*10] + sep + sMap[rem]
		}
		return res
	}

	spellHundreds := func(d uint64) string {
		var s string
		var res string
		rem := d % 100
		quo := d / 100

		if quo > 0 {
			s = " " + sMap[100]
			switch {
			case quo < 10:
				res = spellOnes(quo) + s
			case 21 <= quo && quo <= 99:
				res = spellTens(quo) + s
			default:
				res = sMap[quo] + s
			}
		} else {
			res = "one"
		}

		if rem != 0 {
			switch {
			case rem < 10:
				res += " " + spellOnes(rem)
			case 21 <= rem && rem <= 99:
				res += " " + spellTens(rem)
			default:
				res += " " + sMap[rem]
			}
		}
		return res
	}

	spellThousands := func(d uint64) string {
		var s string
		var res string
		rem := d % 1000
		quo := d / 1000

		if quo > 0 {
			s = " " + sMap[1000]
			switch {
			case quo < 10:
				res = spellOnes(quo) + s
			case (21 <= quo) && (quo <= 99):
				res = spellTens(quo) + s
			case quo > 99:
				res = spellHundreds(quo) + s
			default:
				res = sMap[quo] + s
			}
		} else {
			res = "one"
		}

		if rem != 0 {
			switch {
			case rem < 10:
				res += " " + spellOnes(rem)
			case 21 <= rem && rem <= 99:
				res += " " + spellTens(rem)
			case rem > 99 && rem < 1000:
				res += " " + spellHundreds(rem)
			default:
				res += " " + sMap[rem]
			}
		}
		return res
	}

	spellMillions := func(d uint64) string {
		var s string
		var res string
		rem := d % 1000000
		quo := d / 1000000

		if quo > 0 {
			s = " " + sMap[1000000]
			switch {
			case quo < 10:
				res = spellOnes(quo) + s
			case (21 <= quo) && (quo <= 99):
				res = spellTens(quo) + s
			case quo > 99 && quo < 1000:
				res = spellHundreds(quo) + s
			case quo > 1000 && quo < 1000000:
				res = spellHundreds(quo) + s
			default:
				res = sMap[quo] + s
			}
		} else {
			res = "one"
		}

		if rem != 0 {
			switch {
			case rem < 10:
				res += " " + spellOnes(rem)
			case 21 <= rem && rem <= 99:
				res += " " + spellTens(rem)
			case rem > 99 && rem < 1000:
				res += " " + spellHundreds(rem)
			case rem >= 1000 && rem < 1000000:
				res += " " + spellThousands(rem)

			default:
				res += " " + sMap[rem]
			}
		}
		return res
	}

	spellBillions := func(d uint64) string {
		var s string
		var res string
		rem := d % 1000000000
		quo := d / 1000000000

		if quo > 0 {
			s = " " + sMap[1000000000]
			switch {
			case quo < 10:
				res = spellOnes(quo) + s
			case 21 <= quo && quo <= 99:
				res = spellTens(quo) + s
			case quo > 99 && quo < 1000:
				res = spellHundreds(quo) + s
			case quo > 1000 && quo < 1000000:
				res = spellHundreds(quo) + s
			case quo > 1000000 && quo < 1000000000:
				res = spellMillions(quo) + s
			default:
				res = sMap[quo] + s
			}
		} else {
			res = "one"
		}

		if rem != 0 {
			switch {
			case rem < 10:
				res += " " + spellOnes(rem)
			case 21 <= rem && rem <= 99:
				res += " " + spellTens(rem)
			case rem > 99 && rem < 1000:
				res += " " + spellHundreds(rem)
			case rem >= 1000 && rem < 1000000:
				res += " " + spellThousands(rem)
			case rem >= 1000000 && rem < 1000000000:
				res += " " + spellMillions(rem)

			default:
				res += " " + sMap[rem]
			}
		}
		return res
	}

	switch {

	case d <= 20:
		result = prefix + spellOnes(d)

	case (21 < d) && (d <= 99):
		result = prefix + spellTens(d)

	case (3 <= dig) && (dig < 4):
		result = prefix + spellHundreds(d)

	case (4 <= dig) && (dig < 7):
		result = prefix + spellThousands(d)

	case (7 <= dig) && (dig < 9):
		result = prefix + spellMillions(d)

	case (9 <= dig) && (dig <= 12):
		result = prefix + spellBillions(d)

	default:
		result = "количество разрядов: " + strconv.Itoa(dig)
	}
	return result
}
