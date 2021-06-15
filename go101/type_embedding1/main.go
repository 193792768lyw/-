package main

import (
	"fmt"
	"reflect"
)

type x string

func (x) M() {}

type A struct {
	x
}

func (A) y(int) bool {
	return false
}

type y struct {
	z byte
}

type B struct {
	y
}

func (B) x(string) {}

func mainc() {
	var v struct {
		A
		B
	}
	//_ = v.x // error: 模棱两可的v.x
	//_ = v.y // error: 模棱两可的v.y
	_ = v.M // ok. <=> v.A.x.M
	_ = v.z // ok. <=> v.B.y.z

}

//
//type A struct {
//	x string
//}
//func (A) y(int) bool {
//	return false
//}
//
//type B struct {
//	y bool
//}
//func (B) x(string) {}
//
//type C struct {
//	B
//}
//
//var v1 struct {
//	A
//	B
//}
//
//func f1() {
//	_ = v1.x // error: 模棱两可的v1.x
//	_ = v1.y // error: 模棱两可的v1.y
//}
//
//var v2 struct {
//	A
//	C
//}
//
//func f2() {
//	fmt.Printf("%T \n", v2.x) // string
//	fmt.Printf("%T \n", v2.y) // func(int) bool
//}
