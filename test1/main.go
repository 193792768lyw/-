package main

import (
	"fmt"
	"time"
)

var qq int

func main() {
	go func() {
		fmt.Println(qq)
	}()

	go func() {
		qq++
	}()

	time.Sleep(3*time.Second)
}
