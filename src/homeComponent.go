package main

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// homeComponent
type homeComponent struct {
	vecty.Core
	nav      *nav
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
			vecty.Class("darker"),
		),
		elem.Header(
			vecty.Markup(
				vecty.Class(
					"container",
					"padding-16",
					"darker",
					"bar",
					"center",
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
					vecty.Attribute(
						"name",
						"Refresh homepage.",
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
					vecty.Class(
						"right",
					),
					vecty.Attribute(
						"name",
						"contact",
					),
					prop.Href("/contact"),
				),
				elem.Image(
					vecty.Markup(
						vecty.Class(
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
						vecty.Attribute("Alt", "Profile picture of author."),
						prop.Src("/app/avatar.jpg"),
					),
				),
			),
		),
		vecty.If(v.text != "", elem.Div(
			vecty.Markup(
				vecty.Class("row", "dark"),
			),
			elem.Div(vecty.Markup(
				vecty.Class(
					"col", "l2", "m2",
					"padding",
				),
			),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class(
						"col", "l8", "m8",
						"markdown-body",
						"animate-opacity",
						"container",
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.UnsafeHTML(v.text),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class(
							"padding-32",
						),
					),
					v.pagination(),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class(
						"col", "l2", "m2",
					),
				),
			),
		),
			footer()),
	)
}

func (v *homeComponent) pagination() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("bar"),
		),
		vecty.If(v.nav.prev() != 0,
			elem.Anchor(
				vecty.Markup(
					prop.Href(fmt.Sprintf("/%d", v.nav.prev())),
					vecty.Class(
						"left",
						"btn",
						"darker",
						"round",
					),
				),
				vecty.Text("« Previous"),
			),
		),
		vecty.If(v.nav.next() != 0,
			elem.Anchor(
				vecty.Markup(
					prop.Href(fmt.Sprintf("/%d", v.nav.next())),
					vecty.Class(
						"right",
						"btn",
						"darker",
						"round",
					),
				),
				vecty.Text("Next »"),
			),
		),
	)
}
