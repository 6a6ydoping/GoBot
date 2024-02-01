package main

import (
	"fmt"
	"github.com/6a6ydoping/GoBot/bot"
	"github.com/6a6ydoping/GoBot/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		fmt.Println("Error reading config")
		return
	}

	bot.Start(&cfg.Bot)
	<-make(chan struct{})
	return
}
