package reminder

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"strconv"
	"strings"
	"time"
)

type Manager struct {
	UserReminders map[string][]Reminder
	cronScheduler *cron.Cron
}

func NewManager() *Manager {
	return &Manager{UserReminders: make(map[string][]Reminder, 100), cronScheduler: cron.New()} //TODO: replace to config map cap
}

type Reminder struct {
	Message   string
	Scheduled time.Time
}

func (mg Manager) SetReminder(s *discordgo.Session, m *discordgo.MessageCreate) {
	commandParts := strings.Fields(m.Content)[1:]
	if len(commandParts) < 4 {
		s.ChannelMessageSend(m.ChannelID, "Invalid command format. Use !reminder date=<date> time=<time> gmt=<gmt> text=<message>")
		return
	}

	var date, timeStr, gmt, messageText string
	values := GetAllTextsInQuotes(m.Content)
	date = values[0]
	timeStr = values[1]
	gmt = values[2]
	messageText = values[3]
	
	if date == "" || timeStr == "" || gmt == "" || messageText == "" {
		s.ChannelMessageSend(m.ChannelID, "Invalid command format. Use !reminder date=<date> time=<time> gmt=<gmt> text=<message>")
		return
	}

	scheduledTime, err := time.Parse("02.01.2006 15:04", fmt.Sprintf("%s %s", date, timeStr))
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Invalid date or time format.")
		return
	}

	gmtInt, err := strconv.Atoi(gmt)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Invalid date or time format.")
		return
	}

	location := time.FixedZone("UTC", gmtInt*60*60)

	reminder := Reminder{
		Message:   messageText,
		Scheduled: time.Date(scheduledTime.Year(), scheduledTime.Month(), scheduledTime.Day(), scheduledTime.Hour(), scheduledTime.Minute(), scheduledTime.Second(), 0, location),
	}

	go scheduleReminder(s, m, reminder, gmtInt)
	s.ChannelMessageSend(m.ChannelID, "Got it! I will remind you at "+scheduledTime.Format("15:04"))
}
