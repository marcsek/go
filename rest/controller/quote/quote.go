package controller

import (
	"encoding/json"
	"net/http"
	"rest/model"
)

func GetQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	quote, err := model.GetQuote(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(quote)
}

type CharacterInput struct {
	Name string `json:"name"`
}

type CharacterOutput struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var input CharacterInput
	err := decoder.Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	output := CharacterOutput{Name: input.Name, Id: 1}

	json.NewEncoder(w).Encode(output)
}
