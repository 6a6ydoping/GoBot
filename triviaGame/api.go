package triviaGame

type TriviaResponse struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
	Found  bool   `json:"found"`
	Type   string `json:"type"`
}
