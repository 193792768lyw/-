package main

import (
	"fmt"
	"math/bits"
	"time"
)

func main() {

	t := time.Now()
	zero_tm := time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, t.Location()).Unix()
	fmt.Println(zero_tm)

}

func readBinaryWatch1(turnedOn int) (ans []string) {
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return
}

func readBinaryWatch(turnedOn int) (ans []string) {
	for i := 0; i < 1024; i++ {
		h, m := i>>6, i&63 // 用位运算取出高 4 位和低 6 位
		if h < 12 && m < 60 && bits.OnesCount(uint(i)) == turnedOn {
			ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
		}
	}
	return
}
