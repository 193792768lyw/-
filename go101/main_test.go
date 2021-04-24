package main

import (
	"fmt"
	"testing"
)

/*
uintptr、int以及uint类型的值的尺寸依赖于具体编译器实现。 通常地，在64位的架构上，int和uint类型的值是64位的；在32位的架构上，它们是32位的。
编译器必须保证uintptr类型的值的尺寸能够存下任意一个内存地址。

Go中大多数的类型不确定值都属于字面常量和本文即将介绍的有名常量。
少数类型不确定值包括刚提到的nil和以后会逐步解触到的某些操作的布尔返回值。
*/
func TestMain1(t *testing.T) {
	const N = 123

	var y float32

	y = N // ok: N被隐式转换为类型float32
	fmt.Println(y)
}

func TestMain2(t1 *testing.T) {
	const (
		k = 3 // 在此处，iota == 0

		m float32 = iota + .5 // m float32 = 1 + .5
		n                     // n float32 = 2 + .5

		p    = 9          // 在此处，iota == 3
		q    = iota * 2   // q = 4 * 2
		_                 // _ = 5 * 2
		r                 // r = 6 * 2
		s, t = iota, iota // s, t = 7, 7
		u, v              // u, v = 8, 8
		_, w              // _, w = 9, 9
	)

	const x = iota // x = 0 （iota == 0）
	const (
		y = iota // y = 0 （iota == 0）
		z        // z = 1
	)

	println(m)             // +1.500000e+000
	println(n)             // +2.500000e+000
	println(q, r)          // 8 12
	println(s, t, u, v, w) // 7 7 8 8 9
	println(x, y, z)       // 0 0 1

}

func TestMain3(t1 *testing.T) {
	var (
		c, _ int16 = 15, -6
		e    uint8 = 7
	)

	//// 这些行编译没问题。
	//_ = 12 + 'A' // 两个类型不确定操作数（都为数值类型）
	//_ = 12 - a   // 12将被当做a的类型（float32）使用。
	//_ = a * b    // 两个同类型的类型确定操作数。
	//_ = c % d
	//_, _ = c + int16(e), uint8(c) + e
	//_, _, _, _ = a / b, c / d, -100 / -9, 1.23 / 1.2
	//_, _, _, _ = c | d, c & d, c ^ d, c &^ d
	//_, _, _, _ = d << e, 123 >> e, e >> 3, 0xF << 0
	//_, _, _, _ = -b, +c, ^e, ^-1
	//
	//// 这些行编译将失败。
	//_ = a % b   // error: a和b都不是整数
	//_ = a | b   // error: a和b都不是整数
	//_ = c + e   // error: c和e的类型不匹配
	//_ = b >> 5  // error: b不是一个整数
	//_ = c >> -5 // error: -5不是一个无符号整数

	_ = e << uint(c) // 编译没问题
	_ = e << c       // 从Go 1.13开始，此行才能编译过
	//_ = e << -c      // 从Go 1.13开始，此行才能编译过。
	// 将在运行时刻造成恐慌。
	//_ = e << -1      // error: 右操作数不能为负（常数）
}

type name interface {
	get() string
}

type a struct {
}

func (a1 a) get() string {
	return "a6666"
}

func pr(na *name) {
	fmt.Println((*na).get())
}
func TestMain4(t1 *testing.T) {
	ff := a{}
	var dd name = ff
	pr(&dd)

}

/*
输入：nums = [1,2,3]
输出：[1,2]
解释：[1,3] 也会被视为正确答案。
*/

func TestMain6(t1 *testing.T) {

}
