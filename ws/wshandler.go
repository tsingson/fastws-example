package ws

import (
	"fmt"
	"os"

	"github.com/tsingson/fastws"

	"github.com/tsingson/fastws-example/apis/xone/genflat"
)

func WsHandler(conn *fastws.Conn) {
	fmt.Printf("Opened connection\n")

	l := &genflat.LoginRequestT{
		MsgID:    1,
		Username: "1",
		Password: "1",
	}

	_, _ = conn.Write(l.Marshal())

	var msg []byte
	var mod fastws.Mode
	var err error
	for {
		mod, msg, err = conn.ReadMessage(msg[:0])
		if err != nil {
			if err != fastws.EOF {
				fmt.Fprintf(os.Stderr, "error reading message: %s\n", err)
			}
			break
		}

		l := genflat.UnmarshalLoginRequestT(msg)
		if l.MsgID > 0 {
			fmt.Println("mod > ", mod, " id > ", l.MsgID, " u > ", l.Username, " pw > ", l.Password)
		}

		_, err = conn.Write(msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing message: %s\n", err)
			break
		}
	}

	fmt.Printf("Closed connection\n")
}
