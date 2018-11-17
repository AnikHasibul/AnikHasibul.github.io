package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func footer() *vecty.HTML {
	return elem.Footer(
		vecty.Markup(
			vecty.Class(
				"light-grey",
				"text-grey",
				"small",
				"center",
				"container",
				"padding",
				"padding-32",
			),
			vecty.Style("width", "100%"),
		),
		elem.Paragraph(
			vecty.Text(
				`This site is being maintained by Anik Hasibul.
For further contact: anikhasibul@outlook.com
				`,
			),
		),
		elem.Anchor(
			vecty.Markup(
				prop.Href(
					"https://github.com/AnikHasibul",
				),
				vecty.Attribute(
					"target",
					"_blank",
				),
				vecty.Class(
					"btn",
					"black",
					"round",
				),
				vecty.Style(
					"width",
					"40px",
				),
				vecty.Style(
					"height",
					"36px",
				),
			),
			vecty.Text("G"),
		),
		elem.Anchor(
			vecty.Markup(
				prop.Href(
					"https://fb.me/AnikHasibul.sh",
				),
				vecty.Attribute(
					"target",
					"_blank",
				),
				vecty.Class(
					"btn",
					"indigo",
					"round",
				),
				vecty.Style(
					"width",
					"40px",
				),
				vecty.Style(
					"height",
					"36px",
				),
				vecty.Style(
					"margin-left",
					"3px",
				),
			),
			vecty.Text("F"),
		),
		elem.Anchor(
			vecty.Markup(
				prop.Href(
					"https://twitter.com/iAmAnikHasibul",
				),
				vecty.Attribute(
					"target",
					"_blank",
				),
				vecty.Class(
					"btn",
					"round",
				),
				vecty.Style(
					"width",
					"40px",
				),
				vecty.Style(
					"background",
					"#4289f4",
				),
				vecty.Style(
					"color",
					"#fff",
				),
				vecty.Style(
					"height",
					"36px",
				),
				vecty.Style(
					"margin-left",
					"3px",
				),
			),
			vecty.Text("T"),
		),
		elem.Anchor(
			vecty.Markup(
				prop.Href(
					"https://medium.com/@AnikHasibul",
				),
				vecty.Attribute(
					"target",
					"_blank",
				),
				vecty.Class(
					"btn",
					"round",
					"black",
				),
				vecty.Style(
					"width",
					"40px",
				),
				vecty.Style(
					"height",
					"36px",
				),
				vecty.Style(
					"margin-left",
					"3px",
				),
			),
			vecty.Text("M"),
		),
	)
}
