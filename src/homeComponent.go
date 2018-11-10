package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// homeComponent
type homeComponent struct {
	vecty.Core
	disabled bool
	text     string
}

// Render renders the body
func (v *homeComponent) Render() vecty.ComponentOrHTML {
	return elem.Body(l(v.disabled),
		vecty.Markup(
			vecty.Style(
				"overflow",
				"hidden",
			),
		),
		elem.Header(
			vecty.Markup(
				vecty.Class(
					"container",
					"padding-16",
					"white",
					"bar",
					"top",
					"border-blue",
				),
				vecty.Style(
					"border-bottom",
					"1px",
				),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class(
						"xlarge",
						"left",
					),
					prop.Href("/"),
					vecty.Style(
						"text-decoration",
						"none",
					),
				),
				vecty.Text("Home"),
			),
			elem.Anchor(
				vecty.Markup(
					prop.Href("/contact"),
				),
				elem.Image(
					vecty.Markup(
						vecty.Class(
							"xlarge",
							"right",
							"circle",
						),
						vecty.Style(
							"width",
							"32px",
						),
						vecty.Style(
							"height",
							"32px",
						),
						prop.Src("https://avatars0.githubusercontent.com/u/25927971?s=220&v=4"),
					),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("row"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class(
						"padding-32",
					),
				),
			),
			elem.Div(vecty.Markup(
				vecty.Class(
					"col", "l3", "m2",
					"padding",
				),
			),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class(
						"col", "l6", "m8",
						"markdown-body",
						"animate-left",
						"container",
						"padding-32",
					),
					vecty.UnsafeHTML(v.text),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class(
						"col", "l3", "m2",
					),
				),
			),
		),
		vecty.If(v.text != "", footer()),
	)
}
