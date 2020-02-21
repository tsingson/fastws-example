package ws

import (
	"fmt"
	"log"

	"github.com/tsingson/fastws"
	"github.com/valyala/fasthttp"

	"github.com/tsingson/fastws-example/apis/xone/genflat"
)

func StartClient(urlws, urlset string) error {
	c, err := fastws.Dial(urlws)
	if err != nil {
		return err
	}

	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	// cookie := fasthttp.AcquireCookie()
	// defer fasthttp.ReleaseCookie(cookie)
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	/**
	req.SetRequestURI(urlset)

	err = fasthttp.Do(req, res)
	checkErr(err)

	cookie.SetKeyBytes(cookieKey)
	if !res.Header.Cookie(cookie) {
		panic("cookie not found in response")
	}
	req.Reset()
	req.Header.SetCookieBytesKV(cookie.Key(), cookie.Value())
	*/
	// c, err = fastws.DialWithHeaders(urlws, req)
	// checkErr(err)
	// defer c.Close()

	log.Println("Connected")
	if c != nil {
		var count int32
		for {

			mod, msg, er1 := c.ReadMessage(nil)
			if er1 != nil {
				continue
			}

			l := genflat.UnmarshalLoginRequestT(msg)
			if l.MsgID > 0 {
				fmt.Println("mod > ", mod, " id > ", l.MsgID, " u > ", l.Username, " pw > ", l.Password)
			}

			fmt.Println(string(msg))

			x := &genflat.LoginRequestT{
				MsgID:    count,
				Username: "1",
				Password: "1",
			}

			count++

			_, _ = c.Write(x.Marshal())
		}
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
