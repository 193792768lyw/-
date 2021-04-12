package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	wg.Done() // 通知当前任务已经完成。
}

func TestMain1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2) // 注册两个新任务。
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	wg.Wait() // 阻塞在这里，直到所有任务都已完成。
}

// 延迟函数调用（deferred function call）
/*
在Go中，一个函数调用可以跟在一个defer关键字后面，形成一个延迟函数调用。 和协程调用类似，被延迟的函数调用的所有返回值必须全部被舍弃。
当一个函数调用被延迟后，它不会立即被执行。它将被推入由当前协程维护的一个延迟调用堆栈。 当一个函数调用（可能是也可能不是一个延迟调用）
返回并进入它的退出阶段后，所有在此函数调用中已经被推入的延迟调用将被按照它们被推入堆栈的顺序逆序执行。
当所有这些延迟调用执行完毕后，此函数调用也就真正退出了

事实上，每个协程维护着两个调用堆栈。
一个是正常的函数调用堆栈。在此堆栈中，相邻的两个调用存在着调用关系。晚进入堆栈的调用被早进入堆栈的调用所调用。 此堆栈中最早被推入的调用是对应协程的启动调用。
另一个堆栈是上面提到的延迟调用堆栈。处于延迟调用堆栈中的任意两个调用之间不存在调用关系。
*/
func TestMain2(t *testing.T) {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if false {
		defer fmt.Println("not reachable")
	}
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
	defer fmt.Println("not reachable")
}

// 一个延迟调用可以修改包含此延迟调用的最内层函数的返回值
//func Triple(n int) int {
//	r := n + n
//	defer func() {
//		r += n // 修改返回值
//	}()
//	return  r  // <=> 一个匿名的变量 = r; return
//}
//
//func TestMain3(t *testing.T) {
//	fmt.Println(Triple(5)) // 10
//}

func Triple(n int) (r int) {
	defer func() {
		r += n // 修改返回值
	}()

	return n + n // <=> r = n + n; return
}
func TestMain9(t *testing.T) {
	fmt.Println(Triple(5)) // 15
}

/*
协程和延迟调用的实参的估值时刻
一个协程调用或者延迟调用的实参是在此调用发生时被估值的。更具体地说，
	对于一个延迟函数调用，它的实参是在此调用被推入延迟调用堆栈的时候被估值的。
	对于一个协程调用，它的实参是在此协程被创建的时候估值的。
一个匿名函数体内的表达式是在此函数被执行的时候才会被逐个估值的，不管此函数是被普通调用还是延迟/协程调用。
*/

func TestMain4(t *testing.T) {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i)
			}()
		}
	}()
}

/*
第一个匿名函数中的循环打印出2、1和0这个序列，但是第二个匿名函数中的循环打印出三个3。
因为第一个循环中的i是在fmt.Println函数调用被推入延迟调用堆栈的时候估的值，
而第二个循环中的i是在第二个匿名函数调用的退出阶段估的值（此时循环变量i的值已经变为3）。
我们可以对第二个循环略加修改（使用两种方法），使得它和第一个循环打印出相同的结果。
		for i := 0; i < 3; i++ {
			defer func(i int) {
				// 此i为形参i，非实参循环变量i。
				fmt.Println("b:", i)
			}(i)
		}
或者
for i := 0; i < 3; i++ {
			i := i // 在下面的调用中，左i遮挡了右i。
			       // <=> var i = i
			defer func() {
				// 此i为上面的左i，非循环变量i。
				fmt.Println("b:", i)
			}()
		}
*/

// 同样的估值时刻规则也适用于协程调用。下面这个例子程序将打印出123 789。
func TestMain5(t *testing.T) {
	var a = 123
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Println(x, a) // 123 789
	}(a)

	a = 789

	time.Sleep(2 * time.Second)
}

