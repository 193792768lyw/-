package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestOther(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(3)
	go worker(&wg)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total )

}

var total uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total, i)
	}
}

//
//var total struct {
//	sync.Mutex
//	value int
//}
//
//func worker(wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for i := 0; i <= 100; i++ {
//		//total.Lock()
//		total.value += i
//		//total.Unlock()
//	}
//}
////15150



type singleton struct {}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singleton {

	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}