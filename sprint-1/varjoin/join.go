//go:build !solution

package varjoin

import "strings"

func Join(sep string, args ...string) string {
	var stringBuilder strings.Builder
	for i, arg := range args {
		stringBuilder.WriteString(arg)
		if i != len(args)-1 {
			stringBuilder.WriteString(sep)
		}
	}
	return stringBuilder.String()
}
