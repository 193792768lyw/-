package main

import (
	"fmt"
	"testing"
)

/*
参考答案及解析：不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作。可以使用 make() 初始化之后再用。
同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。
*/

func TestOne(t *testing.T) {
	//list := new([]int)
	//list = append(list, 1)
	//fmt.Println(list)
}

/*
不能通过编译。append() 的第二个参数不能直接使用 slice，需使用 … 操作符，将一个切片追加到另一个切片上：append(s1,s2…)。或者直接跟上元素，形如：append(s1,1,2,3)。
*/
func TestTwo(t *testing.T) {
	//s1 := []int{1, 2, 3}
	//s2 := []int{4, 5}
	//s1 = append(s1, s2)
	//fmt.Println(s1)
}

/*
参考答案及解析：编译不通过 invalid operation: sm1 == sm2
这道题目考的是结构体的比较，有几个需要注意的地方：
结构体只能比较是否相等，但是不能比较大小。
相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1 就是不同的结构体；
     sn3:= struct {
            name string
            age  int
        }{age:11,name:"qq"}
如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
那什么是可比较的呢，常见的有 bool、数值型、字符、指针、数组等，像切片、map、函数等是不能比较的。 具体可以参考 Go 说明文档。https://golang.org/ref/spec#Comparison_operators
*/
func TestThree(t *testing.T) {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
}

/*
参考答案及解析：编译不通过，cannot use i (type int) as type MyInt1 in assignment。
这道题考的是类型别名与类型定义的区别。
第 5 行代码是基于类型 int 创建了新类型 MyInt1，第 6 行代码是创建了 int 的类型别名 MyInt2，注意类型别名的定义时 = 。所以，第 10 行代码相当于是将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过；而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值。
第 10 行代码的赋值可以使用强制类型转化 var i1 MyInt1 = MyInt1(i).
*/
type MyInt1 int
type MyInt2 = int

func TestFour(t *testing.T) {

	//var i int =0
	//var i1 MyInt1 = i
	//var i2 MyInt2 = i
	//fmt.Println(i1,i2)
}

const (
	x = iota
	_
	y
	z = "vdf"
	k
	p = iota
)

func TestFive(t *testing.T) {
	fmt.Println(x, y, z, k, p)
}

func TestSix(t *testing.T) {
	/*
		A. var x = nil
		B. var x interface{} = nil
		C. var x string = nil
		D. var x error = nil
	*/
}

/*
参考答案及解析：BD。知识点：nil 值。nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。强调下 D 选项的 error 类型，
它是一种内置接口类型，看下方贴出的源码就知道，所以 D 是对的。
*/
