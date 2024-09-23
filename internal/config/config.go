package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	BaseURL    string                   `mapstructure:"base_url"`
	APIToken   string                   `mapstructure:"token"`
	APIVersion string                   `mapstructure:"api_version"`
	Projects   []map[string]interface{} `mapstructure:"projects"`
}

var GlobalConfig Config

func Load(configObj *Config) {
	cmdName := "mrglab"

	viper.SetConfigName(cmdName)
	viper.SetConfigType("toml")

	viper.AddConfigPath(fmt.Sprintf("$HOME/.config/%s/", cmdName))
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal: %w", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Debug("config", func() {
			log.Println("Config file changed:", e.Name)
		})
	})
	viper.WatchConfig()

	// Env vars
	viper.SetEnvPrefix(cmdName)
	e := viper.BindEnv("token")
	if e != nil {
		logger.Debug("e", func() {
			log.Print(e)
		})
	}
	token := viper.Get("token")

	if token == nil {
		// TODO: report in statusline this error
		logger.Error(errors.New("api-token not set"))
		token = ""
	}
	GlobalConfig.APIToken = token.(string)
}
