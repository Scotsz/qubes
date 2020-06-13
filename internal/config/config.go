package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Test string
	HTTP struct {
		Host string
		Port string
	}
	Redis struct {
		URL      string
		Username string
		Password string
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
	v := viper.New()
	v.SetConfigFile(file)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	v.AutomaticEnv()
	v.BindEnv("redis.url", "REDIS_URL")
	v.BindEnv("redis.username", "REDIS_USERNAME")
	v.BindEnv("redis.password", "REDIS_PASSWORD")
	v.BindEnv("http.host", "HTTP_HOST")
	v.BindEnv("http.port", "HTTP_PORT")

	var appConfig AppConfig
	err = v.Unmarshal(&appConfig)
	if err != nil {
		return nil, err
	}
	return &appConfig, err
}
