//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	if len(args) == 0 {
		return ""
	}

	str := args[0]
	for _, val := range args[1:] {
		str += sep + val
	}
	return str
}
