package main

import "fmt"

// 88. 合并两个有序数组
func main() {
	arr := []int{1}
	merge(arr, 1, []int{}, 0)
	fmt.Println(arr)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	m -= 1
	n -= 1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[index] = nums1[m]
			m--
		} else {
			nums1[index] = nums2[n]
			n--
		}
		index--
	}
	if n >= 0 {
		for n >= 0 {

			nums1[index] = nums2[n]
			index--
			n--
		}
	}
}
