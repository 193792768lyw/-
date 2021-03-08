package main

import "fmt"

func main() {
	arr := Constructor1([]int{})
	fmt.Println(arr)
	//fmt.Println(arr.SumRange(0, 5))
}

type NumArray struct {
	arr       []int
	prefixSum []int
}

func Constructor1(nums []int) NumArray {
	prefixSum := make([]int, len(nums))
	if len(nums) != 0 {
		prefixSum[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			prefixSum[i] = prefixSum[i-1] + nums[i]
		}
	}

	res := NumArray{
		arr:       nums,
		prefixSum: prefixSum,
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	fmt.Println(this.prefixSum)
	if i == 0 {
		return this.prefixSum[j]
	}
	return this.prefixSum[j] - this.prefixSum[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
