package speller

var (
	oneToNine         = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tenToNineteen = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	multipleOfTen    = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	endingName = []string{"", " thousand", " million", " billion", " trillion", " quadrillion", " quintillion"}
)

func OneRegister(n int64) string{
	return(oneToNine[n])
}

func TwoRegister(n int64) string{
	switch{
		case n < 10: {
			return(OneRegister(n))
		}
		case n < 19: {
			return(tenToNineteen[n-10])
		}
		default: {
			if n % 10 != 0{
				return(multipleOfTen[n/10-2] + "-" + oneToNine[n%10])
			} else {
				return(multipleOfTen[n/10-2])
			}
		}
	}
}

func Spell(n int64) string {

	//check sign
	var sign string
	switch {
		case n ==0: {
			return("zero")
		}
		case n < 0: {
			sign= "minus"
			n = -n
		}
		default: {
			sign= ""
		}
	}

	var result string
	var plus string
	var count=0
	var rest int64

	for number:=n; number >0 ;{
		rest = number % 1000
			switch{
				case rest ==0:{
					plus = ""
				}
				case rest< 10: {
					plus = OneRegister(rest) + endingName[count]
				}
				case rest <100: {
					plus = TwoRegister(rest) + endingName[count]
				}
				default: {
					if TwoRegister(rest%100) == ""{
						plus = OneRegister(rest/100) + " hundred"
					} else {
						plus = OneRegister(rest/100) + " hundred " + TwoRegister(rest%100) + endingName[count]

					}
				}
			}
			if plus == ""{
				
			} else if result == ""{
				result = plus
			} else {
				result = plus + " " + result
			}
			
		number = number/1000
		count++
	}

	//build result with sign
	if sign != ""{
		return sign + " " + result
	} else {
		return result
	}

}