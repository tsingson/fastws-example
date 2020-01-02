package main

import (
	"bytes"
	"log"
	"time"

	"github.com/tsingson/fastws"
	"github.com/valyala/fasthttp"

	"github.com/tsingson/fastws-example/ws"
)

func main() {
	// Configure websocket upgrader.
	/**
	upgr := fastws.Upgrader{
		UpgradeHandler: checkCookies,
		Handler:        websocketHandler,
	}

	// Configure router handler.
	r := router.New()
	r.GET("/set", setCookieHandler)
	r.GET("/ws", upgr.Upgrade)

	wsserver := fasthttp.Server{
		Handler: r.Handler,
	}
	go wsserver.ListenAndServe(":8080")


	*/

	_ = ws.StartClient("ws://localhost:8080/ws", "http://localhost:8080/set")
	select {}
}

func websocketHandler(c *fastws.Conn) {
	c.WriteString("Hello world")
	_, msg, err := c.ReadMessage(nil)
	if err != nil {
		panic(err)
	}
	log.Printf("Readed %s\n", msg)
	c.Close()
}

var (
	cookieKey   = []byte("cookiekey")
	cookieValue = []byte("thisisavalidcookievalue")
)

func checkCookies(ctx *fasthttp.RequestCtx) bool {
	cookie := ctx.Request.Header.CookieBytes(cookieKey)
	if bytes.Equal(cookie, cookieValue) {
		return true
	}
	ctx.Error("You don't have a cookie D:", fasthttp.StatusBadRequest)
	return false
}

func setCookieHandler(ctx *fasthttp.RequestCtx) {
	setCookieWithTimeout(ctx, time.Time{})
}

func delCookieHandler(ctx *fasthttp.RequestCtx) {
	setCookieWithTimeout(ctx, time.Now())
}

func setCookieWithTimeout(ctx *fasthttp.RequestCtx, t time.Time) {
	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)

	cookie.SetKeyBytes(cookieKey)
	cookie.SetValueBytes(cookieValue)

	if !t.IsZero() {
		cookie.SetExpire(t)
	}

	ctx.Response.Header.SetCookie(cookie)
}
