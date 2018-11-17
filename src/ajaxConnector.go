package main

import (
	"honnef.co/go/js/xhr"
)

// fetch fetches a page with get request
func fetch(endpoint string) (string, *respErr) {
	resp, err := fetchFromCache(endpoint)
	if err == nil {
		return resp, nil
	}
	return fetchFromServer(endpoint)
}

// fetchFromCache fetches an url from cache
func fetchFromCache(endpoint string) (string, *respErr) {
	resp, err := getFromCache(endpoint)
	if err != nil {
		return "", &respErr{
			Text: err.Error(),
		}
	}
	go fetchFromServer(endpoint)
	return resp, nil
}

// fetchFromServer fetches an url from the server
func fetchFromServer(endpoint string) (string, *respErr) {
	req := xhr.NewRequest(
		"GET",
		"https://loserdevapi.herokuapp.com/"+endpoint,
	)
	req.Timeout = 15000
	req.ResponseType = xhr.Text
	err := req.Send(nil)
	if err != nil {
		return "", &respErr{
			Status: req.Status,
			Text:   err.Error(),
		}
	}
	content := req.ResponseText
	go updateCache(
		endpoint,
		content,
	)
	return content, nil
}
