package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Bot     BotConfig     `yaml:"bot"`
	Weather WeatherConfig `yaml:"weather"`
}

type BotConfig struct {
	Token string `yaml:"token"`
}

type WeatherConfig struct {
	AppID            string `yaml:"AppID"`
	WeatherByCityURL string `yaml:"ByCityURL"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cfg, nil
}
