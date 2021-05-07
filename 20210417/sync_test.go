package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestA1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const N = 5
	var values [N]int32

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		i := i
		go func() {
			values[i] = 50 + rand.Int31n(50)
			fmt.Println("Done:", i)
			wg.Done() // <=> wg.Add(-1)
		}()
	}

	wg.Wait()
	// 所有的元素都保证被初始化了。
	fmt.Println("values:", values)

}

/*
一个*sync.WaitGroup值的Wait方法可以在多个协程中调用。 当对应的sync.WaitGroup值维护的计数降为0，
这些协程都将得到一个（广播）通知而结束阻塞状态。

一个WaitGroup可以在它的一个Wait方法返回之后被重用。 但是请注意，当一个WaitGroup值维护的基数为零时，
它的带有正整数实参的Add方法调用不能和它的Wait方法调用并发运行，否则将可能出现数据竞争。
这一点可以查看add方法的注释
*/
func TestA2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const N = 5
	var values [N]int32

	var wgA, wgB sync.WaitGroup
	wgA.Add(N)
	wgB.Add(1)

	for i := 0; i < N; i++ {
		i := i
		go func() {
			wgB.Wait() // 等待广播通知
			log.Printf("values[%v]=%v \n", i, values[i])
			wgA.Done()
		}()
	}

	// 下面这个循环保证将在上面的任何一个
	// wg.Wait调用结束之前执行。
	for i := 0; i < N; i++ {
		values[i] = 50 + rand.Int31n(50)
	}
	wgB.Done() // 发出一个广播通知
	wgA.Wait()
}

/*
sync.Once类型
每个*sync.Once值有一个Do(f func())方法。 此方法只有一个类型为func()的参数。

对一个可寻址的sync.Once值o，o.Do()（即(&o).Do()的简写形式）方法调用可以在多个协程中被多次并发地执行，
这些方法调用的实参应该（但并不强制）为同一个函数值。 在这些方法调用中，有且只有一个调用的实参函数（值）将得到调用。
此被调用的实参函数保证在任何o.Do()方法调用返回之前退出。 换句话说，被调用的实参函数内的代码将在任何o.Do()方法返回调用之前被执行。

一般来说，一个sync.Once值被用来确保一段代码在一个并发程序中被执行且仅被执行一次。

一个例子：
在此例中，Hello将仅被输出一次，而world!将被输出5次，并且Hello肯定在所有的5个world!之前输出。
*/

func TestA3(t *testing.T) {
	log.SetFlags(0)

	x := 0
	doSomething := func() {
		x++
		log.Println("Hello")
	}

	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(doSomething)
			log.Println("world!")
		}()
	}

	wg.Wait()
	log.Println("x =", x) // x = 1
}

/*
sync.Mutex（互斥锁）和sync.RWMutex（读写锁）类型
*sync.Mutex和*sync.RWMutex类型都实现了sync.Locker接口类型。 所以这两个类型都有两个方法：Lock()和Unlock()，
用来保护一份数据不会被多个使用者同时读取和修改。

除了Lock()和Unlock()这两个方法，*sync.RWMutex类型还有两个另外的方法：RLock()和RUnlock()，用来支持多个读取者并发读取一份
数据但防止此份数据被某个数据写入者和其它数据访问者（包括读取者和写入者）同时使用。

（注意：这里的数据读取者和数据写入者不应该从字面上理解。有时候某些数据读取者可能修改数据，而有些数据写入者可能只读取数据。）

一个Mutex值常称为一个互斥锁。 一个Mutex零值为一个尚未加锁的互斥锁。 一个（可寻址的）Mutex值m只有在未加锁状态时才能通过
m.Lock()方法调用被成功加锁。 换句话说，一旦m值被加了锁（亦即某个m.Lock()方法调用成功返回）， 一个新的加锁试图将导致当前协程进入阻塞状态，
直到此Mutex值被解锁为止（通过m.Unlock()方法调用）。

注意：m.Lock()和m.Unlock()分别是(&m).Lock()和(&m).Unlock()的简写形式。

一个使用sync.Mutex的例子：
*/

type Counter struct {
	m sync.Mutex
	n uint64
}

func (c *Counter) Value() uint64 {
	c.m.Lock()
	defer c.m.Unlock()
	return c.n
}

func (c *Counter) Increase(delta uint64) {
	c.m.Lock()
	c.n += delta
	c.m.Unlock()
}

/*
在上面这个例子中，一个Counter值使用了一个Mutex字段来确保它的字段n永远不会被多个协程同时使用。
*/
func TestA4(t *testing.T) {
	var c Counter
	for i := 0; i < 100; i++ {
		go func() {
			for k := 0; k < 100; k++ {
				c.Increase(1)
			}
		}()
	}

	// 此循环仅为演示目的。
	for c.Value() < 10000 {
		runtime.Gosched()
	}
	fmt.Println(c.Value()) // 10000
}

