package bot

import "github.com/bwmarrin/discordgo"

func echoHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content != "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, m.Content)
	}
}
