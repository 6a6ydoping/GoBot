package main

import (
	"fmt"
	"github.com/6a6ydoping/GoBot/bot"
	"github.com/6a6ydoping/GoBot/config"
	"github.com/6a6ydoping/GoBot/reminder"
	"github.com/6a6ydoping/GoBot/weather"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		fmt.Println("Error reading config")
		return
	}
	fmt.Println(cfg.Weather.WeatherByCityURL)
	botManager := bot.NewManager(weather.NewManager(cfg.Weather), reminder.NewManager())

	botManager.Start(&cfg.Bot)
	<-make(chan struct{})
	return
}
