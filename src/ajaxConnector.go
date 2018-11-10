package main

import (
	"honnef.co/go/js/xhr"
)

// fetch fetches a page with get request
func fetch(endpoint string) (string, *respErr) {
	defer func() {
	}()
	req := xhr.NewRequest(
		"GET",
		endpoint,
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
	return req.ResponseText, nil
}