// 恐慌（panic）和恢复（recover）
/*
进入恐慌状况是另一种使当前函数调用开始返回的途径。 一旦一个函数调用产生一个恐慌，此函数调用将立即进入它的退出阶段，
在此函数调用中被推入堆栈的延迟调用将按照它们被推入的顺序逆序执行。
通过在一个延迟函数调用之中调用内置函数recover，当前协程中的一个恐慌可以被消除，从而使得当前协程重新进入正常状况。

一个recover函数的返回值为其所恢复的恐慌在产生时被一个panic函数调用所消费的参数。

在一个处于恐慌状况的协程退出之前，其中的恐慌不会蔓延到其它协程。 如果一个协程在恐慌状况下退出，它将使整个程序崩溃。
*/

func TestMain6(t *testing.T) {
	defer func() {
		v := recover() // panic 只会被最近的一个recover
		fmt.Println("正常退出", v)
	}()
	fmt.Println("嗨！")
	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了：", v)
	}()
	panic("拜拜！") // 产生一个恐慌
	fmt.Println("执行不到这里")
}

// 下面的例子在一个新协程里面产生了一个恐慌，并且此协程在恐慌状况下退出，所以整个程序崩溃了。
func TestMain7(t *testing.T) {
	fmt.Println("hi!")

	go func() {
		time.Sleep(time.Second)
		panic(123)
	}()

	for {
		time.Sleep(time.Second)
	}
}

func TestMain8(t *testing.T) {
	fmt.Println("hi!")

	go func() {
		defer func() {
			v := recover()
			fmt.Println("恐慌被恢复了：", v)
		}()
		time.Sleep(time.Second)
		panic(123)
	}()

	for {
		time.Sleep(time.Second)
	}
}

// 通道  Go提供了一种独特的并发同步技术来实现通过通讯来共享内存。此技术即为通道。
/*
通道的主要作用是用来实现并发同步。
通过共享内存来通讯和通过通讯来共享内存是并发编程中的两种编程风格。 当通过共享内存来通讯的时候，我们需要一些传统的并发同步技术（比如互斥锁）来避免数据竞争。


通道操作
Go中有五种通道相关的操作。假设一个通道（值）为ch，下面列出了这五种操作的语法或者函数调用。
1 、 调用内置函数close来关闭一个通道：
	close(ch)
	传给close函数调用的实参必须为一个通道值，并且此通道值不能为单向接收的。 func close(c chan<- Type)
*/

func TestMain10(t *testing.T) {
	//ch := make(chan int)
	//var b <- chan int
	//b = ch

	//close(b)

}

// 一个简单的通过一个非缓冲通道实现的请求/响应的例子：

func TestMain11(t *testing.T) {

	c := make(chan int) // 一个非缓冲通道
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		// <-ch    // 此操作编译不通过
		ch <- x * x // 阻塞在此，直到发送的值被接收
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch      // 阻塞在此，直到有值发送到c
		fmt.Println(n) // 9
		// ch <- 123   // 此操作编译不通过
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)
	<-done // 阻塞在此，直到有值发送到done
	fmt.Println("bye")
}

// 下面的例子使用了一个缓冲通道。此例子程序并非是一个并发程序，它只是为了展示缓冲通道的使用。

func TestMain12(t *testing.T) {
	c := make(chan int, 2) // 一个容量为2的缓冲通道
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok)          // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	x, ok = <-c
	fmt.Println(x, ok)          // 0 false
	fmt.Println(len(c), cap(c)) // 0 2
	close(c)                    // 此行将产生一个恐慌
	c <- 7                      // 如果上一行不存在，此行也将产生一个恐慌。
}

// 一场永不休场的足球比赛：
func TestMain13(t *testing.T) {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Print(<-ball, "传球", "\n")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("张三")
	go kickBall("李四")
	go kickBall("王二麻子")
	go kickBall("刘大")
	ball <- "裁判"    // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
}

// 在下面这个例子中，数据接收和发送操作被用在两个for循环的初始化和步尾语句。
func TestMain14(t *testing.T) {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63); c <- y { // 步尾语句
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}
	c := fibonacci()
	//for x, ok := <-c; ok; x, ok = <-c { // 初始化和步尾语句
	//	time.Sleep(time.Second)
	//	fmt.Println(x)
	//}
	for x := range c {
		time.Sleep(time.Second)
		fmt.Println(x)
	}

}

