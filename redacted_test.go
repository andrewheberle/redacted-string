package redacted

import "testing"

func Test_redact(t *testing.T) {
	tests := []struct {
		name string
		s    string
		c    rune
		size uint
		want string
	}{
		{"short", "a", '*', 8, "********"},
		{"long", "this is a long string", '*', 8, "********"},
		{"full", "this is a long string", '*', 0, "*********************"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redact(tt.s, tt.c, tt.size); got != tt.want {
				t.Errorf("redact() = %v, want %v", got, tt.want)
			}
		})
	}
}
