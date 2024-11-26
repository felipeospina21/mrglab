package exec

import (
	"errors"
	"os/exec"
	"runtime"

	"github.com/felipeospina21/mrglab/internal/logger"
)

func Openbrowser(url string) {
	var err error
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		errors.New("unsupported platform")
	}

	err = cmd.Start()
	if err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)
	}
	err = cmd.Wait()
	if err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)
	}
}
