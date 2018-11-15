package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func (v *postComponent) loadDisqus() *vecty.HTML {
	if v.text == "" {
		return nil
	}
	return disqus()
}

func disqus() *vecty.HTML {
	fireEvent()
	return elem.Div(
		elem.Div(
			vecty.Markup(
				vecty.Class(
					"panel",
					"padding",
					"container",
					"padding-16",
				),
			),
			elem.Div(
				vecty.Markup(
					prop.ID("disqus_thread"),
				),
			),
		),
		elem.Script(
			vecty.Markup(
				prop.Src(
					"https://loserdev.disqus.com/embed.js",
				),
				vecty.Attribute("data-timestamp", fmt.Sprintf("%d", time.Now().Unix())),
			),
		),
	)

}

func fireEvent() {
	js.Global.Set("disqus_config", func(this *js.Object) {
		this.Get("page").Set("url", w.Location().Href)
		this.Get("page").Set("identifier ", w.Location().Href)
	})
}
