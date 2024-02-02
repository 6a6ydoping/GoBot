// Package translator provides functionality for text translation using an external API.
package translator

import (
	"encoding/json"
	"fmt"
	"github.com/6a6ydoping/GoBot/config"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Manager represents the main instance for interacting with the translation API.
type Manager struct {
	BaseURl string
}

func NewManager(config config.TranslatorConfig) *Manager {
	return &Manager{BaseURl: config.BaseURL}
}

// Translate translates the given text from the source language to the target language.
func (mg Manager) Translate(s *discordgo.Session, m *discordgo.MessageCreate, text, fromLang, toLang string) {
	reqBody := strings.NewReader(fmt.Sprintf("source_language=%s&target_language=%s&text=%s",
		url.QueryEscape(fromLang),
		url.QueryEscape(toLang),
		url.QueryEscape(text)))

	req, err := http.NewRequest("POST", mg.BaseURl, reqBody)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Couldn't create the request")
		return
	}
	// TODO: forgot to replace keys to the config
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-RapidAPI-Key", "8a0f63e73dmsh95db0ad37870d82p1c1e6ajsncc0ba05f46d4")
	req.Header.Set("X-RapidAPI-Host", "text-translator2.p.rapidapi.com")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Couldn't translate the text")
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Couldn't read the response body")
		return
	}

	var translationResponse TranslationResponse
	err = json.Unmarshal([]byte(respBody), &translationResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	decodedText, err := strconv.Unquote(`"` + translationResponse.Data.TranslatedText + `"`)
	if err != nil {
		fmt.Println("Error decoding Unicode escape sequences:", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, decodedText)
}
