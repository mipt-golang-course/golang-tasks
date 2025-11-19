//go:build !solution

package varjoin

import (
	"strings"
)

func Join(sep string, args ...string) string {
	var str strings.Builder

	str.Grow(len(args) * 2)

	for _, value := range args {
		str.WriteString(value)
		str.WriteString(sep)
	}
	return strings.TrimRight(str.String(), sep)

}
