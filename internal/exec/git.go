package exec

import (
	"os/exec"
	"strings"
)

// CurrentGitBranch returns the current git branch name, or empty string on error.
func CurrentGitBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
