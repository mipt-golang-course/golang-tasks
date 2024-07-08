//go:build !solution

package varjoin

import (
	"strings"
)

// Join объединяет строки через разделитель
func Join(sep string, elements ...string) string {
	return strings.Join(elements, sep)
}
