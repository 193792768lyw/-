package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSquares(12))
}

/*
输入：n = 12
输出：3
解释：12 = 4 + 4 + 4
*/

//func numSquares(n int) int {
//	nu := int(math.Floor(math.Sqrt(float64(n))))
//	res := 0
//	for n != 0 {
//		for n >= nu * nu  {
//			n -= nu*nu
//			res++
//		}
//		nu --
//	}
//	return res
//}

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := i
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
