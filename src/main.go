package main

import (
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

var (
	w     = dom.GetWindow()
	route = router.New()
)

func main() {
	route.HandleFunc(
		"/contact",
		contactPageHandler,
	)
	route.HandleFunc(
		"/posts/{page}",
		blogPostHandler,
	)
	route.HandleFunc(
		"/{pageNum}",
		homeHandler,
	)
	route.ShouldInterceptLinks = true
	route.Start()
}
