package types

const (
	MODEL_GPT_DAVINCI   string = "text-davinci-003"
	MODEL_CODEX_DAVINCI string = "code-davinci-002"
	COMPLETIONAPIURL string = "https://api.openai.com/v1/completions"
)

type Client struct{
	BearerToken string `json:"token"`
}

type Choice struct {
	Text string `json:"text"`
	Index int64 `json:"index"`
	Logprobs Bool `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type Usage struct{
	PromptTokens int64 `json:"prompt_token"`
	CompletionTokens int64 `json:"completion_token"`
	TotalTokens int64 `json:"total_token"`
}

type CompletionRequest struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	MaxToken int `json:"max_token"`
	Temperature float64 `json:"temperature"`
}

type CompletionResponse struct{
	ID string `json:"id"`
	Object string `json:"object"`
	Created int64 `json:"created"`
	Model string `json:"model"`
	Choices []Choice `json:"choices"`
	Use Usage `json:"use"`
}