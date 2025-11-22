//go:build !solution

package varjoin

func Join1(sep string, args ...string) string {

	result := ""
	s := false
	for _, v := range args {
		if !s {
			result = result + v
			s = true
		} else {
			result = result + sep + v
		}
	}

	return result
}

func Join2(sep string, args ...string) string {

	result := ""
	for i, v := range args {
		if i == 0 {
			result = result + v
		} else {
			result = result + sep + v
		}
	}
	return result
}

func Join3(sep string, args ...string) string {

	result := ""
	s := ""

	for _, v := range args {
		result = result + s + v
		s = sep
	}
	return result
}
