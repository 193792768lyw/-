package main

import "fmt"

func main() {
	//fmt.Println(countBits(5))
	test_recover()
	fmt.Println("after recover")
}

func test_recover() {
	defer func() {
		fmt.Println("defer func")
		switch recover().(type) {
		case string:
			fmt.Println("recover success",recover())
		}

	}()

	panic("crash")
	fmt.Println("after panic")
}

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		temp := 0
		for j := i; j != 0; {
			temp++
			j &= j - 1
		}
		res[i] = temp
	}

	return res
}
//
//对于所有的数字，只有两类：
//奇数：二进制表示中，奇数一定比前面那个偶数多一个 1，因为多的就是最低位的 1。
//举例：
//0 = 0       1 = 1
//2 = 10      3 = 11
//偶数：二进制表示中，偶数中 1 的个数一定和除以 2 之后的那个数一样多。因为最低位是 0，除以 2 就是右移一位，也就是把那个 0 抹掉而已，所以 1 的个数是不变的。
//举例：
//2 = 10       4 = 100       8 = 1000
//3 = 11       6 = 110       12 = 1100


func countBits1(num int) []int {
	result := make([]int, num+1)
	result[0] = 0
	for i := 1; i <= num; i++ {
		if i % 2 == 1{
			result[i] = result[i-1] + 1
		} else {
			result[i] = result[i/2]
		}

	}
	return result
}
