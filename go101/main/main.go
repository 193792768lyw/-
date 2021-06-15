package main

import (
	"Study/go101/foo"
	"fmt"
)

type B struct {
	n bool
}

func (b B) m() {
	fmt.Println("B", b.n)
}

type C struct {
	foo.A
	B
}

func main() {

	var dd foo.I
	var c C = C{
		A: foo.A{},
		B: B{true},
	}
	c.m()      // B false
	foo.Bar(c) // A 0
	dd = A1{
		n: 444,
	}
	fmt.Println(dd)
}

type A1 struct {
	n int
}

func (a A1) m() {
	fmt.Println("A", a.n)
}
