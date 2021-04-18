package main

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"testing"
)

//func longTimeRequest() <-chan int32 {
//	r := make(chan int32)
//
//	go func() {
//		time.Sleep(time.Second * 3) // 模拟一个工作负载
//		r <- rand.Int31n(100)
//	}()
//
//	return r
//}
//
//func sumSquares(a, b int32) int32 {
//	return a*a + b*b
//}
//// 返回单向接收通道做为函数返回结果
///*
//在下面这个例子中，sumSquares函数调用的两个实参请求并发进行。 每个通道读取操作将阻塞到请求返回结果为止。
//两个实参总共需要大约3秒钟（而不是6秒钟）准备完毕（以较慢的一个为准）。
// */
//func TestMain1(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	a, b := longTimeRequest(), longTimeRequest()
//	fmt.Println(sumSquares(<-a, <-b))
//}

//func longTimeRequest(r chan<- int32)  {
//	time.Sleep(time.Second * 3) // 模拟一个工作负载
//	r <- rand.Int31n(100)
//}
//
//func sumSquares(a, b int32) int32 {
//	return a*a + b*b
//}
////将单向发送通道类型用做函数实参
///*
//和上例一样，在下面这个例子中，sumSquares函数调用的两个实参的请求也是并发进行的。
//和上例不同的是longTimeRequest函数接收一个单向发送通道类型参数而不是返回一个单向接收通道结果。
//对于上面这个特定的例子，我们可以只使用一个通道来接收回应结果，因为两个参数的作用是对等的。
//...
//
//	results := make(chan int32, 2) // 缓冲与否不重要
//	go longTimeRequest(results)
//	go longTimeRequest(results)
//
//	fmt.Println(sumSquares(<-results, <-results))
//}
// */
//func TestMain2(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	ra, rb := make(chan int32), make(chan int32)
//	go longTimeRequest(ra)
//	go longTimeRequest(rb)
//
//	fmt.Println(sumSquares(<-ra, <-rb))
//}
//
//func source(c chan<- int32) {
//	ra, rb := rand.Int31(), rand.Intn(3) + 1
//	// 睡眠1秒/2秒/3秒
//	time.Sleep(time.Duration(rb) * time.Second)
//	c <- ra
//}
//
//
//// 采用最快回应
///*
//本用例可以看作是上例中只使用一个通道变种的增强。
//有时候，一份数据可能同时从多个数据源获取。这些数据源将返回相同的数据。 因为各种因素，这些数据源的回应速度参差不一，
//甚至某个特定数据源的多次回应速度之间也可能相差很大。 同时从多个数据源获取一份相同的数据可以有效保障低延迟。
//我们只需采用最快的回应并舍弃其它较慢回应。
//注意：如果有N个数据源，为了防止被舍弃的回应对应的协程永久阻塞，则传输数据用的通道必须为一个容量至少为N-1的缓冲通道。
// */
//func TestMain3(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	startTime := time.Now()
//	c := make(chan int32, 5) // 必须用一个缓冲通道
//	for i := 0; i < cap(c); i++ {
//		go source(c)
//	}
//	rnd := <- c // 只有第一个回应被使用了
//	fmt.Println(time.Since(startTime))
//	fmt.Println(rnd)
//
//}
//
///*
//使用通道实现通知
//通知可以被看作是特殊的请求/回应用例。在一个通知用例中，我们并不关心回应的值，我们只关心回应是否已发生。
//所以我们常常使用空结构体类型struct{}来做为通道的元素类型，因为空结构体类型的尺寸为零，能够节省一些内存（虽然常常很少量）
// */
//
//
//
//
//
//
//
//
////向一个通道发送一个值来实现单对单通知
///*
//我们已知道，如果一个通道中无值可接收，则此通道上的下一个接收操作将阻塞到另一个协程发送一个值到此通道为止。
//所以一个协程可以向此通道发送一个值来通知另一个等待着从此通道接收数据的协程。
//
//在下面这个例子中，通道done被用来做为一个信号通道来实现单对单通知。
// */
//func TestMain4(t *testing.T) {
//
//	values := make([]byte, 32 * 1024 * 1024)
//	if _, err := rand.Read(values); err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	done := make(chan struct{}) // 也可以是缓冲的
//
//	// 排序协程
//	go func() {
//		sort.Slice(values, func(i, j int) bool {
//			return values[i] < values[j]
//		})
//		done <- struct{}{} // 通知排序已完成
//	}()
//
//	// 并发地做一些其它事情...
//
//	<- done // 等待通知
//	fmt.Println(values[0], values[len(values)-1])
//}
//
////从一个通道接收一个值来实现单对单通知
//
///*
//如果一个通道的数据缓冲队列已满（非缓冲的通道的数据缓冲队列总是满的）但它的发送协程队列为空，则向此通道发送一个值将阻塞，
//直到另外一个协程从此通道接收一个值为止。 所以我们可以通过从一个通道接收数据来实现单对单通知。一般我们使用非缓冲通道来实现这样的通知。
//
//这种通知方式不如上例中介绍的方式使用得广泛。
// */
//
//func TestMain5(t *testing.T) {
//	done := make(chan struct{})
//	// 此信号通道也可以缓冲为1。如果这样，则在下面
//	// 这个协程创建之前，我们必须向其中写入一个值。
//
//	go func() {
//		fmt.Print("Hello")
//		// 模拟一个工作负载。
//		time.Sleep(time.Second * 2)
//
//		// 使用一个接收操作来通知主协程。
//		<- done
//	}()
//
//	done <- struct{}{} // 阻塞在此，等待通知
//	fmt.Println(" world!")
//}
//
//// 多对单和单对多通知
//
//type T = struct{}
//
//func worker(id int, ready <-chan T, done chan<- T) {
//	<-ready // 阻塞在此，等待通知
//	log.Print("Worker#", id, "开始工作")
//	// 模拟一个工作负载。
//	time.Sleep(time.Second * time.Duration(id+1))
//	log.Print("Worker#", id, "工作完成")
//	done <- T{} // 通知主协程（N-to-1）
//}
///*
//事实上，上例中展示的多对单和单对多通知实现方式在实践中用的并不多。 在实践中，我们多使用sync.WaitGroup来实现多对单通知，
//使用关闭一个通道的方式来实现单对多通知（详见下一个用例）。
// */
//func TestMain6(t *testing.T) {
//	log.SetFlags(0)
//
//	ready, done := make(chan T), make(chan T)
//	go worker(0, ready, done)
//	go worker(1, ready, done)
//	go worker(2, ready, done)
//
//	// 模拟一个初始化过程
//	time.Sleep(time.Second * 3 / 2)
//	// 单对多通知
//	ready <- T{}; ready <- T{}; ready <- T{}
//	// 等待被多对单通知
//	<-done; <-done; <-done
//}
//
//func AfterDuration(d time.Duration) <- chan struct{} {
//	c := make(chan struct{}, 1)
//	go func() {
//		time.Sleep(d)
//		c <- struct{}{}
//	}()
//	return c
//}
//
////定时通知（timer）
////用通道实现一个一次性的定时通知器是很简单的。 下面是一个自定义实现：
///*
//事实上，time标准库包中的After函数提供了和上例中AfterDuration同样的功能。 在实践中，我们应该尽量使用time.After函数以使代码看上去更干净。
//
//注意，操作<-time.After(aDuration)将使当前协程进入阻塞状态，而一个time.Sleep(aDuration)函数调用不会如此。
//
//<-time.After(aDuration)经常被使用在后面将要介绍的超时机制实现中。
// */
//func TestMain7(t *testing.T) {
//
//	fmt.Println("Hi!")
//	<- AfterDuration(time.Second)
//	fmt.Println("Hello!")
//	<- AfterDuration(time.Second)
//	fmt.Println("Bye!")
//}
//
///*
//将通道用做互斥锁（mutex）
//上面的某个例子提到了容量为1的缓冲通道可以用做一次性二元信号量。
//事实上，容量为1的缓冲通道也可以用做多次性二元信号量（即互斥锁）尽管这样的互斥锁效率不如sync标准库包中提供的互斥锁高效。
//
//有两种方式将一个容量为1的缓冲通道用做互斥锁：
//通过发送操作来加锁，通过接收操作来解锁；
//通过接收操作来加锁，通过发送操作来解锁。
//下面是一个通过发送操作来加锁的例子。
// */
//
//func TestMain8(t *testing.T) {
//	mutex := make(chan struct{}, 1) // 容量必须为1
//
//	counter := 0
//	increase := func() {
//		mutex <- struct{}{} // 加锁
//		counter++
//		<-mutex // 解锁
//	}
//
//	increase1000 := func(done chan<- struct{}) {
//		for i := 0; i < 1000; i++ {
//			increase()
//		}
//		done <- struct{}{}
//	}
//
//	done := make(chan struct{})
//	go increase1000(done)
//	go increase1000(done)
//	<-done; <-done
//	fmt.Println(counter) // 2000
//}
//// 下面是一个通过接收操作来加锁的例子，其中只显示了相对于上例而修改了的部分。
//func TestMain9(t *testing.T) {
//
//	mutex := make(chan struct{}, 1)
//	mutex <- struct{}{} // 此行是必需的
//
//
//	counter := 0
//	increase := func() {
//		<-mutex // 加锁
//		counter++
//		mutex <- struct{}{} // 解锁
//	}
//
//	increase1000 := func(done chan<- struct{}) {
//		for i := 0; i < 1000; i++ {
//			increase()
//		}
//		done <- struct{}{}
//	}
//
//	done := make(chan struct{})
//	go increase1000(done)
//	go increase1000(done)
//	<-done; <-done
//	fmt.Println(counter) // 2000
//
//}
//
///*
//将通道用做计数信号量（counting semaphore）
//缓冲通道可以被用做计数信号量。 计数信号量可以被视为多主锁。如果一个缓冲通道的容量为N，那么它可以被看作是一个在任何时刻最多可有N个主人的锁。
//上面提到的二元信号量是特殊的计数信号量，每个二元信号量在任一时刻最多只能有一个主人。
//
//计数信号量经常被使用于限制最大并发数。
//
//和将通道用做互斥锁一样，也有两种方式用来获取一个用做计数信号量的通道的一份所有权。
//通过发送操作来获取所有权，通过接收操作来释放所有权；
//通过接收操作来获取所有权，通过发送操作来释放所有权。
//下面是一个通过接收操作来获取所有权的例子：
// */
//
//type Seat int
//type Bar chan Seat
//
//func (bar Bar) ServeCustomer(c int) {
//	log.Print("顾客#", c, "进入酒吧")
//	seat := <- bar // 需要一个位子来喝酒
//	log.Print("++ customer#", c, " drinks at seat#", seat)
//	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
//	time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
//	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
//	bar <- seat // 释放座位，离开酒吧
//}
//
//
//func TestMain10(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	bar24x7 := make(Bar, 10) // 此酒吧有10个座位
//	// 摆放10个座位。
//	for seatId := 0; seatId < cap(bar24x7); seatId++ {
//		bar24x7 <- Seat(seatId) // 均不会阻塞
//	}
//
//	for customerId := 0; ; customerId++ {
//		time.Sleep(time.Second)
//		go bar24x7.ServeCustomer(customerId)
//	}
//	for {time.Sleep(time.Second)} // 睡眠不属于阻塞状态
//}
//
///*
//在上例中，只有获得一个座位的顾客才能开始饮酒。 所以在任一时刻同时在喝酒的顾客数不会超过座位数10。
//
//上例main函数中的最后一行for循环是为了防止程序退出。 后面将介绍一种更好的实现此目的的方法。
//
//在上例中，尽管在任一时刻同时在喝酒的顾客数不会超过座位数10，但是在某一时刻可能有多于10个顾客进入了酒吧，因为某些顾客在排队等位子。
//在上例中，每个顾客对应着一个协程。虽然协程的开销比系统线程小得多，但是如果协程的数量很多，则它们的总体开销还是不能忽略不计的。
//所以，最好当有空位的时候才创建顾客协程。
// */
//
//func (bar Bar) ServeCustomerAtSeat(c int, seat Seat) {
//	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
//	time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
//	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
//	bar <- seat // 释放座位，离开酒吧
//}
//
//func TestMain11(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	bar24x7 := make(Bar, 10)
//	for seatId := 0; seatId < cap(bar24x7); seatId++ {
//		bar24x7 <- Seat(seatId)
//	}
//
//	// 这个for循环和上例不一样。
//	for customerId := 0; ; customerId++ {
//		time.Sleep(time.Second)
//		seat := <- bar24x7 // 需要一个空位招待顾客
//		go bar24x7.ServeCustomerAtSeat(customerId, seat)
//	}
//	for {time.Sleep(time.Second)}
//}
//
///*
//}
//在上面这个修改后的例子中，在任一时刻最多只有10个顾客协程在运行（但是在程序的生命期内，仍旧会有大量的顾客协程不断被创建和销毁）。
//
//在下面这个更加高效的实现中，在程序的生命期内最多只会有10个顾客协程被创建出来。
// */
//
//func (bar Bar) ServeCustomerAtSeat1(consumers chan int) {
//	for c := range consumers {
//		seatId := <- bar
//		log.Print("++ 顾客#", c, "在第", seatId, "个座位开始饮酒")
//		time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
//		log.Print("-- 顾客#", c, "离开了第", seatId, "个座位")
//		bar <- seatId // 释放座位，离开酒吧
//	}
//}
//
//func TestMain14(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	bar24x7 := make(Bar, 10)
//	for seatId := 0; seatId < cap(bar24x7); seatId++ {
//		bar24x7 <- Seat(seatId)
//	}
//
//	consumers := make(chan int)
//	for i := 0; i < cap(bar24x7); i++ {
//		go bar24x7.ServeCustomerAtSeat1(consumers)
//	}
//
//	for customerId := 0; ; customerId++ {
//		time.Sleep(time.Second)
//		consumers <- customerId
//	}
//}
//
//
///*
//题外话：当然，如果我们并不关心座位号（这种情况在编程实践中很常见），则实际上bar24x7计数信号量是完全不需要的：
// */
//func ServeCustomer(consumers chan int) {
//	for c := range consumers {
//		log.Print("++ 顾客#", c, "开始在酒吧饮酒")
//		time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
//		log.Print("-- 顾客#", c, "离开了酒吧")
//	}
//}
//
//func TestMain12(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	const BarSeatCount = 10
//	consumers := make(chan int)
//	for i := 0; i < BarSeatCount; i++ {
//		go ServeCustomer(consumers)
//	}
//
//	for customerId := 0; ; customerId++ {
//		time.Sleep(time.Second)
//		consumers <- customerId
//	}
//}
//
///// 通过发送操作来获取所有权的实现相对简单一些，省去了摆放座位的步骤。
////type Customer struct{id int}
////type Bar chan Customer
////
////func (bar Bar) ServeCustomer(c Customer) {
////	log.Print("++ 顾客#", c.id, "开始饮酒")
////	time.Sleep(time.Second * time.Duration(3 + rand.Intn(16)))
////	log.Print("-- 顾客#", c.id, "离开酒吧")
////	<- bar // 离开酒吧，腾出位子
////}
////
////func TestMain13(t *testing.T) {
////	rand.Seed(time.Now().UnixNano())
////
////	bar24x7 := make(Bar, 10) // 最对同时服务10位顾客
////	for customerId := 0; ; customerId++ {
////		time.Sleep(time.Second * 2)
////		customer := Customer{customerId}
////		bar24x7 <- customer // 等待进入酒吧
////		go bar24x7.ServeCustomer(customer)
////	}
////	for {time.Sleep(time.Second)}
////}
//
///*
//对话（或称乒乓）
//两个协程可以通过一个通道进行对话，整个过程宛如打乒乓球一样。 下面是一个这样的例子，它将打印出一系列斐波那契（Fibonacci）数。
// */
//type Ball uint64
//
//func Play(playerName string, table chan Ball) {
//	var lastValue Ball = 1
//	for {
//		ball := <- table // 接球
//		fmt.Println(playerName, ball)
//		ball += lastValue
//		if ball < lastValue { // 溢出结束
//			os.Exit(0)
//		}
//		lastValue = ball
//		table <- ball // 回球
//		time.Sleep(time.Second)
//	}
//}
//
//func TestMain21(t *testing.T) {
//	table := make(chan Ball)
//	go func() {
//		table <- 1 // （裁判）发球
//	}()
//	go Play("A:", table)
//	Play("B:", table)
//}
//
///*
//使用通道传送传输通道
//一个通道类型的元素类型可以是另一个通道类型。 在下面这个例子中， 单向发送通道类型chan<- int是另一个通道类型chan chan<- int的元素类型。
// */
//var counter = func (n int) chan<- chan<- int {
//	requests := make(chan chan<- int)
//	go func() {
//		for request := range requests {
//			if request == nil {
//				n++ // 递增计数
//			} else {
//				request <- n // 返回当前计数
//			}
//		}
//	}()
//	return requests // 隐式转换到类型chan<- (chan<- int)
//}(0)
//
//
//func TestMain22(t *testing.T) {
//	increase1000 := func(done chan<- struct{}) {
//		for i := 0; i < 1000; i++ {
//			counter <- nil
//		}
//		done <- struct{}{}
//	}
//
//	done := make(chan struct{})
//	go increase1000(done)
//	go increase1000(done)
//	<-done; <-done
//
//	request := make(chan int, 1)
//	counter <- request
//	fmt.Println(<-request) // 2000
//}
//
///*
//使当前协程永久阻塞
//Go中的选择机制（select）是一个非常独特的特性。它给并发编程带来了很多新的模式和技巧。
//
//我们可以用一个无分支的select流程控制代码块使当前协程永久处于阻塞状态。 这是select流程控制的最简单的应用。
//事实上，上面很多例子中的for {time.Sleep(time.Second)}都可以换为select{}。
//
//一般，select{}用在主协程中以防止程序退出。
//
//一个例子：
// */
//func DoSomething() {
//	for {
//		// 做点什么...
//
//		runtime.Gosched() // 防止本协程霸占CPU不放
//	}
//}
//
//func TestMain23(t *testing.T) {
//	go DoSomething()
//	go DoSomething()
//
//	select{}
//}
//
//func TestMain24(t *testing.T) {
//	a := chan struct{}(nil)
//	fmt.Println(a)
//}
//
///*
//尝试发送和尝试接收
//含有一个default分支和一个case分支的select代码块可以被用做一个尝试发送或者尝试接收操作，取决于case关键字后跟随的是一个发送操作还是一个接收操作。
//如果case关键字后跟随的是一个发送操作，则此select代码块为一个尝试发送操作。 如果case分支的发送操作是阻塞的，则default分支将被执行，发送失败；
//否则发送成功，case分支得到执行。
//如果case关键字后跟随的是一个接收操作，则此select代码块为一个尝试接收操作。 如果case分支的接收操作是阻塞的，则default分支将被执行，接收失败；
//否则接收成功，case分支得到执行。
//尝试发送和尝试接收代码块永不阻塞。
//
//标准编译器对尝试发送和尝试接收代码块做了特别的优化，使得它们的执行效率比多case分支的普通select代码块执行效率高得多。
//
//下例演示了尝试发送和尝试接收代码块的工作原理。
// */
//func TestMain25(t *testing.T) {
//
//
//	type Book struct{id int}
//	bookshelf := make(chan Book, 3)
//
//	for i := 0; i < cap(bookshelf) * 2; i++ {
//		select {
//		case bookshelf <- Book{id: i}:
//			fmt.Println("成功将书放在书架上", i)
//		default:
//			fmt.Println("书架已经被占满了")
//		}
//	}
//
//	for i := 0; i < cap(bookshelf) * 2; i++ {
//		select {
//		case book := <-bookshelf:
//			fmt.Println("成功从书架上取下一本书", book.id)
//		default:
//			fmt.Println("书架上已经没有书了")
//		}
//	}
//
//}
//
///*
//另一种“采用最快回应”的实现方式
//在上面的“采用最快回应”用例一节已经提到，我们也可以使用选择机制来实现“采用最快回应”用例。
//每个数据源协程只需使用一个缓冲为1的通道并向其尝试发送回应数据即可。示例代码如下：
//
//注意，使用选择机制来实现“采用最快回应”的代码中使用的通道的容量必须至少为1，以保证最快回应总能够发送成功。
//否则，如果数据请求者因为种种原因未及时准备好接收，则所有回应者的尝试发送都将失败，从而所有回应的数据都将被错过。
// */
//
//func source(c chan<- int32) {
//	ra, rb := rand.Int31(), rand.Intn(3)+1
//	// 休眠1秒/2秒/3秒
//	time.Sleep(time.Duration(rb) * time.Second)
//	select {
//	case c <- ra:
//	default:
//	}
//}
//
//func TestMain27(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	c := make(chan int32, 1) // 此通道容量必须至少为1
//	for i := 0; i < 5; i++ {
//		go source(c)
//	}
//	rnd := <-c // 只采用第一个成功发送的回应数据
//	fmt.Println(rnd)
//}
//
///*
//第三种“采用最快回应”的实现方式
//如果一个“采用最快回应”用例中的数据源的数量很少，比如两个或三个，我们可以让每个数据源使用一个单独的缓冲通道来回应数据，
//然后使用一个select代码块来同时接收这三个通道。 示例代码如下：
//
//注意：如果上例中使用的通道是非缓冲的，未被选中的case分支对应的两个source函数调用中开辟的协程将处于永久阻塞状态，从而造成内存泄露。
// */
//
//func source() <-chan int32 {
//	c := make(chan int32, 1) // 必须为一个缓冲通道
//	go func() {
//		ra, rb := rand.Int31(), rand.Intn(3)+1
//		time.Sleep(time.Duration(rb) * time.Second)
//		c <- ra
//	}()
//	return c
//}
//
//
//func TestMain28(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	var rnd int32
//	// 阻塞在此直到某个数据源率先回应。
//	select{
//	case rnd = <-source():
//	case rnd = <-source():
//	case rnd = <-source():
//	}
//	fmt.Println(rnd)
//}
///*
//超时机制（timeout）
//在一些请求/回应用例中，一个请求可能因为种种原因导致需要超出预期的时长才能得到回应，有时甚至永远得不到回应。
//对于这样的情形，我们可以使用一个超时方案给请求者返回一个错误信息。 使用选择机制可以很轻松地实现这样的一个超时方案。
//
//下面这个例子展示了如何实现一个支持超时设置的请求：
// */
//func requestWithTimeout(timeout time.Duration) (int, error) {
//	c := make(chan int)
//	go doRequest(c) // 可能需要超出预期的时长回应
//
//	select {
//	case data := <-c:
//		return data, nil
//	case <-time.After(timeout):
//		return 0, errors.New("超时了！")
//	}
//}
//
///*
//脉搏器（ticker）
//我们可以使用尝试发送操作来实现一个每隔一定时间发送一个信号的脉搏器。
//事实上，time标准库包中的Tick函数提供了同样的功能，但效率更高。 我们应该尽量使用标准库包中的实现。
// */
//func Tick(d time.Duration) <-chan struct{} {
//	c := make(chan struct{}, 1) // 容量最好为1
//	go func() {
//		for {
//			time.Sleep(d)
//			select {
//			case c <- struct{}{}:
//			default:
//			}
//		}
//	}()
//	return c
//}
//
//
//func TestMain29(t *testing.T) {
//	t1 := time.Now()
//	for range Tick(time.Second) {
//		fmt.Println(time.Since(t1))
//	}
//}
///*
//速率限制（rate limiting）
//上面已经展示了如何使用尝试发送实现峰值限制。 同样地，我们也可以使用使用尝试机制来实现速率限制，
//但需要前面刚提到的定时器实现的配合。 速率限制常用来限制吞吐和确保在一段时间内的资源使用不会超标。
//
//下面的例子借鉴了官方Go维基中的例子。 在此例中，任何一分钟时段内处理的请求数不会超过200。
// */
//type Request interface{}
//func handle(r Request) {fmt.Println(r.(int))}
//
//const RateLimitPeriod = time.Minute
//const RateLimit = 200 // 任何一分钟内最多处理200个请求
//
//func handleRequests(requests <-chan Request) {
//	quotas := make(chan time.Time, RateLimit)
//
//	go func() {
//		tick := time.NewTicker(RateLimitPeriod / RateLimit)
//		defer tick.Stop()
//		for t := range tick.C {
//			select {
//			case quotas <- t:
//			default:
//			}
//		}
//	}()
//
//	for r := range requests {
//		<-quotas
//		go handle(r)
//	}
//}
//
//
//func TestMain30(t *testing.T) {
//	requests := make(chan Request)
//	go handleRequests(requests)
//	// time.Sleep(time.Minute)
//	for i := 0; ; i++ {requests <- i}
//}
//
///*
//开关
//通道一文提到了向一个nil通道发送数据或者从中接收数据都属于阻塞操作。 利用这一事实，
//我们可以将一个select流程控制中的case操作中涉及的通道设置为不同的值，以使此select流程控制选择执行不同的分支。
//
//下面是另一个乒乓模拟游戏的实现。此实现使用了选择机制。在此例子中，两个case操作中的通道有且只有一个为nil，
//所以只能是不为nil的通道对应的分支被选中。 每个循环步将对调这两个case操作中的通道，从而改变两个分支的可被选中状态。
// */
//
//type Ball uint8
//func Play(playerName string, table chan Ball, serve bool) {
//	var receive, send chan Ball
//	if serve {
//		receive, send = nil, table
//	} else {
//		receive, send = table, nil
//	}
//	var lastValue Ball = 1
//	for {
//		select {
//		case send <- lastValue:
//		case value := <- receive:
//			fmt.Println(playerName, value)
//			value += lastValue
//			if value < lastValue { // 溢出了
//				os.Exit(0)
//			}
//			lastValue = value
//		}
//		receive, send = send, receive // 开关切换
//		time.Sleep(time.Second)
//	}
//}
//
//
//func TestMain31(t *testing.T) {
//	table := make(chan Ball)
//	go Play("A:", table, false)
//	Play("B:", table, true)
//}
//
///*
//控制代码被执行的几率
//我们可以通过在一个select流程控制中使用重复的case操作来增加对应分支中的代码的执行几率。
//
//一个例子：在上面这个例子中，函数f的调用执行几率大致为函数g的两倍。
// */
//func TestMain32(t *testing.T) {
//	foo, bar := make(chan struct{}), make(chan struct{})
//	close(foo); close(bar) // 仅为演示目的
//	x, y := 0.0, 0.0
//	f := func(){x++}
//	g := func(){y++}
//	for i := 0; i < 100000; i++ {
//		select {
//		case <-foo: f()
//		case <-foo: f()
//		case <-bar: g()
//		}
//	}
//	fmt.Println(x/y) // 大致为2
//}

