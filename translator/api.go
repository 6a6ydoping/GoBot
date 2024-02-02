package translator

type TranslationRequest struct {
	Text           string `json:"text"`
	SourceLanguage string `json:"source_language"`
	TargetLanguage string `json:"target_language"`
}

type TranslationResponse struct {
	Status string `json:"status"`
	Data   struct {
		TranslatedText string `json:"translatedText"`
	} `json:"data"`
}
