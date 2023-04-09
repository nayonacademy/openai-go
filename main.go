package main

import "github.com/nayonacademy/openai-go"

func main() {
	chat := openai.NewGPT()
	fmt.Println("Hello", chat)
}