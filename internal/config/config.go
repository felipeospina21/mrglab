// Package config handles loading and watching the mrglab TOML configuration file.
package config

import (
	"errors"
	"flag"
	"fmt"

	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Project represents a GitLab project entry in the config file.
type Project struct {
	Name     string `mapstructure:"name"`
	FullPath string `mapstructure:"fullPath"`
	ID       string `mapstructure:"id"`
}

// Filter holds the project filter list from the config file.
type Filter struct {
	Projects []Project `mapstructure:"projects"`
}
// Config holds the application configuration loaded from the TOML file and environment.
type Config struct {
	BaseURL  string `mapstructure:"base_url"`
	APIToken string `mapstructure:"token"`
	Filters  Filter `mapstructure:"filters"`
	DevMode  bool
}

// GlobalConfig is the singleton config instance used throughout the application.
var (
	GlobalConfig Config
	cmdName      = "mrglab"
)

// Load reads the config file, unmarshals it, and loads environment variables.
// In dev mode, missing config file and token are tolerated and mock projects are used as fallback.
func Load(config *Config) error {
	config.DevMode = isDevMode()

	l, f := logger.New(logger.NewLogger{})
	defer f.Close()

	viper.SetConfigName(cmdName)
	viper.SetConfigType("toml")

	viper.AddConfigPath(fmt.Sprintf("$HOME/.config/%s/", cmdName))
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if config.DevMode {
			config.BaseURL = "https://gitlab.com"
			config.Filters.Projects = mockProjects
			return nil
		}
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal: %w", err))
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://gitlab.com"
	}

	if config.DevMode {
		config.Filters.Projects = mockProjects
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		l.Info(fmt.Sprintf("Config file changed: %s", e.Name))
	})
	viper.WatchConfig()

	// Env vars
	err = loadEnvVars(config)
	if err != nil && !config.DevMode {
		return err
	}
	return nil
}

func loadEnvVars(config *Config) error {
	l, f := logger.New(logger.NewLogger{})
	defer f.Close()

	viper.SetEnvPrefix(cmdName)
	err := viper.BindEnv("token")
	if err != nil {
		l.Error(err)
	}
	token := viper.Get("token")

	if token == nil {
		err := errors.New("api-token not set")
		l.Error(err)
		config.APIToken = ""
		return err
	}

	config.APIToken = token.(string)
	return nil
}

var devFlag = flag.Bool("dev", false, "activates dev mode to use mocked data instead of calling api")

func isDevMode() bool {
	if !flag.Parsed() {
		flag.Parse()
	}
	return *devFlag
}

var mockProjects = []Project{
	{Name: "Payments API", ID: "10234567", FullPath: "acme-corp/payments-api"},
	{Name: "Web Dashboard", ID: "10234568", FullPath: "acme-corp/web-dashboard"},
	{Name: "Auth Service", ID: "10234569", FullPath: "acme-corp/auth-service"},
	{Name: "Mobile App", ID: "10234570", FullPath: "acme-corp/mobile-app"},
}
