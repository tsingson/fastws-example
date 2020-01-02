package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/fasthttp/router"
	"github.com/tsingson/fastws"
	"github.com/valyala/fasthttp"

	"github.com/tsingson/fastws-example/pkg/vtils"
	"github.com/tsingson/fastws-example/ws"
)

func main() {
	r := router.New()

	path, _ := vtils.GetCurrentExecDir()
	path = path + "/www"
	r.ServeFiles("/www/*filepath", path)
	r.GET("/ws", fastws.Upgrade(ws.WsHandler))

	server := fasthttp.Server{
		Handler: r.Handler,
	}
	go server.ListenAndServe(":8080")

	fmt.Println("Visit http://localhost:8080")

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh
	signal.Stop(sigCh)
	signal.Reset(os.Interrupt)
	server.Shutdown()
}
