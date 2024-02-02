// Package config provides functionality for reading and initializing configuration settings
// for the bot and weather components.
package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config represents the overall configuration structure
type Config struct {
	Bot        BotConfig        `yaml:"bot"`
	Weather    WeatherConfig    `yaml:"weather"`
	Translator TranslatorConfig `yaml:"translator"`
	TriviaGame TriviaGameConfig `yaml:"triviaGame"`
}

// BotConfig represents the configuration settings for the bot.
type BotConfig struct {
	Token string `yaml:"token"`
}

// WeatherConfig represents the configuration settings for weather-related functionality.
type WeatherConfig struct {
	AppID            string `yaml:"AppID"`
	WeatherByCityURL string `yaml:"ByCityURL"`
}

type TranslatorConfig struct {
	BaseURL     string `yaml:"BaseURL"`
	ContentType string `yaml:"ContentType"`
	APIKey      string `yaml:"APIKey"`
	APIHost     string `yaml:"APIHost"`
}

type TriviaGameConfig struct {
	BaseURL   string `yaml:"BaseURL"`
	MinNumber int    `yaml:"MinNumber"`
	MaxNumber int    `yaml:"MaxNumber"`
	APIKey    string `yaml:"APIKey"`
	APIHost   string `yaml:"APIHost"`
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
