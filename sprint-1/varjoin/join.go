//go:build !solution

package varjoin

import (
	"strings"
)

func Join(sep string, args ...string) string {
	var str strings.Builder

	lagrs := len(args)

	if lagrs == 0 {
		return str.String()
	}

	grow := 0
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
