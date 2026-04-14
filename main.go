package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/tui/app"
)

func main() {
	m := app.NewApp()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
