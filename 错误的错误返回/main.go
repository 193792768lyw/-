package main

import (
	"fmt"
	"log"
)

func main() {
	//if r := recover(); r != nil {
	//	log.Fatal(r)
	//}
	//
	//panic(123)
	//
	//if r := recover(); r != nil {
	//	log.Fatal(r)
	//}
	//defer func() {
	//	defer func() {
	//		// 无法捕获异常
	//		if r := recover(); r != nil {
	//			fmt.Println(r)
	//		}
	//	}()
	//}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("pppppppppp")
		}
		// 虽然总是返回nil, 但是可以恢复异常状态
	}()

	//警告: 用`nil`为参数抛出异常
	panic(nil)
	//foo()
}

func foo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			if r == "TODO" {
				fmt.Println("0000000000")

			}
			//switch x := r.(type) {
			//case string:
			//	err = errors.New(x)
			//case error:
			//	err = x
			//default:
			//	err = fmt.Errorf("Unknown panic: %v", r)
			//}
		}
	}()

	panic("TODO")
}

func MyRecover() interface{} {
	log.Println("trace...")
	return recover()
}

func returnsError() error {
	var p *MyError = nil
	if true {
		return (*MyError)(&MyError{})
		//p = &MyError{}
	}
	return p // Will always return a non-nil error.
}

type MyError struct {
	error
}

func (e *MyError) Error() string {
	return e.error.Error()
}
