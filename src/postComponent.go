package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// postComponent
type postComponent struct {
	vecty.Core
	disabled bool
	text     string
}

// Render renders the body
func (v *postComponent) Render() vecty.ComponentOrHTML {
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
					vecty.MarkupIf(
						v.disabled,
						vecty.Class(
							"disabled",
						),
					),
					prop.Href("/"),
					vecty.Style(
						"text-decoration",
						"none",
					),
				),
				vecty.Text("« Home"),
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
						prop.Src("/app/avatar.jpg"),
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
			elem.Div(
				vecty.Markup(
					vecty.Class(
						"col", "l3", "m2",
						"padding",
					),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class(
						"col", "l4", "m8",
						"markdown-body",
						"animate-opacity",
						"container",
						"padding-32",
					),
					vecty.UnsafeHTML(v.text),
				),
			),
		),

		vecty.If(v.text != "", v.loadDisqus(), footer()),
	)
}
