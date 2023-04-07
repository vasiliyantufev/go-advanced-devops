package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	ID    string   `json:"id"`
	MType string   `json:"mType"`
	Value *float64 `json:"value"`
	Hash  string   `json:"hash"`
}

func MakeHTTPCall(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func MakeHTTPWithBodyJSONCall(url string) (*http.Response, *Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	respBody := &Response{}
	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, nil, err
	}

	return resp, respBody, nil
}

func MakeHTTPWithBodyValueJSONCall(url string) (*http.Response, []byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, nil
}
