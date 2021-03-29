package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("pppppppp")
		time.Sleep(4 * time.Second)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("===")
	wg.Wait()
	fmt.Println("main over")

}
