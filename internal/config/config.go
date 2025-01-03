package config

import (
	"errors"
	"flag"
	"fmt"

	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Project struct {
	Name     string `mapstructure:"name"`
	FullPath string `mapstructure:"fullPath"`
	ID       string `mapstructure:"id"`
}

type Filter struct {
	Projects []Project `mapstructure:"projects"`
}
type Config struct {
	BaseURL  string `mapstructure:"base_url"`
	APIToken string `mapstructure:"token"`
	Filters  Filter `mapstructure:"filters"`
	DevMode  bool
}

var (
	GlobalConfig Config
	cmdName      = "mrglab"
)

func Load(config *Config) error {
	l, f := logger.New(logger.NewLogger{})
	defer f.Close()

	viper.SetConfigName(cmdName)
	viper.SetConfigType("toml")

	viper.AddConfigPath(fmt.Sprintf("$HOME/.config/%s/", cmdName))
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal: %w", err))
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://gitlab.com"
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		l.Info(fmt.Sprintf("Config file changed: %s", e.Name))
	})
	viper.WatchConfig()

	// Env vars
	err = loadEnvVars(config)
	if err != nil {
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
	config.DevMode = isDevMode()
	return nil
}

func isDevMode() bool {
	isDevMode := flag.Bool("dev", false, "activates dev mode to use mocked data instead of calling api")
	flag.Parse()

	return *isDevMode
}
