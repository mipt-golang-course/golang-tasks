//go:build !solution

package varjoin

import "strings"

func Join(sep string, args ...string) string {
	switch argc := len(args); argc {
	case 0:
		return ""
	case 1:
		return args[0]
	default:
		sepLen := len(sep)
		totalLen := (argc - 1) * sepLen

		for _, s := range args {
			totalLen += len(s)
		}

		var output strings.Builder
		output.Grow(totalLen)
		output.WriteString(args[0])

		if sepLen > 0 {
			for i := 1; i < argc; i++ {
				output.WriteString(sep)
				output.WriteString(args[i])
			}
		} else {
			for i := 1; i < argc; i++ {
				output.WriteString(args[i])
			}
		}

		return output.String()
	}
}
