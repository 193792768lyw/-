package main

import "fmt"

func main() {
	fmt.Println(calculate("3+5 / 2"))
}

//"3+2*2"
func calculate(s string) int {
	// 使用 map 维护一个运算符优先级
	// 这里的优先级划分按照「数学」进行划分即可
	map1 := map[uint8]int{'-': 1,
		'+': 1,
		'*': 2,
		'/': 2,
		'%': 2,
		'^': 3,
	}
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
		b := nums[len(nums)-1]
		a := nums[len(nums)-2]
		nums = nums[:len(nums)-2]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		ans := 0
		if op == '+' {
			ans = a + b
		} else if op == '-' {
			ans = a - b
		} else if op == '*' {
			ans = a * b
		} else if op == '/' {
			ans = a / b
		} else if op == '%' {
			ans = a % b
		}
		nums = append(nums, ans)
	}
	for i := 0; i < len(s); i++ {
		value := s[i]
		if value == ' ' { // 如果是空格 : 跳过
			continue
		} else if value == '(' { // 如果是 ( : 直接加入 ops 中，等待与之匹配的 )
			ops = append(ops, value)
			if s[i+1] == '-' || s[i+1] == '+' {
				nums = append(nums, 0)
			}
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
				// 有一个新操作要入栈时，先把栈内可以算的都算了
				// 但是和只有 +/- 的情况不同的是，只有比当前的「运算符」优先级高，才能进行运算
				// 如果是 +/- : 需要将操作放入 ops 中。在放入之前先把栈内可以算的都算掉，使用现有的 nums 和 ops 进行计算，直到没有操作或者遇到左括号，计算结果放到 nums
				for len(ops) != 0 && ops[len(ops)-1] != '(' {
					if map1[ops[len(ops)-1]] >= map1[value] {
						ca()
					} else {
						break
					}
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
