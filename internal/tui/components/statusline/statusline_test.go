package statusline

import (
	"testing"

	"github.com/felipeospina21/mrglab/internal/tui/style"
)

func TestModeBackground(t *testing.T) {
	tests := []struct {
		name   string
		status string
		want   string
	}{
		{"normal", ModesEnum.Normal, style.StatuslineModeNormal},
		{"loading", ModesEnum.Loading, style.StatuslineModeLoading},
		{"error", ModesEnum.Error, style.StatuslineModeError},
		{"demo", ModesEnum.Demo, style.StatuslineModeDev},
		{"unknown falls back to normal", "UNKNOWN", style.StatuslineModeNormal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := modeBackground(tt.status)
			if got == nil {
				t.Fatal("modeBackground returned nil")
			}
		})
	}
}
