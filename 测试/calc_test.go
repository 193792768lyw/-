package main

import (
	"fmt"
	"sync/atomic"
	"testing"
)


func TestAdd(t *testing.T) {
	//if ans := Add(1, 2); ans != 3 {
	//	t.Errorf("1 + 2 expected be 3, but %d got", ans)
	//}
	//
	//if ans := Add(-10, -20); ans != -30 {
	//	t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	//}


		//生成10个并发序列号
		for i := 0; i < 10; i++ {
			go GenID()
		}

		fmt.Println(GenID())

}

var (
	// 序列号
	seq int64
)

// 序列号生成器
func GenID() int64 {

	// 尝试原子的增加序列号
	atomic.AddInt64(&seq, 1)
	return seq
}

