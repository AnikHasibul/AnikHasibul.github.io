package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func (v *postComponent) suggestionCOMP() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class(
				"panel",
			),
		),
	)
}
