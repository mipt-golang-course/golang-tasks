package balance

import "testing"

func TestBalance(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"only letters", "abcXYZ", true},
		{"simple valid parentheses", "()", true},
		{"multiple valid", "()[]{}", true},
		{"nested valid", "({[]})", true},
		{"valid with letters", "a(b[c]{d}e)f", true},

		{"wrong order", "(]", false},
		{"mismatched", "([)]", false},
		{"extra opening", "(((", false},
		{"extra closing", "())", false},
		{"closing first", ")", false},
		{"mixed with letters wrong", "a(b]c)", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Balance(tt.input)
			if got != tt.expected {
				t.Errorf("Balance(%q) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}
