// Package translator provides functionality for text translation using an external API.
package translator

// TranslationRequest is the base structure for constructing translation requests
type TranslationRequest struct {
	Text           string `json:"text"`
	SourceLanguage string `json:"source_language"`
	TargetLanguage string `json:"target_language"`
}

// TranslationResponse is the base structure for handling translation API responses
type TranslationResponse struct {
	Status string `json:"status"`
	Data   struct {
		TranslatedText string `json:"translatedText"`
	} `json:"data"`
}
