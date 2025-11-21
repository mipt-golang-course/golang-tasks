//go:build !solution

package varjoin

import (
	"strings"
)

func Join(sep string, args ...string) string {
	lagrs := len(args)

	if lagrs == 0 {
		return ""
	}

	var str strings.Builder
	len_sep := (len(sep) * len(args)) - 1
	grow := len_sep
	for _, i := range args {
		grow += len(i)
	}

	str.Grow(grow)

	for i := 0; i < lagrs-1; i++ {

		str.WriteString(args[i])
		str.WriteString(sep)
	}

	str.WriteString(args[lagrs-1])

	return str.String()

}