/*
在下面这个例子中，default分支将铁定得到执行，因为两个case分支后的操作均为阻塞的。
*/
func TestMain15(t *testing.T) {
	var c chan struct{} // nil
	select {
	case <-c: // 阻塞操作
	case c <- struct{}{}: // 阻塞操作
	default:
		fmt.Println("Go here.")
	}
}

/*
下面这个例子中实现了尝试发送（try-send）和尝试接收（try-receive）。 它们都是用含有一个case分支和一个default分支的select-case代码块来实现的。
*/
func TestMain16(t *testing.T) {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default: // 如果c的缓冲已满，则执行默认分支。
		}
	}
	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-" // 如果c的缓冲为空，则执行默认分支。
		}
	}
	trySend("Hello!") // 发送成功
	trySend("Hi!")    // 发送成功
	trySend("Bye!")   // 发送失败，但不会阻塞。
	// 下面这两行将接收成功。
	fmt.Println(tryReceive()) // Hello!
	fmt.Println(tryReceive()) // Hi!
	// 下面这行将接收失败。
	fmt.Println(tryReceive()) // -
}

/*
下面这个程序有50%的几率会因为恐慌而崩溃。 此程序中select-case代码块中的两个case操作均不阻塞，所以随机一个将被执行。
如果第一个case操作（向已关闭的通道发送数据）被执行，则一个恐慌将产生。
*/
func TestMain17(t *testing.T) {
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}: // 若此分支被选中，则产生一个恐慌
	case <-c:
	}
}

/*
在Go中，如果我们能够保证从不会向一个通道发送数据，那么有一个简单的方法来判断此通道是否已经关闭。
如前所述，此方法并不是一个通用的检查通道是否已经关闭的方法。

事实上，即使有一个内置closed函数用来检查一个通道是否已经关闭，它的有用性也是十分有限的。
原因是当此函数的一个调用的结果返回时，被查询的通道的状态可能已经又改变了，导致此调用结果并不能反映出被查询的通道的最新状态。
虽然我们可以根据一个调用closed(ch)的返回结果为true而得出我们不应该再向通道ch发送数据的结论，
但是我们不能根据一个调用closed(ch)的返回结果为false而得出我们可以继续向通道ch发送数据的结论。
*/

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func TestFun1(t *testing.T) {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
}

// 通道关闭原则

/*
不要在数据接收方或者在有多个发送者的情况下关闭通道 , 换句话说，我们只应该让一个通道唯一的发送者关闭此通道。
*/
// 礼貌地关闭通道的方法
/*
type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}


type MyChannel struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
}

func (mc *MyChannel) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}


上述方式不能完全有效地避免数据竞争。 目前的Go白皮书并不保证发生在一个通道上的并发关闭操作和发送操纵不会产生数据竞争。
如果一个SafeClose函数和同一个通道上的发送操作同时运行，则数据竞争可能发生（虽然这样的数据竞争一般并不会带来什么危害）。

*/

// 优雅地关闭通道的方法
// 情形一：M个接收者和一个发送者。发送者通过关闭用来传输数据的通道来传递发送结束信号
// 这是最简单的一种情形。当发送者欲结束发送，让它关闭用来传输数据的通道即可。
func TestMain19(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)

	// 发送者
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				// 此唯一的发送者可以安全地关闭此数据通道。
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// 接收数据直到通道dataCh已关闭
			// 并且dataCh的缓冲队列已空。
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

