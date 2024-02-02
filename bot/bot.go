// Package bot provides functionality for creating and managing a Telegram bot using the
// discordGo library. It includes features such as command handling, message processing,
// and interaction with the external APIs.
package bot

import (
	"fmt"
	"github.com/6a6ydoping/GoBot/config"
	"github.com/6a6ydoping/GoBot/reminder"
	"github.com/6a6ydoping/GoBot/translator"
	"github.com/6a6ydoping/GoBot/weather"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	goBot *discordgo.Session
)

type Manager struct {
	WeatherManager    *weather.Manager
	ReminderManager   *reminder.Manager
	TranslatorManager *translator.Manager
}

func NewManager(wManager *weather.Manager, rManager *reminder.Manager, tManager *translator.Manager) *Manager {
	return &Manager{wManager, rManager, tManager}
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
