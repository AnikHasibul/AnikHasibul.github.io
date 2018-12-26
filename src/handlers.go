package main

import (
	"fmt"
	"time"

	"github.com/go-humble/router"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/websocket/websocketjs"
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
		nav:      NewNav(),
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
	homeCOMP.nav.setCurrent(c.Params["pageNum"])
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
		disabled: true,
	}
	vecty.RenderBody(contactCOMP)
}

// chatHandler : the single page application handler
func chatHandler(c *router.Context) {
	vecty.SetTitle("Ask a Loser :: Bot :: Hasibul Hasan (Anik) | Personal blog | @AnikHasibul")
	chatCOMP := &chatComponent{
		disabled: true,
	}
	vecty.RenderBody(chatCOMP)

	go func() {
		var closesig = make(chan bool, 1)
		ws, err := websocketjs.New(
			"ws://localhost:5000/ask",
		)
		if err != nil {
			js.Global.Get("window").
				Call("alert", err.Error())
		}

		sendToWs := func(text string) {
			println(text)
			chatCOMP.chats = append(
				chatCOMP.chats,
				chatModel{
					sender:  "self",
					message: text,
				},
			)
			vecty.Rerender(chatCOMP)
			err := ws.Send(text)
			if err != nil {
				println(err.Error())
			}
		}
		go func() {
			for {
				select {
				case m := <-message:
					sendToWs(m)
				case <-closesig:
					return
				}
			}
		}()
		ws.AddEventListener(
			"open",
			false,
			func(e *js.Object) {
				sendToWs(".")
				chatCOMP.disabled = false
				vecty.Rerender(chatCOMP)
			})

		ws.AddEventListener(
			"message",
			false,
			func(e *js.Object) {
				chatCOMP.chats = append(
					chatCOMP.chats,
					chatModel{
						sender: "bot",
						message: e.Get("data").
							String(),
					},
				)
				vecty.Rerender(chatCOMP)
				// BUG: scroll to bottom not working
				//js.Global.Get("document").Call("getElementById", "chatview").Set("scrollTop", js.Global.Get("document").Call("getElementById", "chatview").Get("scrollHeight"))
			})

		ws.AddEventListener(
			"close",
			false,
			func(e *js.Object) {
				js.Global.Get("window").
					Call("alert", "closed")
				chatCOMP.disabled = true
				vecty.Rerender(chatCOMP)
				close(closesig)
			})
		ws.AddEventListener(
			"error",
			false,
			func(e *js.Object) {
				js.Global.Get("window").
					Call("alert", "error")
				chatCOMP.disabled = true
				vecty.Rerender(chatCOMP)
			})

		//	err = ws.Close()
	}()
}
