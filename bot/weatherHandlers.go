// Package bot provides functionality for creating and managing a Telegram bot using the
// discordGo library. It includes features such as command handling, message processing,
// and interaction with the external APIs.
package bot

import (
	"fmt"
	"github.com/6a6ydoping/GoBot/weather"
	"github.com/bwmarrin/discordgo"
	"os"
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
	// Extracting relevant weather details from the response
	weatherDetails := resp.Weather[0]
	temperature := resp.Main.Temp
	description := weatherDetails.Description
	iconFileName := weatherDetails.Icon + "@2x.png"

	// Open the file for reading
	file, err := os.Open("icons/" + iconFileName)
	if err != nil {
		fmt.Println("Error while opening icon file")
		return
	}
	defer file.Close()

	// Create a discord file to attach
	discordFile := &discordgo.File{
		Name:   iconFileName,
		Reader: file,
	}

	// Combining the message text
	messageText := fmt.Sprintf("Current weather in %s:\n\n", resp.Name)
	messageText += fmt.Sprintf("Temperature: %.2f°C\n", weather.KelvinToCelsius(temperature))
	messageText += fmt.Sprintf("Description: %s\n", description)

	// Sending message + image
	_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: messageText,
		Files:   []*discordgo.File{discordFile},
	})

	return
}
