//go:build !solution

package varjoin

import "strings"

func Join(sep string, args ...string) string {
	str := ""
	for _, value := range args {
		str = str + value + sep
	}
	return strings.TrimRight(str, sep)
}
