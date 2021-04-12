package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{0, 0}))
}

/*
输入：nums = [3,30,34,5,9]
输出："9534330"
*/
func largestNumber(nums []int) string {
	strNum := make([]string, len(nums))
	for i := range nums {
		strNum[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strNum, func(i, j int) bool {
		s1 := strNum[i] + strNum[j]
		s2 := strNum[j] + strNum[i]
		if s1 > s2 {
			return true
		}
		return false
	})
	var by strings.Builder
	for _, s := range strNum {
		by.WriteString(s)
	}
	res := by.String()
	if res[0] == '0' {
		return "0"
	}
	return res
}
