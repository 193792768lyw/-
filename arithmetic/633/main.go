package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(11))
}

/*
输入：c = 5
输出：true
解释：1 * 1 + 2 * 2 = 5
*/
func judgeSquareSum1(c int) bool {
	for a := 0; a*a <= c; a++ {
		rt := math.Sqrt(float64(c - a*a))
		if rt == math.Floor(rt) {
			return true
		}
	}
	return false
}

func judgeSquareSum5(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}

/*
费马平方和定理告诉我们：

一个非负整数 cc 如果能够表示为两个整数的平方和，当且仅当 cc 的所有形如 4k + 34k+3 的质因子的幂均为偶数。

证明请见 这里。

因此我们需要对 cc 进行质因数分解，再判断所有形如 4k + 34k+3 的质因子的幂是否均为偶数即可。

*/

func judgeSquareSum(c int) bool {
	for base := 2; base*base <= c; base++ {
		// 如果不是因子，枚举下一个
		if c%base > 0 {
			continue
		}

		// 计算 base 的幂
		exp := 0
		for ; c%base == 0; c /= base {
			exp++
		}

		// 根据 Sum of two squares theorem 验证
		if base%4 == 3 && exp%2 != 0 {
			return false
		}
	}

	// 例如 11 这样的用例，由于上面的 for 循环里 base * base <= c ，base == 11 的时候不会进入循环体
	// 因此在退出循环以后需要再做一次判断
	return c%4 != 3
}
