package main

import (
	"github.com/go-humble/router"
	"github.com/gopherjs/vecty"
	"honnef.co/go/js/dom"
)

var (
	w     = dom.GetWindow()
	route = router.New()
)

func main() {
	vecty.AddStylesheet("/app/style.css")
	route.HandleFunc(
		"/contact",
		contactPageHandler,
	)
	route.HandleFunc(
		"/posts/{page}",
		blogPostHandler,
	)
	route.HandleFunc(
		"/ask",
		chatHandler,
	)
	route.HandleFunc(
		"/{pageNum}",
		homeHandler,
	)
	route.ShouldInterceptLinks = true
	route.Start()
}
