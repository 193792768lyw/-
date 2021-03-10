package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate(

		"(1)"))
}

/*

"(1)"
"- (3 + (4 + 5))"
"-2+ 1"
"  30"
"2147483647"
如果是空格 : 跳过
如果是 ( : 直接加入 ops 中，等待与之匹配的 )
如果是 ) : 使用现有的 nums 和 ops 进行计算，直到遇到左边最近的一个左括号为止，计算结果放到 nums
如果是 数字 : 从当前位置开始继续往后取，将整一个连续数字整体取出，加入 nums
如果是 +/- : 需要将操作放入 ops 中。在放入之前先把栈内可以算的都算掉，使用现有的 nums 和 ops 进行计算，直到没有操作或者遇到左括号，计算结果放到 nums

*/

func calculate(s string) int {
	nums := make([]int, 0)
	nums = append(nums, 0)
	ops := make([]uint8, 0)
	var ca func()
	ca = func() {
		if len(nums) == 0 || len(nums) < 2 {
			return
		}
		if len(ops) == 0 {
			return
		}
		nu1 := nums[len(nums)-1]
		nu2 := nums[len(nums)-2]
		nums = nums[:len(nums)-2]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		if op == 43 {
			nums = append(nums, nu1+nu2)
		} else if op == 45 {
			nums = append(nums, nu2-nu1)
		}
	}
	for i := 0; i < len(s); i++ {
		value := s[i]
		if value == ' ' { // 如果是空格 : 跳过
			continue
		} else if value == '(' { // 如果是 ( : 直接加入 ops 中，等待与之匹配的 )
			ops = append(ops, value)
		} else if value == ')' { // 如果是 ) : 使用现有的 nums 和 ops 进行计算，直到遇到左边最近的一个左括号为止，计算结果放到 nums
			// 计算到最近一个左括号为止
			for i := len(ops) - 1; i >= 0; i-- {
				op := ops[len(ops)-1]
				if op != '(' {
					ca()
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
		} else {
			if 48 <= value && value <= 57 {
				u := 0
				j := i
				// 将从 i 位置开始后面的连续数字整体取出，加入 nums
				for j < len(s) && 48 <= s[j] && s[j] <= 57 {
					u = u*10 + (int)(s[j]-'0')
					j++
				}
				nums = append(nums, u)
				i = j - 1
			} else {
				// 如果是 +/- : 需要将操作放入 ops 中。在放入之前先把栈内可以算的都算掉，使用现有的 nums 和 ops 进行计算，直到没有操作或者遇到左括号，计算结果放到 nums

				for len(ops) != 0 && ops[len(ops)-1] != '(' {
					ca()
				}
				ops = append(ops, value)
			}
		}
	}
	for len(ops) != 0 {
		ca()
	}
	return nums[len(nums)-1]
}
