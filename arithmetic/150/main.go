package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// 150. 逆波兰表达式求值
// 波兰式、逆波兰式与表达式求值
// https://blog.csdn.net/linraise/article/details/20459751
func main() {
	// context 测试学习
	ctx, _ := context.WithTimeout(context.Background(), time.Second*80)
	go func(ctx context.Context) {
		ctx, fff := context.WithCancel(ctx)
		go jjj(ctx)
		time.Sleep(3 * time.Second)
		fff()
		fmt.Println("000000000000000")
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("----------")
	}
	// fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func jjj(ctx context.Context) {
	select {
	case <-ctx.Done():

		fmt.Println(errors.New("rpc client: call failed: " + ctx.Err().Error()))

	}

}
func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
func evalRPN(tokens []string) int {

	stack := make([]int, 0)
	for _, nu := range tokens {
		if IsNum(nu) {
			numb, _ := strconv.Atoi(nu)
			stack = append(stack, numb)
		} else {
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch nu {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
