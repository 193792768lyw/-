package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Jeffail/tunny"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sync"
	"testing"
	"time"
)

// 控制协程(goroutine)的并发数量

func TestOne1(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < math.MaxInt32; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}

/*
panic: too many concurrent operations on a single file or socket (max 1048575)
对单个 file/socket 的并发操作个数超过了系统上限，这个报错是 fmt.Printf 函数引起的，fmt.Printf 将格式化后的字符串打印到屏幕，即标准输出。
在 linux 系统中，标准输出也可以视为文件，内核(kernel)利用文件描述符(file descriptor)来访问文件，标准输出的文件描述符为 1，
错误输出文件描述符为 2，标准输入的文件描述符为 0。

简而言之，系统的资源被耗尽了。

那如果我们将 fmt.Printf 这行代码去掉呢？那程序很可能会因为内存不足而崩溃。这一点更好理解，每个协程至少需要消耗 2KB 的空间，
那么假设计算机的内存是 2GB，那么至多允许 2GB/2KB = 1M 个协程同时存在。
那如果协程中还存在着其他需要分配内存的操作，那么允许并发执行的协程将会数量级地减少。
*/

// 利用 channel 的缓存区限制并发的协程数量
func TestOne2(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}

// 利用第三方库 目前有很多第三方库实现了协程池，可以很方便地用来控制协程的并发数量
// https://github.com/Jeffail/tunny
// https://github.com/panjf2000/ants
func TestOne3(t *testing.T) {
	// 第一个参数是协程池的大小(poolSize)，第二个参数是协程运行的函数(worker)。
	pool := tunny.NewFunc(3, func(i interface{}) interface{} {
		log.Println(i)
		time.Sleep(time.Second)
		return nil
	})
	// 关闭协程池。
	defer pool.Close()
	// 将参数 i 传递给协程池定义好的 worker 处理。
	for i := 0; i < 10; i++ {
		go pool.Process(i)
	}
	time.Sleep(time.Second * 4)
}

// Go sync.Pool
/*
1 sync.Pool 的使用场景
一句话总结：保存和复用临时对象，减少内存分配，降低 GC 压力。
*/
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func TestOne4(t *testing.T) {

	stu := &Student{}
	json.Unmarshal(buf, stu)
}

func TestOne5(t *testing.T) {
	var studentPool = sync.Pool{
		New: func() interface{} {
			return new(Student)
		},
	}
	stu := studentPool.Get().(*Student)
	json.Unmarshal(buf, stu)
	studentPool.Put(stu)
	fmt.Println(studentPool.Get())
}

func TestOne6(t *testing.T) {
	content, _ := ioutil.ReadFile("test.jpeg")
	//	_ = ioutil.WriteFile("test.jpg.txt", content, 0666)

	encodedContent := base64.StdEncoding.EncodeToString(content)
	f, _ := os.Create("test.jpg.txt") //创建文件
	//	f, err1 := os.OpenFile("test.jpg.txt", os.O_APPEND, 0666) //打开文件
	//fmt.Println(f, err1)
	n, err := io.WriteString(f, encodedContent) //写入文件(字符串)
	fmt.Println(n, err)

}
