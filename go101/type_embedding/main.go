package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type A struct {
	x int
}

func (a A) MethodA() {}

type B struct {
	A
}
type C struct {
	B
}

func main() {
	var c C

	// 下面的赋值语句中的选择器都是相互等价的。
	_ = c.B.A.x
	_ = c.B.x
	_ = c.A.x
	_ = c.x // x被称为类型C的一个提升字段

	// 这下面的几个方法调用都是互相等价的。
	c.B.A.MethodA()
	c.B.MethodA()
	c.A.MethodA()
	c.MethodA()
}

func main6() {
	t := reflect.TypeOf(Singer{}) // the Singer type
	fmt.Println(t, "has", t.NumField(), "fields:")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
	}
	fmt.Println(t, "has", t.NumMethod(), "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	}

	pt := reflect.TypeOf(&Singer{}) // the *Singer type
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}
func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person // 通过内嵌Person类型来扩展之
	works  []string
}

func main5() {

	var gaga = Singer{Person: Person{"Gaga", 30}}
	gaga.PrintName() // Name: Gaga

	gaga.Name = "Lady Gaga"
	(&gaga).SetAge(31)
	(&gaga).PrintName()   // Name: Lady Gaga
	fmt.Println(gaga.Age) // 31
}

func main3() {
	//type Encoder interface {Encode([]byte) []byte}
	//type Person struct {name string; age int}
	//type Alias = struct {name string; age int}
	//type AliasPtr = *struct {name string; age int}
	//type IntPtr *int
	//type AliasPP = *IntPtr
	//
	//// 这些类型或别名都可以被内嵌。
	//Encoder
	//Person
	//*Person
	//Alias
	//*Alias
	//AliasPtr
	//int
	//*int
	//
	//// 这些类型或别名都不能被内嵌。
	//AliasPP          // 基类型为一个指针类型
	//*Encoder         // 基类型为一个接口类型
	//*AliasPtr        // 基类型为一个指针类型
	//IntPtr           // 定义的指针类型
	//*IntPtr          // 基类型为一个指针类型
	//*chan int        // 基类型为一个非定义类型
	//struct {age int} // 非定义非指针类型
	//map[string]int   // 非定义非指针类型
	//[]int64          // 非定义非指针类型
	//func()           // 非定义非指针类型
}

func main1() {
	type P = *bool
	type M = map[int]int
	var x struct {
		string      // 一个定义的非指针类型
		error       // 一个定义的接口类型
		*int        // 一个非定义指针类型
		P           // 一个非定义指针类型的别名
		M           // 一个非定义类型的别名
		http.Header // 一个定义的映射类型
	}
	x.string = "Go"
	x.error = nil
	x.int = new(int)
	x.P = new(bool)
	x.M = make(M)
	x.Header = http.Header{}

}
