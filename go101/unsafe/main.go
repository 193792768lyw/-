package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main1() {
	var x struct {
		a int64
		b bool
		c string
	}
	const M, N = unsafe.Sizeof(x.c), unsafe.Sizeof(x)
	fmt.Println(M, N) // 16 32

	fmt.Println(unsafe.Alignof(x.a)) // 8
	fmt.Println(unsafe.Alignof(x.b)) // 1
	fmt.Println(unsafe.Alignof(x.c)) // 8

	fmt.Println(unsafe.Offsetof(x.a)) // 0
	fmt.Println(unsafe.Offsetof(x.b)) // 8
	fmt.Println(unsafe.Offsetof(x.c)) // 16
}

func main2() {
	type T struct {
		c string
	}
	type S struct {
		b bool
	}
	var x struct {
		a int64
		*S
		T
	}

	fmt.Println(unsafe.Offsetof(x.a)) // 0

	fmt.Println(unsafe.Offsetof(x.S)) // 8
	fmt.Println(unsafe.Offsetof(x.T)) // 16

	// 此行可以编译过，因为选择器x.c中的隐含字段T为非指针。
	fmt.Println(unsafe.Offsetof(x.c)) // 16

	// 此行编译不过，因为选择器x.b中的隐含字段S为指针。
	//fmt.Println(unsafe.Offsetof(x.b)) // error

	// 此行可以编译过，但是它将打印出字段b在x.S中的偏移量.
	fmt.Println(unsafe.Offsetof(x.S.b)) // 0
}

//
//func main() {
//	a := [16]int{3: 3, 9: 9, 11: 11}
//	fmt.Println(a)
//	eleSize := int(unsafe.Sizeof(a[0]))
//	p9 := &a[9]
//	up9 := unsafe.Pointer(p9)
//	p3 := (*int)(unsafe.Add(up9, unsafe.IntegerType(-6*eleSize)))
//	fmt.Println(*p3) // 3
//	s := unsafe.Slice((*unsafe.ArbitraryType)(p9), 5)[:3]
//	fmt.Println(s)              // [9 0 11]
//	fmt.Println(len(s), cap(s)) // 3 5
//
//	// 下面是两个不正确的调用。因为它们
//	// 的返回结果引用了未知的内存块。
//	_ = unsafe.Add(up9, unsafe.IntegerType(7*eleSize))
//	_ = unsafe.Slice((*unsafe.ArbitraryType)(p9), 8)
//}

// 假设此函数不会被内联（inline）。
func createInt() *int {
	return new(int)
}

func foo() {
	p0, y, z := createInt(), createInt(), createInt()
	var p1 = unsafe.Pointer(y) // 和y一样引用着同一个值
	var p2 = uintptr(unsafe.Pointer(z))

	// 此时，即使z指针值所引用的int值的地址仍旧存储
	// 在p2值中，但是此int值已经不再被使用了，所以垃圾
	// 回收器认为可以回收它所占据的内存块了。另一方面，
	// p0和p1各自所引用的int值仍旧将在下面被使用。

	// uintptr值可以参与算术运算。
	p2 += 2
	p2--
	p2--

	*p0 = 1                         // okay
	*(*int)(p1) = 2                 // okay
	*(*int)(unsafe.Pointer(p2)) = 3 // 危险操作！
}

func maina() {
	x := 123                // 类型为int
	p := unsafe.Pointer(&x) // 类型为unsafe.Pointer
	pp := &p                // 类型为*unsafe.Pointer
	p = unsafe.Pointer(pp)
	pp = (*unsafe.Pointer)(p)
}

func mainb() {
	//type MyString string
	//ms := []MyString{"C", "C++", "Go"}
	//fmt.Printf("%s\n", ms)  // [C C++ Go]
	//// ss := ([]string)(ms) // 编译错误
	//ss := *(*[]string)(unsafe.Pointer(&ms))
	//ss[1] = "Rust"
	//fmt.Printf("%s\n", ms) // [C Rust Go]
	//// ms = []MyString(ss) // 编译错误
	//ms = *(*[]MyString)(unsafe.Pointer(&ss))

	fmt.Println(len("aaa我"))
	fmt.Println(String2ByteSlice("aaa我"))
}
func String2ByteSlice(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s)) // 危险
}

func mainc() {
	type T struct{ a int }
	var t T
	fmt.Printf("%p\n", &t)                          // 0xc6233120a8
	println(&t)                                     // 0xc6233120a8
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t))) // c6233120a8
	fmt.Println(uintptr(unsafe.Pointer(&t)))
}

type T struct {
	x bool
	y [3]int16
}

const N = unsafe.Offsetof(T{}.y)
const M = unsafe.Sizeof(T{}.y[0])

func maind() {
	//t := T{y: [3]int16{123, 456, 789}}
	//p := unsafe.Pointer(&t)
	//// "uintptr(p) + N + M + M"为t.y[2]的内存地址。
	//ty2 := (*int16)(unsafe.Pointer(uintptr(p)+N+M+M))
	//fmt.Println(*ty2) // 789
	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
	fmt.Println(p)
}

func maine() {
	a := [...]byte{'G', 'o', 'l', 'a', 'n', 'g'}
	s := "Java"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&a))
	hdr.Len = len(a)
	fmt.Println(s) // Golang
	// 现在，字符串s和切片a共享着底层的byte字节序列，
	// 从而使得此字符串中的字节变得可以修改。
	a[2], a[3], a[4], a[5] = 'o', 'g', 'l', 'e'
	fmt.Println(s) // Google
}

func mainf() {
	a := [6]byte{'G', 'o', '1', '0', '1'}
	bs := []byte("Golang")
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	hdr.Data = uintptr(unsafe.Pointer(&a))

	hdr.Len = 2
	hdr.Cap = len(a)
	fmt.Printf("%s\n", bs) // Go
	bs = bs[:cap(bs)]
	fmt.Println(string(bs))
}

func String2ByteSlice1(str string) (bs []byte) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sliceHdr.Data = strHdr.Data
	sliceHdr.Cap = strHdr.Len
	sliceHdr.Len = strHdr.Len
	return
}

func main() {
	// str := "Golang"
	// 对于官方标准编译器来说，上面这行将使str中的字节
	// 开辟在不可修改内存区。所以这里我们使用下面这行。
	str := strings.Join([]string{"Go", "land"}, "")
	s := String2ByteSlice1(str)
	fmt.Printf("%s\n", s) // Goland
	s[5] = 'g'
	fmt.Println(str) // Golang
	var ddd uintptr = 99
	ddd += 1
	fmt.Println(ddd)
}
