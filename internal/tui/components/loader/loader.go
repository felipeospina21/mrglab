// Package loader provides a reusable loading indicator component.
package loader

import (
	tsloader "github.com/felipeospina21/tuishell/loader"
	"github.com/felipeospina21/tuishell/style"
)

// View renders a spinner frame with a "Loading..." label.
func View(spinnerView string) string {
	return tsloader.View(style.DefaultTheme(), spinnerView)
}
