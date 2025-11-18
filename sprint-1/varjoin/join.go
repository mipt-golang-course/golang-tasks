//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	str := ""
	if len(args) < 1 {
		return str
	}
	for _, value := range args[:len(args)-1] {
		str = str + value + sep
	}
	return str + args[len(args)-1]
}
