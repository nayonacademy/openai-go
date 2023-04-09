package client

type ClientProperty struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	MaxToken int64 `json:"max_token"`
	Temperatur bool `json:"temperatur"`
}

func OpenaiReq(){
	client := &http.Client{}
	var data = strings.NewReader(`{
    "model": "text-davinci-003",
    "prompt": "Say this is a test",
    "max_tokens": 7,
    "temperature": 0
  }`)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer ")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}