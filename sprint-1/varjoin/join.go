//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	ret := ""
	if len(args) > 0 {
		for _, val := range args[:len(args)-1] {
			ret += val + sep
		}
		ret += args[len(args)-1]
	}

	return ret
}