// 情形二：一个接收者和N个发送者，此唯一接收者通过关闭一个额外的信号通道来通知发送者不要在发送数据了
/*
此情形比上一种情形复杂一些。我们不能让接收者关闭用来传输数据的通道来停止数据传输，因为这样做违反了通道关闭原则。
但是我们可以让接收者关闭一个额外的信号通道来通知发送者不要在发送数据了。

如此例中的注释所述，对于此额外的信号通道stopCh，它只有一个发送者，即dataCh数据通道的唯一接收者。
dataCh数据通道的接收者关闭了信号通道stopCh，这是不违反通道关闭原则的。

在此例中，数据通道dataCh并没有被关闭。是的，我们不必关闭它。 当一个通道不再被任何协程所使用后，它将逐渐被垃圾回收掉，无论它是否已经被关闭。
所以这里的优雅性体现在通过不关闭一个通道来停止使用此通道。
*/
func TestMain20(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh是一个额外的信号通道。它的
	// 发送者为dataCh数据通道的接收者。
	// 它的接收者为dataCh数据通道的发送者。

	// 发送者
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// 这里的第一个尝试接收用来让此发送者
				// 协程尽早地退出。对于这个特定的例子，
				// 此select代码块并非必需。
				select {
				case <-stopCh:
					return
				default:
				}

				// 即使stopCh已经关闭，此第二个select
				// 代码块中的第一个分支仍很有可能在若干个
				// 循环步内依然不会被选中。如果这是不可接受
				// 的，则上面的第一个select代码块是必需的。
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// 接收者
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				// 此唯一的接收者同时也是stopCh通道的
				// 唯一发送者。尽管它不能安全地关闭dataCh数
				// 据通道，但它可以安全地关闭stopCh通道。
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}

// 情形三：M个接收者和N个发送者。它们中的任何协程都可以让一个中间调解协程帮忙发出停止数据传送的信号
/*
这是最复杂的一种情形。我们不能让接收者和发送者中的任何一个关闭用来传输数据的通道，我们也不能让多个接收者之一关闭一个额外的信号通道。
这两种做法都违反了通道关闭原则。 然而，我们可以引入一个中间调解者角色并让其关闭额外的信号通道来通知所有的接收者和发送者结束工作。
具体实现见下例。注意其中使用了一个尝试发送操作来向中间调解者发送信号。

在此例中，通道关闭原则依旧得到了遵守。

请注意，信号通道toStop的容量必须至少为1。 如果它的容量为0，则在中间调解者还未准备好的情况下就已经有某个协程向toStop发送信号时，
此信号将被抛弃。

我们也可以不使用尝试发送操作向中间调解者发送信号，但信号通道toStop的容量必须至少为数据发送者和数据接收者的数量之和，
以防止向其发送数据时（有一个极其微小的可能）导致某些发送者和接收者协程永久阻塞。

...
toStop := make(chan string, NumReceivers + NumSenders)
...
			value := rand.Intn(Max)
			if value == 0 {
				toStop <- "sender#" + id
				return
			}
...
				if value == Max-1 {
					toStop <- "receiver#" + id
					return
				}
...
*/
func TestMain21(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh是一个额外的信号通道。它的发送
	// 者为中间调解者。它的接收者为dataCh
	// 数据通道的所有的发送者和接收者。
	toStop := make(chan string, 1)
	// toStop是一个用来通知中间调解者让其
	// 关闭信号通道stopCh的第二个信号通道。
	// 此第二个信号通道的发送者为dataCh数据
	// 通道的所有的发送者和接收者，它的接收者
	// 为中间调解者。它必须为一个缓冲通道。

	var stoppedBy string

	// 中间调解者
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// 发送者
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					// 为了防止阻塞，这里使用了一个尝试
					// 发送操作来向中间调解者发送信号。
					select {
					case toStop <- "发送者#" + id:
					default:
					}
					return
				}

				// 此处的尝试接收操作是为了让此发送协程尽早
				// 退出。标准编译器对尝试接收和尝试发送做了
				// 特殊的优化，因而它们的速度很快。
				select {
				case <-stopCh:
					return
				default:
				}

				// 即使stopCh已关闭，如果这个select代码块
				// 中第二个分支的发送操作是非阻塞的，则第一个
				// 分支仍很有可能在若干个循环步内依然不会被选
				// 中。如果这是不可接受的，则上面的第一个尝试
				// 接收操作代码块是必需的。
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// 和发送者协程一样，此处的尝试接收操作是为了
				// 让此接收协程尽早退出。
				select {
				case <-stopCh:
					return
				default:
				}

				// 即使stopCh已关闭，如果这个select代码块
				// 中第二个分支的接收操作是非阻塞的，则第一个
				// 分支仍很有可能在若干个循环步内依然不会被选
				// 中。如果这是不可接受的，则上面尝试接收操作
				// 代码块是必需的。
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						// 为了防止阻塞，这里使用了一个尝试
						// 发送操作来向中间调解者发送信号。
						select {
						case toStop <- "接收者#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("被" + stoppedBy + "终止了")
}

