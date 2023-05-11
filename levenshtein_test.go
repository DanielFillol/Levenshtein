package levenshtein

import (
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want int
	}{
		{"kitten", "sitting", 3},
		{"", "foo", 3},
		{"bar", "", 3},
		{"", "", 0},
		{"abcdefg", "ABCDEFG", 7},
		{"book", "back", 2},
	}

	for _, tt := range tests {
		got := Distance(tt.str1, tt.str2)
		if got != tt.want {
			t.Errorf("Distance(%q, %q) = %v, want %v", tt.str1, tt.str2, got, tt.want)
		}
	}
}
