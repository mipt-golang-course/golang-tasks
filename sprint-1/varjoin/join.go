//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	if len(args) == 0 {
		return ""
	}

	accum := args[0]
	for i := range args[1:] {
		accum += sep + args[i+1]
	}

	return accum
}
