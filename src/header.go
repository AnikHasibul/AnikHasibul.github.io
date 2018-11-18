package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func logo() *vecty.HTML {
	return elem.Bold(
		vecty.Markup(
			vecty.Class(
				"hide-small",
				"xlarge",
				"center",
			),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class(
					"text-red",
				),
			),
			vecty.Text("Loser"),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class(
					"text-black",
				),
			),
			vecty.Text("Dev"),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class(
					"text-grey",
				),
			),
			vecty.Text("."),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class(
					"text-orange",
				),
			),
			vecty.Text("me"),
		),
	)

}
