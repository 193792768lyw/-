package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	fmt.Printf("下一个伪随机数总是%v。\n", rand.Uint32())
}
