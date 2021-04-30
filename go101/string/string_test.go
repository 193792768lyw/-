package main

import "fmt"

func main() {
	const World = "world"
	var hello = "hello"

	// 衔接字符串。
	var helloWorld = hello + " " + World
	helloWorld += "!"
	fmt.Println(helloWorld) // hello world!

	// 比较字符串。
	fmt.Println(hello == "hello")   // true
	fmt.Println(hello > helloWorld) // false
}
