package varjoin

func Join(sep string, args ...string) string {
	if len(args) == 0 {
		return ""
	}

	result := ""
	for k, v := range args {
		if k > 0 {
			result += sep
		}
		result += v
	}
	return result
}
