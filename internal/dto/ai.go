package dto

type GenerateRequest struct {
	SessionID        string `json:"session_id"`
	Prompt           string `json:"prompt"`
	AdditionalPrompt string `json:"additional_prompt"`
}

type GenerateResponse struct {
	Response  string   `json:"response"`
	FoodNames []string `json:"food_names"`
}