/*
一个RWMutex值常称为一个读写互斥锁，它的内部包含两个锁：一个写锁和一个读锁。 对于一个可寻址的RWMutex值rwm，数据写入者可以通过方法调用rwm.Lock()对rwm加写锁，
或者通过rwm.RLock()方法调用对rwm加读锁。 方法调用rwm.Unlock()和rwm.RUnlock()用来解开rwm的写锁和读锁。 rwm的读锁维护着一个计数。
当rwm.RLock()调用成功时，此计数增1；当rwm.Unlock()调用成功时，此计数减1； 一个零计数表示rwm的读锁处于未加锁状态；
反之，一个非零计数（肯定大于零）表示rwm的读锁处于加锁状态。

注意rwm.Lock()、rwm.Unlock()、rwm.RLock()和rwm.RUnlock()分别是(&rwm).Lock()、(&rwm).Unlock()、(&rwm).RLock()和(&rwm).RUnlock()的简写形式。

对于一个可寻址的RWMutex值rwm，下列规则存在：
rwm的写锁只有在它的写锁和读锁都处于未加锁状态时才能被成功加锁。 换句话说，rwm的写锁在任何时刻最多只能被一个数据写入者成功加锁，并且rwm的写锁和读锁不能同时处于加锁状态。
当rwm的写锁正处于加锁状态的时候，任何新的对之加写锁或者加读锁的操作试图都将导致当前协程进入阻塞状态，直到此写锁被解锁，这样的操作试图才有机会成功。
当rwm的读锁正处于加锁状态的时候，新的加写锁的操作试图将导致当前协程进入阻塞状态。 但是，一个新的加读锁的操作试图将成功，只要此操作试图发生在任何被阻塞的加写锁的操作试图之前（见下一条规则）。 换句话说，一个读写互斥锁的读锁可以同时被多个数据读取者同时加锁而持有。 当rwm的读锁维护的计数清零时，读锁将返回未加锁状态。
假设rwm的读锁正处于加锁状态的时候，为了防止后续数据写入者没有机会成功加写锁，后续发生在某个被阻塞的加写锁操作试图之后的所有加读锁的试图都将被阻塞。
假设rwm的写锁正处于加锁状态的时候，（至少对于标准编译器来说，）为了防止后续数据读取者没有机会成功加读锁，发生在此写锁下一次被解锁之前的所有加读锁的试图都将在此写锁下一次被解锁之后肯定取得成功，即使所有这些加读锁的试图发生在一些仍被阻塞的加写锁的试图之后。
后两条规则是为了确保数据读取者和写入者都有机会执行它们的操作。

请注意：一个锁并不会绑定到一个协程上，即一个锁并不记录哪个协程成功地加锁了它。 换句话说，一个锁的加锁者和此锁的解锁者可以不是同一个协程，尽管在实践中这种情况并不多见。
*/

// 根据上面列出的后两条规则，下面这个程序最有可能输出abdc。

func TestA5(t *testing.T) {
	var m sync.RWMutex
	go func() {
		m.RLock()
		fmt.Print("a")
		time.Sleep(time.Second)
		m.RUnlock()
	}()
	go func() {
		time.Sleep(time.Second * 1 / 4)
		m.Lock()
		fmt.Print("b")
		time.Sleep(time.Second)
		m.Unlock()
	}()
	go func() {
		time.Sleep(time.Second * 2 / 4)
		m.Lock()
		fmt.Print("c")
		m.Unlock()
	}()
	go func() {
		time.Sleep(time.Second * 3 / 4)
		m.RLock()
		fmt.Print("d")
		m.RUnlock()
	}()
	time.Sleep(time.Second * 3)
	fmt.Println()
}

/*
sync.Mutex和sync.RWMutex值也可以用来实现通知，尽管这不是Go中最优雅的方法来实现通知。 下面是一个使用了Mutex值来实现通知的例子。
*/
func TestA6(t *testing.T) {
	var m sync.Mutex
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Hi")
		m.Unlock() // 发出一个通知
	}()
	m.Lock() // 等待通知
	fmt.Println("Bye")
}

func TestA21(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const N = 10
	var values [N]string

	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d) // 模拟一个工作负载
			cond.L.Lock()
			// 下面的修改必须在cond.L被锁定的时候执行
			values[i] = string('a' + i)
			cond.Broadcast() // 可以在cond.L被解锁后发出通知
			cond.L.Unlock()
			// 上面的通知也可以在cond.L未锁定的时候发出。
			//cond.Broadcast() // 上面的调用也可以放在这里
		}(i)
	}

	// 此函数必须在cond.L被锁定的时候调用。
	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}

	cond.L.Lock()
	defer cond.L.Unlock()
	for !checkCondition() {
		cond.Wait() // 必须在cond.L被锁定的时候调用
	}
}

