package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3}, 4))
}
func findMaxAverage1(nums []int, k int) float64 {

	maxSum := math.MinInt64
	curSum := 0
	for i, num := range nums {
		curSum += num
		if i >= k {
			curSum -= nums[i-k]
		}
		if (i+1)>=k {
			if curSum > maxSum {
				maxSum = curSum
			}
		}
	}
return float64(maxSum) / float64(k)
}


func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
