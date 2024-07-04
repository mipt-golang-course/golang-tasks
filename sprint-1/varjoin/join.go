//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	var result string
	if len(args) != 0 {
		i := 0
		for ; i < len(args)-1; i++ {
			result = result + args[i] + sep
		}
		result = result + args[i]
	}
	return result
}