/*
Cond值所表示的自定义条件可以是一个虚无。对于这种情况，此Cond值纯粹被用来实现通知。 比如，下面这个程序将打印出abc或者bac。
*/
func TestA31(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	go func() {
		cond.L.Lock()
		go func() {
			cond.L.Lock()
			cond.Broadcast()
			cond.L.Unlock()
		}()
		cond.Wait()
		fmt.Print("a")
		cond.L.Unlock()
		wg.Done()
	}()
	cond.Wait()
	fmt.Print("b")
	cond.L.Unlock()

	wg.Wait()
	fmt.Println("c")
}

// 整数原子操作
func TestA41(t *testing.T) {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(atomic.AddInt32(&n, 1))
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(atomic.LoadInt32(&n)) // 1000

}

type Page struct {
	views uint32
}

func (page *Page) SetViews(n uint32) {
	atomic.StoreUint32(&page.views, n)

}

func (page *Page) Views() uint32 {
	return atomic.LoadUint32(&page.views)
}

func TestA51(t *testing.T) {
	var (
		n uint64 = 97
		m uint64 = 1
		k int    = 2
	)
	const (
		a        = 3
		b uint64 = 4
		c uint32 = 5
		d int    = 6
	)

	show := fmt.Println
	atomic.AddUint64(&n, -m)
	show(n) // 96 (97 - 1)

	atomic.AddUint64(&n, -uint64(k))
	show(n) // 94 (96 - 2)

	atomic.AddUint64(&n, ^uint64(a-1))
	show(n) // 91 (94 - 3)

	atomic.AddUint64(&n, ^(b - 1))
	show(n) // 87 (91 - 4)

	atomic.AddUint64(&n, ^uint64(c-1))
	show(n) // 82 (87 - 5)
	atomic.AddUint64(&n, ^uint64(d-1))
	show(n) // 76 (82 - 6)

	x := b
	atomic.AddUint64(&n, -x)
	show(n) // 72 (76 - 4)
	atomic.AddUint64(&n, ^(m - 1))
	show(n) // 71 (72 - 1)
	atomic.AddUint64(&n, ^uint64(k-1))
	show(n) // 69 (71 - 2)
}

/*
SwapT函数调用和StoreT函数调用类似，但是返回修改之前的旧值（因此称为置换操作）。

一个CompareAndSwapT函数调用仅在新值和旧值 相等 的情况下才会执行修改操作，并返回true；否则立即返回false。
*/
func TestA61(t *testing.T) {
	var n int64 = 123
	var old = atomic.SwapInt64(&n, 789)
	fmt.Println(n, old) // 789 123

	swapped := atomic.CompareAndSwapInt64(&n, 123, 456)
	fmt.Println(swapped) // false
	fmt.Println(n)       // 789
	swapped = atomic.CompareAndSwapInt64(&n, 789, 456)
	fmt.Println(swapped) // true
	fmt.Println(n)       // 456
}

func TestA71(t *testing.T) {
	p0 := new(int)   // p0指向一个int类型的零值
	fmt.Println(p0)  // （打印出一个十六进制形式的地址）
	fmt.Println(*p0) // 0

	x := *p0         // x是p0所引用的值的一个复制。
	p1, p2 := &x, &x // p1和p2中都存储着x的地址。
	// x、*p1和*p2表示着同一个int值。
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false
	p3 := &*p0            // <=> p3 := &(*p0)
	// <=> p3 := p0
	// p3和p0中存储的地址是一样的。
	fmt.Println(p0 == p3) // true
	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3) // 789 789 123

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
}

func double(x *int) {
	*x += *x
	x = nil // 此行仅为讲解目的
}

func TestA72(t *testing.T) {
	var a = 3
	double(&a)
	fmt.Println(a) // 6
	p := &a
	double(p)
	fmt.Println(a, p == nil) // 12 false
}
func TestA73(t *testing.T) {
	a := int64(5)
	p := &a

	// 下面这两行编译不通过。
	/*
		p++
		p = (&a) + 8
	*/

	*p++
	fmt.Println(*p, a)   // 6 6
	fmt.Println(p == &a) // true

	*&a++
	*&*&a++
	**&p++
	*&*p++
	fmt.Println(*p, a) // 10 10
}
func TestA74(t *testing.T) {

}

func TestA75(t *testing.T) {

}
func TestA76(t *testing.T) {

}
func TestA77(t *testing.T) {

}
func TestA78(t *testing.T) {

}
