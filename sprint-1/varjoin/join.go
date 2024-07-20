//go:build !solution

package varjoin

import (
	"fmt"
	"strings"
)

func Join(sep string, args ...string) string {
	args_num := len(args)

	switch {
	case args_num == 0:
		return ""
	case args_num == 1:
		return args[0]
	case args_num > 1:
		var string_len int
		sep_len := len(sep)

		for _, str := range args {
			string_len += (len(str) + sep_len)
		}

		var b strings.Builder

		b.WriteString(args[0])
		for i := 1; i < args_num; i++ {
			b.WriteString(sep)
			b.WriteString(args[i])
		}

		return b.String()

	default:
		fmt.Println("Reached unreachable")
		return ""
	}

}
