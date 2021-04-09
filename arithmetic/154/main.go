package main

import "fmt"

type name struct {
	b  int
	ma map[int]int
}

/*
测试结构体内部引用的数据传递时是否会完全copy
*/
func main() {
	bb := name{4, map[int]int{}}
	ch(bb)
	fmt.Println(bb)
	//fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

func ch(name2 interface{}) {
	v := name2.(name)
	v.ma[9] = 99
	v.b = 88
}

// 154. 寻找旋转排序数组中的最小值 II
func findMin(nums []int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if res > nums[i] {
			return nums[i]
		}
	}
	return res
}
