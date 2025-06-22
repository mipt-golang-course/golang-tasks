//go:build !solution

package varjoin

import (
	"strings"
)

func Join(sep string, args ...string) string {
	n := len(args)
	if n == 0 {
		return ""
	}

	output := strings.Builder{}
	output.Grow(TotalLen(sep, args...))

	output.WriteString(args[0])

	if len(sep) == 0 {
		for i := 1; i < n; i++ {
			output.WriteString(args[i])
		}
	} else {
		for i := 1; i < n; i++ {
			output.WriteString(sep)
			output.WriteString(args[i])
		}
	}

	return output.String()
}

func TotalLen(sep string, args ...string) int {
	totalLen := 0
	for i := range args {
		totalLen += len(args[i])
	}

	totalLen += len(sep) * (len(args) - 1)

	return totalLen
}
