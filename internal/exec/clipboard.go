// Package exec provides OS-level utilities for clipboard and browser operations.
package exec

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/felipeospina21/mrglab/internal/logger"
)

// CopyToClipboard copies text to the system clipboard.
func CopyToClipboard(text string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard")
	case "windows":
		cmd = exec.Command("clip")
	default:
		return
	}

	cmd.Stdin = strings.NewReader(text)
	if err := cmd.Run(); err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)
	}
}
