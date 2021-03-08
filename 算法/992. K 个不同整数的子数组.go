package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3))

}

//
//func subarraysWithKDistinct(A []int, K int) int {
//	res := 0
//	for i := 0; i < len(A); i++ {
//		mapNum := make(map[int]struct{}, 0)
//		for j := i; j < len(A); j++ {
//			mapNum[A[j]] = struct{}{}
//			if len(mapNum) == K {
//				res++
//			}
//			if len(mapNum) > K {
//				break
//			}
//		}
//	}
//	return res
//}

func subarraysWithKDistinct(A []int, K int) (ans int) {
	return atMostK(A,K) - atMostK(A,K-1)
}

func atMostK(A []int, K int) int {
	N := len(A)
	left, right := 0, 0
	counter := make(map[int]int)
	distinct := 0
	res := 0
	for right < N {
		if counter[A[right]] == 0 {
			distinct += 1
		}

		counter[A[right]] += 1
		for distinct > K {
			counter[A[left]] -= 1
			if counter[A[left]] == 0 {
				distinct -= 1
			}
			left += 1
		}

		res += right-left+1
		right += 1
	}

	return res
}

