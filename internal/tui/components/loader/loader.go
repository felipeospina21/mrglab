// Package loader provides a reusable loading indicator component.
package loader

import (
	tsloader "github.com/felipeospina21/tuishell/loader"
	"github.com/felipeospina21/tuishell/style"
)

var pkgTheme style.Theme

// SetTheme sets the theme used by the loader package.
func SetTheme(t style.Theme) { pkgTheme = t }

// View renders a spinner frame with a "Loading..." label.
func View(spinnerView string) string {
	return tsloader.View(pkgTheme, spinnerView)
}
