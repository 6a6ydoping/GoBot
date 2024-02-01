package bot

import (
	"fmt"
	"github.com/6a6ydoping/GoBot/weather"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (mg Manager) getWeatherByCity(s *discordgo.Session, m *discordgo.MessageCreate) {
	city := strings.Fields(m.Content)[1]
	if city == "" {
		s.ChannelMessageSend(m.ChannelID, "Please write correct city")
	}
	resp, err := mg.WeatherManager.WeatherByCity(city)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Couldn't find info about this city")
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%f", weather.KelvinToCelsius(resp.Main.Temp)))
	return
}
