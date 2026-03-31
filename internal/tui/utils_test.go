package tui

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a greater", 5, 3, 5},
		{"b greater", 3, 5, 5},
		{"equal", 4, 4, 4},
		{"negative", -1, -5, -1},
		{"zero and positive", 0, 3, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.a, tt.b); got != tt.expected {
				t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a smaller", 3, 5, 3},
		{"b smaller", 5, 3, 3},
		{"equal", 4, 4, 4},
		{"negative", -1, -5, -5},
		{"zero and positive", 0, 3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.a, tt.b); got != tt.expected {
				t.Errorf("Min(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name           string
		v, low, high   int
		expected       int
	}{
		{"within range", 5, 0, 10, 5},
		{"below low", -1, 0, 10, 0},
		{"above high", 15, 0, 10, 10},
		{"at low", 0, 0, 10, 0},
		{"at high", 10, 0, 10, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.v, tt.low, tt.high); got != tt.expected {
				t.Errorf("Clamp(%d, %d, %d) = %d, want %d", tt.v, tt.low, tt.high, got, tt.expected)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		limit    int
		expected string
	}{
		{"short string", "hello", 20, "hello"},
		{"at threshold", "exactly twenty chars", 20, "exactly twenty chars..."},
		{"long string high limit", "this is a string that is longer than twenty characters", 25, "this is a string that is ..."},
		{"limit below 20 uses 20", "short", 5, "short"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.s, tt.limit); got != tt.expected {
				t.Errorf("Truncate(%q, %d) = %q, want %q", tt.s, tt.limit, got, tt.expected)
			}
		})
	}
}
