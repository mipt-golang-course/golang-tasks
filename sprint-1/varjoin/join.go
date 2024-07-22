package varjoin

// Join принимает разделитель и переменное число строк и возвращает одну строку-склейку всех аргументов через разделитель.
func Join(sep string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}

	// Начинаем с первой строки, чтобы избежать начального разделителя
	result := parts[0]

	// Итерируем по остальным частям и добавляем их к результату с разделителем
	for _, part := range parts[1:] {
		result += sep + part
	}

	return result
}
