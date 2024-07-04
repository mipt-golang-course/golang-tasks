//go:build !solution

package speller

import (
	"slices"
	"strings"
)

var digitsSpelling = map[int64]string{
	0: "",
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

var tensSpelling = map[int64]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var specialNumbersSpelling = map[int64]string{
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

var powersOfThousandSpelling = map[int]string{
	0: "",
	1: "thousand",
	2: "million",
	3: "billion",
}

func Spell(n int64) string {
	if n == 0 {
		return "zero"
	}
	var stringBuilder strings.Builder
	if n < 0 {
		stringBuilder.WriteString("minus ")
	}
	triplets := divideIntoTriplets(n)
	for i, triplet := range triplets {
		spelledTriplet := spellTriplet(triplet)
		if len(spelledTriplet) != 0 {
			for _, word := range spelledTriplet {
				stringBuilder.WriteString(word)
				stringBuilder.WriteString(" ")
			}
			stringBuilder.WriteString(powersOfThousandSpelling[len(triplets)-i-1])
			stringBuilder.WriteString(" ")
		}
	}
	return strings.Trim(stringBuilder.String(), " ")
}

func divideIntoTriplets(n int64) []int64 {
	triplets := make([]int64, 0)
	n = getAbsoluteValue(n)
	for {
		triplets = append(triplets, n%1000)
		n /= 1000
		if n == 0 {
			break
		}
	}
	slices.Reverse(triplets)
	return triplets
}

func getAbsoluteValue(n int64) int64 {
	if n < 0 {
		n = -n
	}
	return n
}

func spellTriplet(n int64) []string {
	words := make([]string, 0)
	if n == 0 {
		return words
	}
	if n/100 != 0 {
		words = append(words, digitsSpelling[n/100], "hundred")
	}
	n %= 100
	switch {
	case n/10 == 0:
		words = append(words, digitsSpelling[n%10])
	case n/10 == 1:
		words = append(words, specialNumbersSpelling[n])
	case n%10 == 0:
		words = append(words, tensSpelling[n/10])
	default:
		words = append(words, tensSpelling[n/10]+"-"+digitsSpelling[n%10])
	}
	return words
}
