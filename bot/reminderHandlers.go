// Package bot provides functionality for creating and managing a Telegram bot using the
// discordGo library. It includes features such as command handling, message processing,
// and interaction with the external APIs.
package bot

import "github.com/bwmarrin/discordgo"

func (mg Manager) setReminder(s *discordgo.Session, m *discordgo.MessageCreate) {
	mg.ReminderManager.SetReminder(s, m)
}
