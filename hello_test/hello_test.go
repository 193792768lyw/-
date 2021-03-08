package hello_test

import (
	"fmt"
	"testing"
	"time"
	"unicode/utf8"
)

/*
在以字符串作为参数传递给fmt.Println函数时，字符串的内容并没有被复制——传递的仅仅是字符串的地址和长度（字符串的结构在
reflect.StringHeader中定义）。在Go语言中，函数参数都是以复制的方式(不支持以引用的方式)传递（比较特殊的是，Go语言
闭包函数对外部变量是以引用的方式使用）。
*/

func TestFunc(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9}
	for _, va :=  range arr{
		// 这里不是匿名函数，所以打印没问题
		//go fmt.Println(va)
		// 底下这个会有问题
		go func() {
			fmt.Println(va)
		}()
	}
	time.Sleep(3 * time.Second)
}



func TestHello(t *testing.T) {
	fmt.Println("你好, 世界!")
	fmt.Println(Inc())
}
func Inc() (v int) {
	defer func(){ v++ } ()
	return 42
}
/*
测试遍历数组时修改其中的值起不起作用
结论：当数组里面是指针或者引用类型时可以修改成功
*/
func TestStructSlice(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	persons := [2]*person{
		{
			name: "刘",
			age:  22,
		},
		{
			name: "yyy",
			age:  222,
		},
	}

	for _, p := range persons {
		p.age += 10
	}
	fmt.Printf("%T\n", persons)
	fmt.Printf("%#v\n", persons)
}

func TestArray(t *testing.T) {
	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针

	fmt.Println(a[0], a[1]) // 打印数组的前2个元素
	fmt.Println(b[0], b[1]) // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	fmt.Printf("--------------------\n")
	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}

}

/*
 管道的同步操作
在这里，我们并不关心管道中传输数据的真实类型，其中管道接收和发送操作只是用于消息的同步。对于这种场景，我们用空数组来作为管道类型可以减少管道
元素赋值时的开销。当然一般更倾向于用无类型的匿名结构体代替：
*/
func TestChanSync(t *testing.T) {
	//c1 := make(chan [0]int)
	//go func() {
	//	fmt.Println("c1")
	//	c1 <- [0]int{}
	//}()
	//<-c1
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
	}()
	<-c2
}

func TestString(t *testing.T) {
	//s := "hello, world"
	//hello := s[:5]
	//world := s[7:]
	//
	//s1 := "hello, world"[:5]
	//s2 := "hello, world"[7:]
	//fmt.Printf(hello,world,s1,s2)
	//fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	//fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	//fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5
	//for i, c := range []byte("世界abc") {
	//	fmt.Println(i, c)
	//}
	//for i, c := range "世界abc"{
	//	fmt.Println(i,c)
	//}
	//fmt.Printf("%#v\n", []rune("世界"))      		// []int32{19990, 30028}
	//fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界

}

func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

var (
	a []int            // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
	b = []int{}        // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
	c = []int{1, 2, 3} // 有3个元素的切片, len和cap都为3
	//       0   1   2
	d = c[:2]             // 有2个元素的切片, len为2, cap为3
	e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
	f = c[:0]             // 有0个元素的切片, len为0, cap为3
	g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
	h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
	i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
)

// 往切片中插入值
func TestSlice1(t *testing.T) {
	var a = []int{1, 2, 3}
	i := 2
	a = append(a, 0)     // 切片扩展1个空间
	copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
	a[i] = 909           // 设置新添加的元素
	fmt.Println(a)

	//var a []int
	//a = append(a, 1)               // 追加1个元素
	//a = append(a, 1, 2, 3)         // 追加多个元素, 手写解包方式
	//a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
	//var a = []int{1,2,3}
	//a = append([]int{0}, a...)        // 在开头添加1个元素
	//a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片

}
func TestInsertSlice2(t *testing.T) {
	var a = []int{1, 2, 3}
	i := 2

	//a = append(a[:i], append([]int{44}, a[i:]...)...)       // 在第i个位置插入x
	a = append(a[:i], append([]int{1, 2, 3}, a[i:]...)...) // 在第i个位置插入切片

	fmt.Println(a)
}

// 用copy和append组合也可以实现在中间位置插入多个元素(也就是插入一个切片):
func TestInsertSlice3(t *testing.T) {
	var a = []int{1, 2, 3}
	i := 2
	x := []int{200, 200}
	a = append(a, x...)       // 为x切片扩展足够的空间
	copy(a[i+len(x):], a[i:]) // a[i:]向后移动len(x)个位置
	copy(a[i:], x)            // 复制新添加的切片

	fmt.Println(a)
}

//  删除切片元素
// 根据要删除元素的位置有三种情况：从开头位置删除，从中间位置删除，从尾部删除。其中删除切片尾部的元素最快
func TestDelSlice(t *testing.T) {
	// 尾部删除
	//a = []int{1, 2, 3}
	//a = a[:len(a)-1]   // 删除尾部1个元素
	//a = a[:len(a)-N]   // 删除尾部N个元素
	//// 删除开头的元素
	//// 删除开头的元素可以直接移动数据指针：
	//a = []int{1, 2, 3}
	//a = a[1:] // 删除开头1个元素
	//a = a[N:] // 删除开头N个元素
	//// 删除开头的元素也可以不移动数据指针，但是将后面的数据向开头移动。可以用append原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）：
	//a = []int{1, 2, 3}
	//a = append(a[:0], a[1:]...) // 删除开头1个元素
	//a = append(a[:0], a[N:]...) // 删除开头N个元素
	//a = []int{1, 2, 3}
	//a = a[:copy(a, a[1:])] // 删除开头1个元素
	////a = a[:copy(a, a[N:])] // 删除开头N个元素
	//
	//fmt.Println(a)

	/*
	切片就是一种简化版的动态数组。因为动态数组的长度是不固定，切片的长度自然也就不能是类型的组成部分了。
	对于类型，和数组的最大不同是，切片的类型和长度信息无关，只要是相同类型元素构成的切片均对应相同的切片类型。
		 切片的定义reflect.SliceHeader：
	 type SliceHeader struct {
		Data uintptr
		Len  int
		Cap  int
	}
	 */
	var b []int // b是一个变量,在不断往切片中添加元素的过程中，b变量本身不边，但是其存储的值即 指向（*reflect.SliceHeader）的指针却不断在

	// 变化
	fmt.Printf("%p\n",b)
	//fmt.Println( (*reflect.SliceHeader)(unsafe.Pointer(&b)))
	for i := 0; i < 30 ; i++ {
		b = append(b, i)
		//fmt.Println( (*reflect.SliceHeader)(unsafe.Pointer(&b)))
		fmt.Printf("%p\n",b)
	}
//fmt.Println(string(TrimSpace([]byte("wnvsd uuu p"))))
}


func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}
func TestGrou(t *testing.T) {

}