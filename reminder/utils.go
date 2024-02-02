package reminder

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
	"time"
)

func GetAllTextsInQuotes(input string) []string {
	var texts []string

	// Define a regular expression to find text within quotes
	re := regexp.MustCompile(`"([^"]+)"`)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(input, -1)

	// Extract the texts from the matches
	for _, match := range matches {
		texts = append(texts, match[1])
	}

	return texts
}

func scheduleReminder(s *discordgo.Session, m *discordgo.MessageCreate, reminder Reminder, gmt int) {
	// Calculate the duration until the scheduled time

	duration := reminder.Scheduled.Sub(time.Now().In(time.FixedZone("UTC", gmt*60*60)))

	// Create a timer
	timer := time.NewTimer(duration)

	// Wait for the timer to expire
	<-timer.C

	// Execute the reminder task
	s.ChannelMessageSend(m.ChannelID, reminder.Message)
}
