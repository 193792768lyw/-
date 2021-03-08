package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
使用
TypeOf

获取intaface的类型

ValueOf

获取interface的具体值
*/

func TestInterface(t *testing.T) {
	var num int = 233
	var i interface{} = num
	fmt.Println("type: ", reflect.TypeOf(i))
	fmt.Println("value: ", reflect.ValueOf(i))

	//fmt.Println("value: ", reflect.ValueOf(i).Elem().CanSet())

}

// 类型转换
func TestTypeChane(t *testing.T) {

	var i int = 233

	p := reflect.ValueOf(&i)
	v := reflect.ValueOf(i)
	fmt.Println(p, v)
	fmt.Println("-------------")
	cp := p.Interface().(*int)
	cv := v.Interface().(int)

	switch p.Interface().(type) {
	case *int:
		fmt.Println("*int")
	case int:
		fmt.Println("int")
	}

	fmt.Println(cp, cv)
}

// output:
// type:  int
// value:  233
type t1 struct {
	A string
	B int
	//c string
}

func (t t1) Fn(jj int) {
	fmt.Println("this is t1's method", jj)
}
func (t t1) fn() {
	fmt.Println("this is  kkkkkkkkk t1's method")
}

// 复合类型获取字段类型，方法
func TestStruct(tr *testing.T) {
	st := t1{
		A: "ss",
		B: 233,
	}

	var t interface{} = st

	rt := reflect.TypeOf(t)
	fmt.Println("t's type", rt.Name())

	rv := reflect.ValueOf(t)
	fmt.Println("t's value", rv)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i).Interface()
		fmt.Printf("field's name: %s, type: %v, value: %v\n", field.Name, field.Type, value)
	}

	// 反射似乎不检测非导出方法
	for i := 0; i < rv.NumMethod(); i++ {
		m := rt.Method(i)
		fmt.Printf("method'name: %s, type: %v\n", m.Name, m.Type)
	}

}

// output
// t's type t1
// t's value {ss 233}
// field's name: A, type: string, value: ss
// field's name: B, type: int, value: 233
// method'name: Fn, type: func(main.t1)

//修改原来的值
func TestFuncNumMethod(t1 *testing.T) {
	var i int = 233

	// 指针类型才可以修改值
	p := reflect.ValueOf(&i)

	fmt.Println(p.Type())
	fmt.Println(reflect.TypeOf(&i).Elem())
	fmt.Println(reflect.TypeOf(&i))
	// 获取原始反射对象
	pe := p.Elem()

	fmt.Println("type of pe: ", pe.Type())
	fmt.Println("can set?", pe.CanSet())

	pe.SetInt(23)
	fmt.Println("after change", i)
}

// 动态调用方法

func TestFunCall(te *testing.T) {

	st := t1{
		A: "ss",
		B: 233,
	}

	var t interface{} = st

	rv := reflect.ValueOf(t)

	method1 := rv.MethodByName("Fn")
	// Fn方法不需要参数，这里传len为0的[]reflect.Value既可
	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf(9999))
	method1.Call(args)

}

type IT interface {
	test1()
}

type T struct {
	A string
}

func (t *T) test1() {}

func TestImpleFunc(tk *testing.T) {
	t := &T{}
	ITF := reflect.TypeOf((*IT)(nil)).Elem()
	tv := reflect.TypeOf(t)
	fmt.Println(ITF, tv)
	fmt.Println(tv.Implements(ITF))
	fmt.Println(ITF,tv)
}

type Nexter interface {
	Next() Nexter
}

type Node struct {
	next Nexter
}

func (n *Node) Next() Nexter {
	return nil
}

func TestMainNil(t *testing.T) {
	var p Nexter = nil
	//var p Nexter = (*Node)(nil)

	//if p != nil {
	//	n := p.(*Node) // will not fail IF p really contains a value of type *Node
	//	fmt.Println(n)
	//}

	// This will never fail:
	if n, ok := p.(*Node); ok {
		fmt.Printf("n=%#v", n)
	}

	//var n *Node
	//fmt.Println(n == nil) // will print true
	//n = p.(*Node) // will fail
	//
	//n = (*Node)(nil)
}
