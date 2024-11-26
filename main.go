package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/app"
)

func main() {
	ctx := &context.AppContext{}

	err := config.Load(&config.GlobalConfig)
	if err != nil {
		fmt.Println("Error loading api token", err)
		os.Exit(1)
	}

	m := app.InitMainModel(ctx)

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
