package main

import "fmt"

func main() {
	fmt.Println(clumsy(4))
}

/*
输入：10
输出：12
解释：12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1

根据题意，「笨阶乘」没有显式括号，运算优先级是先「乘除」后「加减」。我们可以从 N 开始，枚举 N−1、N-2直到 1 ，枚举这些数的时候，
认为它们之前的操作符按照「乘」「除」「加」「减」交替进行。

出现乘法、除法的时候可以把栈顶元素取出，与当前的 N 进行乘法运算、除法运算（除法运算需要注意先后顺序），并将运算结果重新压入栈中；

出现加法、减法的时候，把减法视为加上一个数的相反数，然后压入栈，等待以后遇见「乘」「除」法的时候取出。

最后将栈中元素累加即为答案。由于加法运算交换律成立，可以将栈里的元素依次出栈相加。
*/
func clumsy(N int) (ans int) {
	stk := []int{N}
	N--

	index := 0 // 用于控制乘、除、加、减
	for N > 0 {
		switch index % 4 {
		case 0:
			stk[len(stk)-1] *= N
		case 1:
			stk[len(stk)-1] /= N
		case 2:
			stk[len(stk)-1] += N
		default:
			stk = append(stk, -N)
		}
		N--
		index++
	}

	// 累加栈中数字
	for _, v := range stk {
		ans += v
	}
	return
}
