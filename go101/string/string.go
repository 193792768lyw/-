package main

import (
	"fmt"
	"time"

	"unicode/utf8"
)

func main() {

	bs := make([]byte, 1<<26)
	s0 := string(bs)
	s1 := string(bs)
	s2 := s1[:]

	// s0、s1和s2是三个相等的字符串。
	// s0的底层字节序列是bs的一个深复制。
	// s1的底层字节序列也是bs的一个深复制。
	// s0和s1底层字节序列为两个不同的字节序列。
	// s2和s1共享同一个底层字节序列。

	startTime := time.Now()
	_ = s0 == s1
	duration := time.Now().Sub(startTime)
	fmt.Println("duration for (s0 == s1):", duration)

	startTime = time.Now()
	_ = s1 == s2
	duration = time.Now().Sub(startTime)
	fmt.Println("duration for (s1 == s2):", duration)
}
func main2() {
	hello := []byte("Hello ")
	world := "world!"

	// helloWorld := append(hello, []byte(world)...) // 正常的语法
	helloWorld := append(hello, world...) // 语法糖
	fmt.Println(string(helloWorld))

	helloWorld2 := make([]byte, len(hello)+len(world))
	copy(helloWorld2, hello)
	// copy(helloWorld2[len(hello):], []byte(world)) // 正常的语法
	copy(helloWorld2[len(hello):], world) // 语法糖
	fmt.Println(string(helloWorld2))
}

func main1() {
	s := "éक्षिaπ囧"
	fmt.Println(utf8.RuneCountInString(s))
	d := []rune(s)
	fmt.Println(d)
	//for i, rn := range s {
	//	fmt.Printf("%2v: 0x%x %v \n", i, rn, string(rn))
	//}
	//fmt.Println(len(s))
	//s := "éक्षिaπ囧"
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("第%v个字节为0x%x\n", i, s[i])
	//}

	//s := "éक्षिaπ囧"
	//// 这里，[]byte(s)不需要深复制底层字节。
	//for i, b := range []byte(s) {
	//	fmt.Printf("The byte at index %v: 0x%x \n", i, b)
	//}
}
