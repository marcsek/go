package model

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type FetchOutput struct {
	Status  int
	Length  int64
	Content []byte
}

type chanRes struct {
	value *http.Response
	err   error
}

func fetch(url string, ctx context.Context) (FetchOutput, error) {
	ctxwt, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	respch := make(chan chanRes)

	go func() {
		res, err := http.Get(url)
		respch <- chanRes{
			value: res,
			err:   err,
		}
		defer res.Body.Close()
	}()

	select {
	case <-ctxwt.Done():
		return FetchOutput{}, errors.New("Request took too long")
	case resp := <-respch:
		res := resp.value
		content, _ := io.ReadAll(res.Body)
		return FetchOutput{Status: res.StatusCode, Length: res.ContentLength, Content: content}, nil
	}
}

func parseJson[Shape any](bytes []byte, shape Shape) error {
	isValid := json.Valid(bytes)

	if !isValid {
		return errors.New("Invalid JSON.")
	}

	return json.Unmarshal(bytes, shape)
}
