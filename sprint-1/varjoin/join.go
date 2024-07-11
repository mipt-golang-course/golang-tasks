//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	var result string
	for _, line := range args {
		if result != "" {
			result = result + sep + line
		} else {
			result = line
		}
	}
	return result
}
