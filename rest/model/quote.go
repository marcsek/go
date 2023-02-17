package model

import "context"

type Quote struct {
	Quote     string `json:"quote"`
	Character string `json:"character"`
}

const quoteURL = "https://thesimpsonsquoteapi.glitch.me/quotes"

func GetQuote(ctx context.Context) (Quote, error) {
	var quote []Quote

	res, err := fetch(quoteURL, ctx)

	if err != nil {
		return Quote{}, err
	}

	if res.Status == 200 {
		err := parseJson(res.Content, &quote)
		if err != nil {
			return quote[0], err
		}
	}

	return quote[0], nil
}
