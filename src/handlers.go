package main

import (
	"fmt"
	"time"

	"github.com/go-humble/router"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

// blogPostHandler : the single page application handler
func blogPostHandler(c *router.Context) {
	postCOMP := &postComponent{
		disabled: true,
	}
	vecty.RenderBody(postCOMP)
	docRoot := "raw.githubusercontent.com/anikhasibul/anikhasibul.github.io/published/posts/"
	postCOMP.disabled = true
	vecty.Rerender(postCOMP)
	defer func() {
		postCOMP.disabled = false
		vecty.Rerender(postCOMP)
	}()
	resp, err := fetch(
		docRoot + c.Params["page"] + ".md",
	)
	if err != nil {
		resp = err.Error()
	}
	postCOMP.text = string(resp)
	vecty.SetTitle(getTitle(c.Params["page"]) + " :: Hasibul Hasan (Anik) | @AnikHasibul || anikhasibul.github.io")
	vecty.Rerender(postCOMP)
	js.Global.Get("window").
		Call("scrollTo", 0, 0)
}

// homeHandler : the single page application handler
func homeHandler(c *router.Context) {
	vecty.SetTitle("Home :: Hasibul Hasan (Anik) | Personal blog | @AnikHasibul")
	homeCOMP := &homeComponent{
		disabled: true,
	}
	vecty.RenderBody(homeCOMP)
	if c.Params["pageNum"] == "" {
		year, week := time.Now().ISOWeek()
		c.Params["pageNum"] = fmt.Sprintf(
			"%d%d",
			year,
			week,
		)
	}
	docRoot := "raw.githubusercontent.com/anikhasibul/anikhasibul.github.io/published/weekly/"
	homeCOMP.disabled = true
	vecty.Rerender(homeCOMP)
	defer func() {
		homeCOMP.disabled = false
		vecty.Rerender(homeCOMP)
	}()
	resp, err := fetch(
		docRoot + c.Params["pageNum"] + ".md",
	)
	if err != nil {
		resp = err.Error()
	}
	homeCOMP.text = string(resp)
	vecty.RenderBody(homeCOMP)
	js.Global.Get("window").
		Call("scrollTo", 0, 0)
}

// contactPageHandler : the single page application handler
func contactPageHandler(c *router.Context) {
	vecty.SetTitle("Contact :: Hasibul Hasan (Anik) | Personal blog | @AnikHasibul")
	contactCOMP := &contactComponent{
		disabled: false,
	}
	vecty.RenderBody(contactCOMP)
}