/*
数据生成/搜集/加载
一个数据产生者可能通过以下途径生成数据：
加载一个文件、或者读取一个数据库、或者用爬虫抓取网页数据；
从一个软件或者硬件系统搜集各种数据；
产生一系列随机数；
等等。
这里，我们使用一个随机数产生器做为一个数据产生者的例子。 此数据产生者函数没有输入，只有输出。
一个数据产生者可以在任何时刻关闭返回的通道以结束数据生成。
*/
func RandomGenerator() <-chan uint64 {
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
				break
			}
			c <- binary.BigEndian.Uint64(rnds)
		}
	}()
	return c
}

/*
数据聚合
一个数据聚合模块的工作协程将多个数据流合为一个数据流。 假设数据类型为int64，下面这个函数将任意数量的数据流合为一个。
*/
//func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
//	out := make(chan uint64)
//	for _, in := range inputs {
//		go func(in <-chan uint64) {
//			for {
//				out <- <-in // <=> out <- (<-in)
//			}
//		}(in)
//	}
//	return out
//}
// 一个更完美的实现需要考虑一个输入数据流是否已经关闭。（下面要介绍的其它工作协程同理。）
func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(int <-chan uint64) {
			defer wg.Done()
			// 如果通道in被关闭，此循环将最终结束。
			for x := range in {
				output <- x
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

/*
数据分流
数据分流是数据聚合的逆过程。数据分流的实现很简单，但在实践中用的并不多。
*/

func Divisor(input <-chan uint64, outputs ...chan<- uint64) {
	for _, out := range outputs {
		go func(o chan<- uint64) {
			for {
				o <- <-input // <=> o <- (<-input)
			}
		}(out)
	}
}

/*
数据合成
数据合成将多个数据流中读取的数据合成一个。

下面是一个数据合成工作函数的实现中，从两个不同数据流读取的两个uint64值组成了一个新的uint64值。 当然，在实践中，数据的组合比这复杂得多。
*/
func Composor(inA, inB <-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	go func() {
		for {
			a1, b, a2 := <-inA, <-inB, <-inA
			output <- a1 ^ b&a2
		}
	}()
	return output
}

/*
数据分解
数据分解是数据合成的逆过程。一个数据分解者从一个通道读取一份数据，并将此数据分解为多份数据。 这里就不举例了。
数据复制/增殖
数据复制（增殖）可以看作是特殊的数据分解。一份输入数据将被复制多份并输出给多个数据流。

一个例子
*/
func Duplicator(in <-chan uint64) (<-chan uint64, <-chan uint64) {
	outA, outB := make(chan uint64), make(chan uint64)
	go func() {
		for x := range in {
			outA <- x
			outB <- x
		}
	}()
	return outA, outB
}

/*
数据计算/分析
数据计算和数据分析模块的功能因具体程序不同而有很大的差异。 一般来说，数据分析者接收一份数据并对之加工处理后转换为另一份数据。

下面的简单示例中，每个输入的uint64值将被进行位反转后输出。
*/
func Calculator(in <-chan uint64, out chan uint64) <-chan uint64 {
	if out == nil {
		out = make(chan uint64)
	}
	go func() {
		for x := range in {
			out <- ^x
		}
	}()
	return out
}

/*
数据验证/过滤
一个数据验证或过滤者的任务是检查输入数据的合理性并抛弃不合理的数据。 比如，下面的工作者协程将抛弃所有的非素数。
*/

func Filter0(input <-chan uint64, output chan uint64) <-chan uint64 {
	if output == nil {
		output = make(chan uint64)
	}
	go func() {
		bigInt := big.NewInt(0)
		for x := range input {
			bigInt.SetUint64(x)
			if bigInt.ProbablyPrime(1) {
				output <- x
			}
		}
	}()
	return output
}

func Filter(input <-chan uint64) <-chan uint64 {
	return Filter0(input, nil)
}

/*
数据服务/存盘
一般，一个数据服务或者存盘模块为一个数据流系统中的最后一个模块。 这里的实现值是简单地将数据输出到终端。
*/
func Printer(input <-chan uint64) {
	for x := range input {
		fmt.Println(x)
	}
}

/*
组装数据流系统
现在，让我们使用上面的模块工作者函数实现来组装一些数据流系统。 组装数据流仅仅是创建一些工作者协程函数调用，
并为这些调用指定输入数据流和输出数据流。

数据流系统例子1（一个流线型系统）：
*/

func TestMain34(t *testing.T) {
	Printer(
		Filter(
			Calculator(
				RandomGenerator(), nil,
			),
		),
	)
}

// 数据流系统例子2（一个单向无环图系统）
func TestMain35(t *testing.T) {
	filterA := Filter(RandomGenerator())
	filterB := Filter(RandomGenerator())
	filterC := Filter(RandomGenerator())
	filter := Aggregator(filterA, filterB, filterC)
	calculatorA := Calculator(filter, nil)
	calculatorB := Calculator(filter, nil)
	calculator := Aggregator(calculatorA, calculatorB)
	Printer(calculator)
}

func TestMain33(t *testing.T) {
	bigInt := big.NewInt(0)
	bigInt.SetInt64(-1)
	fmt.Println(bigInt.ProbablyPrime(1))
}

func TestMain36(t *testing.T) {

}

func TestMain37(t *testing.T) {

}

func TestMain38(t *testing.T) {

}
