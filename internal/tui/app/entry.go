// Package app wires together all TUI components into the main Bubble Tea model.
//
// NewApp returns a ready-to-use tea.Model for embedding in a launcher or running standalone.
package app

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/gitlab"
)

// NewApp loads config, creates the GitLab client, and returns the initialized app model.
func NewApp() tea.Model {
	ctx := &context.AppContext{}

	err := config.Load(&config.GlobalConfig)
	if err != nil {
		fmt.Println("Error loading config", err)
		os.Exit(1)
	}

	client := gitlab.NewClient(&config.GlobalConfig)
	return InitMainModel(ctx, &config.GlobalConfig, client)
}
