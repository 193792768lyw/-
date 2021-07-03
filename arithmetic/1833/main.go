package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxIceCream([]int{10, 6, 8, 7, 7, 8}, 5))
}

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	res := 0
	for _, v := range costs {
		if coins >= v {
			res += 1
			coins -= v
		}
	}
	return res
}
