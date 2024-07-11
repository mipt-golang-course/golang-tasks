//go:build !solution

package varjoin

import (
	"strings"
)

func Join(sep string, args ...string) string {
	return strings.Join(args, sep)
}
