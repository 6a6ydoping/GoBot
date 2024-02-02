package bot

import "github.com/bwmarrin/discordgo"

func (mg Manager) setReminder(s *discordgo.Session, m *discordgo.MessageCreate) {
	mg.ReminderManager.SetReminder(s, m)
}
