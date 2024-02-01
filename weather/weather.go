package weather

import (
	"encoding/json"
	"fmt"
	"github.com/6a6ydoping/GoBot/config"
	"net/http"
	"net/url"
)

type Manager struct {
	AppID            string
	WeatherByCityURL string
}

func NewManager(config config.WeatherConfig) *Manager {
	return &Manager{config.AppID, config.WeatherByCityURL}
}

func (m Manager) WeatherByCity(city string) (WeatherResponse, error) {
	// Create object to return
	var response WeatherResponse

	// parsing url from string
	parsedURL, err := url.Parse(m.WeatherByCityURL)
	if err != nil {
		fmt.Println("Error parsing url:", err)
		return response, err
	}
	// adding params, make it string
	qParams := parsedURL.Query()
	qParams.Add("appid", m.AppID)
	qParams.Add("q", city) // q - requested city name
	parsedURL.RawQuery = qParams.Encode()
	url := parsedURL.String()

	// request to weather api
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return response, err
	}
	defer resp.Body.Close()
	// only if 200 status
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: API returned status code %d\n", resp.StatusCode)
		return response, err
	}
	// trying to encode
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return response, err
	}
	return response, nil
}
