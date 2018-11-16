package main

import "github.com/go-humble/locstor"

// addToCache adds a page to cache
func addToCache(url, content string) error {
	return locstor.SetItem(url, content)
}

// getFromCache returns a page from csche
func getFromCache(url string) (string, error) {
	return locstor.GetItem(url)
}

// updateCache updates a page from cache
func updateCache(url, content string) error {
	return locstor.SetItem(url, content)
}
