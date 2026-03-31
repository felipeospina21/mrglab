package table

import (
	"testing"
	"time"

	"github.com/felipeospina21/mrglab/internal/tui/icon"
)

func TestFormatTime(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{"seconds ago", now.Add(-30 * time.Second), "30s"},
		{"minutes ago", now.Add(-5 * time.Minute), "5m"},
		{"hours ago", now.Add(-3 * time.Hour), "3h"},
		{"days ago", now.Add(-2 * 24 * time.Hour), "2d"},
		{"weeks ago", now.Add(-14 * 24 * time.Hour), "2w"},
		{"months ago", now.Add(-60 * 24 * time.Hour), "2M"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatTime(tt.input)
			if got != tt.expected {
				t.Errorf("FormatTime() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestFormatPercentage(t *testing.T) {
	tests := []struct {
		name     string
		input    float32
		expected string
	}{
		{"zero", 0, ""},
		{"non-zero", 75.5, "75.50 %"},
		{"small", 0.01, "0.01 %"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatPercentage(tt.input); got != tt.expected {
				t.Errorf("FormatPercentage(%v) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    float32
		expected string
	}{
		{"zero", 0, ""},
		{"positive", 120, "2 m"},
		{"large", 3600, "60 m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDuration(tt.input); got != tt.expected {
				t.Errorf("FormatDuration(%v) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestColWidth(t *testing.T) {
	tests := []struct {
		name     string
		w, p     int
		expected int
	}{
		{"50 percent of 100", 100, 50, 50},
		{"10 percent of 200", 200, 10, 20},
		{"zero percent", 100, 0, 0},
		{"100 percent", 100, 100, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColWidth(tt.w, tt.p); got != tt.expected {
				t.Errorf("ColWidth(%d, %d) = %d, want %d", tt.w, tt.p, got, tt.expected)
			}
		})
	}
}

func TestRenderIcon(t *testing.T) {
	tests := []struct {
		name     string
		b        bool
		icon     string
		expected string
	}{
		{"true returns icon", true, icon.Check, icon.Check},
		{"false returns empty", false, icon.Check, icon.Empty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RenderIcon(tt.b, tt.icon); got != tt.expected {
				t.Errorf("RenderIcon(%v, %q) = %q, want %q", tt.b, tt.icon, got, tt.expected)
			}
		})
	}
}

func TestGetColIndex(t *testing.T) {
	cols := []Column{
		{Name: "title", Title: "Title"},
		{Name: "author", Title: "Author"},
		{Name: "status", Title: "Status"},
	}
	tests := []struct {
		name     string
		colName  string
		expected int
	}{
		{"found first", "title", 0},
		{"found last", "status", 2},
		{"not found", "missing", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetColIndex(cols, tt.colName); got != tt.expected {
				t.Errorf("GetColIndex(cols, %q) = %d, want %d", tt.colName, got, tt.expected)
			}
		})
	}
}

func TestParseTimeString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		zero  bool
	}{
		{"valid RFC3339", "2024-01-15T10:30:00Z", false},
		{"invalid string", "not-a-date", true},
		{"empty string", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseTimeString(tt.input)
			if tt.zero && !got.IsZero() {
				t.Errorf("ParseTimeString(%q) should be zero time", tt.input)
			}
			if !tt.zero && got.IsZero() {
				t.Errorf("ParseTimeString(%q) should not be zero time", tt.input)
			}
		})
	}
}
