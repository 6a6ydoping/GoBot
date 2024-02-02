// Package bot provides functionality for creating and managing a Telegram bot using the
// discordGo library. It includes features such as command handling, message processing,
// and interaction with the external APIs.
package bot

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
)

// TODO: document translator related stuff
func (mg Manager) translate(s *discordgo.Session, m *discordgo.MessageCreate) {
	re := regexp.MustCompile(`!translate \[([^\]]+)\] \[([^\]]+)\] (.+)`)

	matches := re.FindStringSubmatch(m.Content)
	if len(matches) != 4 {
		s.ChannelMessageSend(m.ChannelID, "Invalid command format. Use !translate [from-lang] [to-lang] text")
		return
	}

	fromLang := matches[1]
	toLang := matches[2]
	text := matches[3]

	mg.TranslatorManager.Translate(s, m, text, fromLang, toLang)
}
