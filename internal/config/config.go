package config

import (
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
	viper.SetConfigName("glabt")
	viper.SetConfigType("toml")

	viper.AddConfigPath("$HOME/.config/glabt/")
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
	e := viper.BindEnv("glabt_token")
	if e != nil {
		logger.Debug("e", func() {
			log.Print(e)
		})
	}
	token := viper.Get("glabt_token")
	GlobalConfig.APIToken = token.(string)
}
