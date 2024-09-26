package config

import (
	"errors"
	"fmt"

	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	BaseURL    string                   `mapstructure:"base_url"`
	APIToken   string                   `mapstructure:"token"`
	APIVersion string                   `mapstructure:"api_version"`
	Projects   []map[string]interface{} `mapstructure:"projects"`
	DevMode    bool
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
		// TODO: report in statusline this error
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
	viper.SetEnvPrefix(cmdName)
	viper.BindEnv("dev")

	useMockedData := viper.Get("dev")
	if useMockedData != nil {
		return true
	}

	return false
}
