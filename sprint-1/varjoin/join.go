//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	var result string
	for i := 0; i < len(args); i++ {
		result += args[i]
		if i < (len(args) - 1) {
			result += sep
		}
	}
	return result
}
