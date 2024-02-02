package triviaGame

import (
	"encoding/json"
	"fmt"
	"github.com/6a6ydoping/GoBot/config"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Manager struct {
	BaseURL     string
	MinNumber   int
	MaxNumber   int
	APIKey      string
	APIHost     string
	TriviaState map[string]int
}

func NewManager(config config.TriviaGameConfig) *Manager {
	return &Manager{BaseURL: config.BaseURL, MinNumber: config.MinNumber, MaxNumber: config.MaxNumber, APIKey: config.APIKey, APIHost: config.APIHost}
}

func (mg Manager) Trivia(s *discordgo.Session, m *discordgo.MessageCreate, correctAnswer int) {
	if mg.TriviaState == nil {
		mg.TriviaState = make(map[string]int)
	}

	if state, ok := mg.TriviaState[m.Author.ID]; ok {
		userResponse := strings.TrimSpace(m.Content)
		if userAnswer, err := strconv.Atoi(userResponse); err == nil {
			if userAnswer == state {
				s.ChannelMessageSend(m.ChannelID, "Correct! You got it right.")
			} else {
				s.ChannelMessageSend(m.ChannelID, "Incorrect. The correct answer is "+strconv.Itoa(state))
			}
			delete(mg.TriviaState, m.Author.ID)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Please respond with a valid number.")
		}
	} else {
		url := fmt.Sprintf(mg.BaseURL, getRandomNumberInRange(mg.MinNumber, mg.MaxNumber))

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("X-RapidAPI-Key", mg.APIKey)
		req.Header.Add("X-RapidAPI-Host", mg.APIHost)

		resp, _ := http.DefaultClient.Do(req)

		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Couldn't read the response body")
			return
		}

		var triviaResponse TriviaResponse
		err = json.Unmarshal([]byte(respBody), &triviaResponse)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		correctAnswer := triviaResponse.Number
		mg.TriviaState[m.Author.ID] = correctAnswer

		s.ChannelMessageSend(m.ChannelID, triviaResponse.Text)
	}
}
