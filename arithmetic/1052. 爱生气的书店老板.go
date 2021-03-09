package main

import "fmt"

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}
func maxSatisfied(customers []int, grumpy []int, X int) int {
	res := 0
	for i, value := range grumpy {
		if value == 0 {
			res += customers[i]
		}
	}
	left, right := 0, 0
	temp := 0
	r := 0
	for right < len(grumpy) {
		if grumpy[right] == 1 {
			temp += customers[right]
		}
		if right-left+1 > X {
			if grumpy[left] == 1 {
				temp -= customers[left]
			}
			left++
		}
		if temp > r {
			r = temp
		}
		right++
	}

	return r + res
}
