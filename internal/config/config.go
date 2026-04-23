// Package config handles loading and watching the mrglab TOML configuration file.
package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

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
// ThemeOverrides holds optional color overrides for the theme.
// Each field is a hex color string (e.g. "#FF5733"). Nil means use default.
type ThemeOverrides struct {
	Preset          *string `mapstructure:"preset"`
	Primary         *string `mapstructure:"primary"`
	PrimaryBright   *string `mapstructure:"primary_bright"`
	PrimaryFg       *string `mapstructure:"primary_fg"`
	PrimaryDim      *string `mapstructure:"primary_dim"`
	Info            *string `mapstructure:"info"`
	InfoBright      *string `mapstructure:"info_bright"`
	Success         *string `mapstructure:"success"`
	SuccessBright   *string `mapstructure:"success_bright"`
	Danger          *string `mapstructure:"danger"`
	DangerBright    *string `mapstructure:"danger_bright"`
	Warning         *string `mapstructure:"warning"`
	WarningBright   *string `mapstructure:"warning_bright"`
	Caution         *string `mapstructure:"caution"`
	Text            *string `mapstructure:"text"`
	TextInverse     *string `mapstructure:"text_inverse"`
	TextDimmed      *string `mapstructure:"text_dimmed"`
	Muted           *string `mapstructure:"muted"`
	Dim             *string `mapstructure:"dim"`
	Border          *string `mapstructure:"border"`
	ModalBorder     *string `mapstructure:"modal_border"`
	SurfaceDim      *string `mapstructure:"surface_dim"`
	SelectionBorder *string `mapstructure:"selection_border"`
	StatusText      *string `mapstructure:"status_text"`
	StatusNormal    *string `mapstructure:"status_normal"`
	StatusLoading   *string `mapstructure:"status_loading"`
	StatusError     *string `mapstructure:"status_error"`
	StatusDemo      *string `mapstructure:"status_demo"`
	StatusAccent1   *string `mapstructure:"status_accent1"`
	StatusAccent2   *string `mapstructure:"status_accent2"`
}

// Config holds the application configuration loaded from the TOML file and environment.
type Config struct {
	BaseURL  string         `mapstructure:"base_url"`
	APIToken string         `mapstructure:"token"`
	Filters  Filter         `mapstructure:"filters"`
	Theme    ThemeOverrides `mapstructure:"theme"`
	DemoMode bool
}

// GlobalConfig is the singleton config instance used throughout the application.
var (
	GlobalConfig Config
	cmdName      = "mrglab"
)

// Load reads the config file, unmarshals it, and loads environment variables.
// In dev mode, missing config file and token are tolerated and mock projects are used as fallback.
func Load(config *Config) error {
	config.DemoMode = isDemoMode()

	l, f := logger.New(logger.NewLogger{})
	defer f.Close()

	viper.SetConfigName(cmdName)
	viper.SetConfigType("toml")

	viper.AddConfigPath(fmt.Sprintf("$HOME/.config/%s/", cmdName))
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if config.DemoMode {
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

	if config.DemoMode {
		config.Filters.Projects = mockProjects
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		l.Info(fmt.Sprintf("Config file changed: %s", e.Name))
	})
	viper.WatchConfig()

	// Env vars
	err = loadEnvVars(config)
	if err != nil && !config.DemoMode {
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

var flags = flag.NewFlagSet("mrglab", flag.ExitOnError)
var demoFlag = flags.Bool("demo", false, "use mocked data instead of calling api")

func isDemoMode() bool {
	if !flags.Parsed() {
		flags.Parse(os.Args[1:])
	}
	return *demoFlag
}

var mockProjects = []Project{
	{Name: "Payments API", ID: "10234567", FullPath: "acme-corp/payments-api"},
	{Name: "Web Dashboard", ID: "10234568", FullPath: "acme-corp/web-dashboard"},
	{Name: "Auth Service", ID: "10234569", FullPath: "acme-corp/auth-service"},
	{Name: "Mobile App", ID: "10234570", FullPath: "acme-corp/mobile-app"},
}
