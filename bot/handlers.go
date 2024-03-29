// Package bot provides functionality for creating and managing a Telegram bot using the
// discordGo library. It includes features such as command handling, message processing,
// and interaction with the external APIs.
package bot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// messageCreate entry point in the bot, for processing incoming messages
func (mg Manager) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if state, ok := mg.TriviaGameManager.TriviaState[m.Author.ID]; ok {
		mg.trivia(s, m, state)
		return
	}

	// commands
	command := strings.Fields(m.Content)[0]
	if command[0] != '!' { // Check if this not command
		return
	}
	switch command {
	case "!weather":
		mg.getWeatherByCity(s, m)
	case "!reminder":
		mg.setReminder(s, m)
	case "!translate":
		mg.translate(s, m)
	case "!trivia":
		mg.trivia(s, m, 0)
	case "!help":
		mg.help(s, m)
	}
}
