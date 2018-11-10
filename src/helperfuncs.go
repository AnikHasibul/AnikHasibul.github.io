package main

import "strings"

// getTitle returns a title from the given page name
func getTitle(page string) string {
	return strings.Replace(
		page,
		"-",
		" ",
		-1,
	)
}