// 情形四：“M个接收者和一个发送者”情形的一个变种：用来传输数据的通道的关闭请求由第三方发出
/*
有时，数据通道（dataCh）的关闭请求需要由某个第三方协程发出。对于这种情形，
我们可以使用一个额外的信号通道来通知唯一的发送者关闭数据通道（dataCh）。
上述代码中的stop函数中使用的技巧偷自Roger Peppe在此贴中的一个留言。
*/
func TestMain22(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	closing := make(chan struct{}) // 信号通道
	closed := make(chan struct{})

	// 此stop函数可以被安全地多次调用。
	stop := func() {
		select {
		case closing <- struct{}{}:
			<-closed
		case <-closed:
		}
	}

	// 一些第三方协程
	for i := 0; i < NumThirdParties; i++ {
		go func() {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop()
		}()
	}

	// 发送者
	go func() {
		defer func() {
			close(closed)
			close(dataCh)
		}()

		for {
			select {
			case <-closing:
				return
			default:
			}

			select {
			case <-closing:
				return
			case dataCh <- rand.Intn(Max):
			}
		}
	}()

	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()

}

// 情形五：“N个发送者”的一个变种：用来传输数据的通道必须被关闭以通知各个接收者数据发送已经结束了
/*
在上面的提到的“N个发送者”情形中，为了遵守通道关闭原则，我们避免了关闭数据通道（dataCh）。 但是有时候，
数据通道（dataCh）必须被关闭以通知各个接收者数据发送已经结束。 对于这种“N个发送者”情形，
我们可以使用一个中间通道将它们转化为“一个发送者”情形，然后继续使用上一节介绍的技巧来关闭此中间通道，从而避免了关闭原始的dataCh数据通道。
*/
func TestMain23(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 1000000
	const NumReceivers = 10
	const NumSenders = 1000
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)   // 将被关闭
	middleCh := make(chan int) // 不会被关闭
	closing := make(chan string)
	closed := make(chan struct{})

	var stoppedBy string

	stop := func(by string) {
		select {
		case closing <- by:
			<-closed
		case <-closed:
		}
	}

	// 中间层
	go func() {
		exit := func(v int, needSend bool) {
			close(closed)
			if needSend {
				dataCh <- v
			}
			close(dataCh)
		}

		for {
			select {
			case stoppedBy = <-closing:
				exit(0, false)
				return
			case v := <-middleCh:
				select {
				case stoppedBy = <-closing:
					exit(v, true)
					return
				case dataCh <- v:
				}
			}
		}
	}()

	// 一些第三方协程
	for i := 0; i < NumThirdParties; i++ {
		go func(id string) {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop("3rd-party#" + id)
		}(strconv.Itoa(i))
	}

	// 发送者
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					stop("sender#" + id)
					return
				}

				select {
				case <-closed:
					return
				default:
				}

				select {
				case <-closed:
					return
				case middleCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// 接收者
	for range [NumReceivers]struct{}{} {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

func TestA(t *testing.T) {
	ch := make(chan int)
	stopCh := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Second)
		close(stopCh)
	}()
	select {
	case <-ch:
		fmt.Println("dvnlsubngfs")
	case <-stopCh:
		fmt.Println("b ndfglkbjuslnrglbn")
	}

}

// 测试switch case
func TestMain18(t *testing.T) {
	switch "oo1" {
	case "oo":
		fmt.Println("bgfjln")
		//fallthrough
	case "88":
		fmt.Println("dbfnku")

	}
}
