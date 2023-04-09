package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "strings"
	"encoding/json"
	"github.com/joho/godotenv"
)
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
	Logprobs bool `json:"logprobs"`
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
	MaxToken int `json:"max_tokens"`
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
type ClientProperty struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	MaxToken int64 `json:"max_tokens"`
	Temperatur int64 `json:"temperature"`
}
func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	cProperty := ClientProperty{
		Model: "text-davinci-003",
		Prompt: "Say this is a test",
		MaxToken: 7,
		Temperatur: 0,
	}
	data, err := json.Marshal(cProperty)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	client := &http.Client{}

	req, err := http.NewRequest("POST", COMPLETIONAPIURL, reader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result CompletionResponse
	if err := json.Unmarshal(bodyText, &result); err != nil {   // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result.Choices[0].Text)
}
