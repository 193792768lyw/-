package go_chann

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

/*
Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。对一个已经被close过的channel进行接收操作
依然可以接受到之前已经成功发送的数据；如果channel中已经没有数据的话将产生一个零值的数据。

当一个channel被关闭后，再向该channel发送数据将导致panic异常。当一个被关闭的channel中已经发送的数据都被成功接收后，
后续的接收操作将不再阻塞，它们会立即返回一个零值。

试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。
关闭一个channels还会触发一个广播机制，
*/
func TestChan(t *testing.T) {

	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < 10; i++ {
		va, ok := <-ch
		fmt.Println(va, ok)
	}

}

// 试图关闭一个nil值的channel也将导致panic异常。
func TestChanClose(t *testing.T) {
	var a chan int
	if a == nil {
		fmt.Println("chan a is nil")
	}
	close(a)

}

//因为关闭操作只用于断言不再向channel发送新的数据，所以只有在发送者所在的goroutine才会调用close函数，因此对一个只接收的channel调用close
//将是一个编译错误。
/*
任何双向channel向单向channel变量的赋值操作都将导致该隐式转换。这里并没有反向转换的语法：
也就是不能将一个类似chan<- int类型的单向型的channel转换为chan int类型的双向型的channel。
*/
func TestJson(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
	//对一个只接收的channel调用close将是一个编译错误
	//close(in)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestVerifyEmailFormat(t *testing.T) {
	//pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	//
	//reg := regexp.MustCompile(pattern)
	//  fmt.Println(reg.MatchString("193792768"))
	reg, _ := regexp.Compile("@123u.com$")
	req := reg.ReplaceAllString("193792768@qq.com", "@huanle.com")
	fmt.Println(req)
}

func TestDef(t *testing.T) {
	for _, filename := range []int{1, 2, 3, 4} {
		v := filename
		defer func() {
			switch p := recover(); p {
			case "jjjfbd":
				fmt.Println("0000000000000")
			default:
				panic(p)

			}
			fmt.Println(v)
		}() // NOTE: risky; could run out of file descriptors
		// ...process f…

		panic("jjjfbd1")
	}
}

//对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
//无缓存的Channel上的发送操作总在对应的接收操作完成前发生.

var limit = make(chan int, 3)
var work = []func(){
	func() { println("1"); time.Sleep(1 * time.Second) },
	func() { println("2"); time.Sleep(1 * time.Second) },
	func() { println("3"); time.Sleep(1 * time.Second) },
	func() { println("4"); time.Sleep(1 * time.Second) },
	func() { println("5"); time.Sleep(1 * time.Second) },
}

func TestMain1(t *testing.T) {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	//select{}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
	<-make(chan int)
}

// InterfaceStructure 定义了一个interface{}的内部结构
type InterfaceStructure struct {
	pt uintptr // 到值类型的指针
	pv uintptr // 到值内容的指针
}

// asInterfaceStructure 将一个interface{}转换为InterfaceStructure
func asInterfaceStructure(i interface{}) InterfaceStructure {
	return *(*InterfaceStructure)(unsafe.Pointer(&i))
}

func TestInterface(t *testing.T) {
	//var p *int = nil
	var i interface{} = nil
	fmt.Printf("%v %+v is nil %v\n", i, asInterfaceStructure(i), i == nil)

	//fmt.Println("Hello, 世界")
	//var val uint64 = 730
	//var ratio uint64 = 50000
	//res := val * ratio/1e6
	//fmt.Println("res:", res)
	//
	//var val2 int64 =100
	//var ratio2 int = 3
	////res2 := float64(val2) / float64(ratio2)
	//res2 := float64(val2) / float64(ratio2)
	//fmt.Println("res2:", Decimal(res2))
	//a := map[string]interface{}{
	//
	//}
	//b := map[string]interface{}{
	//
	//}
	//fmt.Println("第一遍：")
	//for k, v := range a {
	//	fmt.Print(k + " : ")
	//	fmt.Println(v)
	//}
	//fmt.Println("第二遍：")
	//for k, v := range a {
	//	fmt.Print(k + " : ")
	//	fmt.Println(v)
	//}

	//fmt.Println(CompareTwoMapInterface(a ,b ))
fmt.Println(IsNum("k"))
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func CompareTwoMapInterface(data1 map[string]interface{},
	data2 map[string]interface{}) bool {
	keySlice := make([]string, 0)
	dataSlice1 := make([]interface{}, 0)
	dataSlice2 := make([]interface{}, 0)
	for key, value := range data1 {
		keySlice = append(keySlice, key)
		dataSlice1 = append(dataSlice1, value)
	}
	for _, key := range keySlice {
		if data, ok := data2[key]; ok {
			dataSlice2 = append(dataSlice2, data)
		} else {
			return false
		}
	}
	dataStr1, _ := json.Marshal(dataSlice1)
	dataStr2, _ := json.Marshal(dataSlice2)

	return string(dataStr1) == string(dataStr2)
}


func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}
