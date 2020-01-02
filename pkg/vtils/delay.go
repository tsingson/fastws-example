package vtils

import (
	"fmt"
	"math/rand"
	"time"
)

// Delay 随机延迟睡眠处理
func Delay(i, j int64) {
	c := DelayRand(i, j)
	// log.Info().Msg(vtils.StrBuilder("------------进入延迟时间  :", fmt.Sprintf("%d", c), "秒.................."))
	fmt.Println(StrBuilder("------------进入延迟时间  :", fmt.Sprintf("%d", c), "秒.................."))
	time.Sleep(time.Duration(c) * time.Second)
	// 	time.Sleep(5 * time.Second)
}

// DelayRand delay between seconds be set
func DelayRand(i, j int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c := r.Int63n(i) + j

	return c
}
