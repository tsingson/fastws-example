package genflat

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginRequestT_Byte(t *testing.T) {
	as := assert.New(t)
	// serialized
	l := &LoginRequestT{
		MsgID:    1,
		Username: "1",
		Password: "1",
	}

	b := l.Byte()

	// un-serialized
	c := ByteLoginRequestT(b)
	if l.MsgID > 0 {
		fmt.Println(" id > ", c.MsgID, " u > ", c.Username, " pw > ", c.Password)
	}

	as.Equal(l.Password, c.Password)

}
