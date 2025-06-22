//func Join(sep string, args ...string) string {
//	str := ""
//
//	for i, val := range args {
//		str = str + val
//		if i < len(args)-1 {
//			str = str + sep
//		}
//	}
//	return str
//}

//func Join(sep string, args ...string) string {
//	if len(args) == 0 {
//		return ""
//	}
//
//	new_str_len := (len(args) - 1) * len(sep)
//	for _, val := range args {
//		new_str_len += len(val)
//	}
//
//	new_str := make([]byte, 0, new_str_len)
//	new_str = append(new_str, args[0]...)
//	for _, val := range args[1:] {
//		new_str = append(new_str, sep...)
//		new_str = append(new_str, val...)
//	}
//	return string(new_str)
//}

//func Join(sep string, args ...string) string {
//	switch {
//	case len(args) == 0:
//		return ""
//	case len(args) == 1:
//		return args[0]
//	default:
//		return args[0] + sep + Join(sep, args[1:]...)
//	}
//}
