//go:build !solution

package varjoin

import "strings"

func Join(sep string, args ...string) string {
	if len(args) == 0 {
		return ""
	}

	str := strings.Builder{}
	str.WriteString(args[0])
	for _, val := range args[1:] {
		str.WriteString(sep)
		str.WriteString(val)
	}
	return str.String()
}
