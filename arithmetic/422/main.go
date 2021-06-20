package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findDuplicates([]int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6}))
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		v := int(math.Abs(float64(nums[i])))

		if nums[v-1] < 0 {
			res = append(res, v)
			continue
		}
		nums[v-1] *= -1
	}

	return res

}
