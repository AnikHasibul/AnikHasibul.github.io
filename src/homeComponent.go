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
					"center",
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
						"animate-opacity",
						"container",
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class(
							"panel",
							"white",
							"card-4",
							"text-grey",
							"padding-16",
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class(
								"center",
								"serif",
								"xlarge",
							),
						),

						vecty.Text("🎉"),
						elem.Break(),
						vecty.Text("Hurray!"),
					),
					vecty.Text("You are seeing a preview version of this blog! A regular version will come on 1st January 2019! 🎉🎉🎉"),
				),
				elem.Div(
					vecty.Markup(
						vecty.UnsafeHTML(v.text),
					),
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
