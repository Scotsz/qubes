package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	HTTP struct {
		Host string
		Port string
	}
}

func NewAppConfig(file string) (*AppConfig, error) {
	config, err := loadConfig(file)
	if err != nil {
		return nil, err
	}
	return config, nil
}
func loadConfig(file string) (*AppConfig, error) {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, err
	}
	return &appConfig, err
}
