package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// l - The top loader bar
func l(val bool) *vecty.HTML {
	if !val {
		return nil
	}
	return elem.Div(
		vecty.Markup(
			vecty.Class("cyan"),
			vecty.Style("position", "fixed"),
			vecty.Style("z-index", "999"),
			vecty.Style("top", "0"),
			vecty.Style("left", "0"),
			vecty.Style("width", "100%"),
			vecty.Style("height", "2px"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class(
					"blue",
					"round-large",
					"animate-zoom-infinite",
				),
				vecty.Style(
					"width",
					"100%",
				),
				vecty.Style(
					"height",
					"2px",
				),
			),
		),
	)
}
