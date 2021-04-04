package main

import (
	"fmt"
	"math"
)

// 781. 森林中的兔子
func main() {
	fmt.Println(numRabbits([]int{0, 3, 2, 0, 3, 3, 4, 2, 4, 3, 2, 4, 4, 3, 0, 1, 3, 4, 4, 3}))
}
func numRabbits(answers []int) int {
	mapAn := map[int]int{}
	for _, v := range answers {
		mapAn[v] += 1
	}
	res := 0
	for k, v := range mapAn {
		if v > k+1 {
			n := int(math.Ceil(float64(v) / float64(k+1)))
			res += n * (k + 1)
		} else {
			res += k + 1
		}
	}

	return res
}
