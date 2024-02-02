package bot

import "github.com/bwmarrin/discordgo"

func (mg Manager) trivia(s *discordgo.Session, m *discordgo.MessageCreate, state int) {
	mg.TriviaGameManager.Trivia(s, m, state)
}
