package bot

import (
	"fmt"
	"github.com/6a6ydoping/GoBot/config"
	"github.com/6a6ydoping/GoBot/weather"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	goBot *discordgo.Session
)

type Manager struct {
	WeatherManager *weather.Manager
}

func NewManager(manager *weather.Manager) *Manager {
	return &Manager{manager}
}

func (mg Manager) Start(config *config.BotConfig) {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	goBot.AddHandler(mg.messageCreate)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running")
}