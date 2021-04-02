package main

import "fmt"

func main() {
	//fmt.Println(trap([]int{0}))
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

/*
核心精神：总体积减去柱子体积就是水的容量

利用左右指针的下标差值计算出每一层雨水和柱子的体积。如下图，第一层体积为11，第二层为8，第三层为1。
累加得到整体体积volume = 20（注意：每一层从左边第一个方格到右边最后一个方格之间一定是被蓝黑两种
颜色的方格填满，不会存在空白，所以我们可以这么求值）
计算柱子的总体积Sum，就是height：[0,1,0,2,1,0,1,3,2,1,2,1]数组之和14
返回结果volume − Sum就是雨水的体积。

 */
func trap(height []int) int {
	res := 0
	left := 0
	right := len(height) - 1
	h := 1
	for left <= right {
		for left <= right && height[left] < h   {
			left++
		}
		for   left <= right && height[right] < h {
			right--
		}
		res += right - left + 1
		h++
	}
	for _, v := range height {
		res -= v
	}

	return res
}
