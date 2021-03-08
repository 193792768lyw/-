package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	fmt.Println("000000099999")
}


func init() {
 fmt.Println("00000000000000")
}

/*
原子操作配合互斥锁可以实现非常高效的单例模式。互斥锁的代价比普通整数的原子读写高很多，在性能敏感的地方可以增加一个数字型的标志位，
通过原子检测标志位状态降低互斥锁的使用次数来提高性能。
 */
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
func main() {

	for i := 0 ;i < 10 ; i ++ {
		go func() {
			fmt.Printf(	"%p\n",Instance())
		}()

	}
	time.Sleep(4 * time.Second)
}
//
//// 基于sync.Once重新实现单例模式
//var (
//	instance *singleton
//	once     sync.Once
//)
//
//func Instance() *singleton {
//	once.Do(func() {
//		instance = &singleton{}
//	})
//	return instance
//}