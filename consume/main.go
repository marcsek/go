package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Quote struct {
	Quote     string `json:"quote"`
	Character string `json:"character"`
	Image     string `json:"image"`
	CharDir   string `json:"characterDirection"`
}

func main() {
	response := fetch("https://thesimpsonsquoteapi.glitch.me/quotes")

	var obj []Quote
	parseJson(response.Content, &obj)

	fmt.Printf("Random Simpsons quote 💅\n📕 Quote: %v\n👨 By: %v", obj[0].Quote, obj[0].Character)
}

type FetchOutput struct {
	Status  int
	Length  int64
	Content []byte
}

func fetch(url string) FetchOutput {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	return FetchOutput{Status: res.StatusCode, Length: res.ContentLength, Content: content}
}

func parseJson[Shape any](bytes []byte, shape Shape) {
	isValid := json.Valid(bytes)
	if isValid {
		err := json.Unmarshal(bytes, shape)

		if err != nil {
			panic(err)
		}
	}
}
