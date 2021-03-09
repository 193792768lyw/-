package main

import (
	"fmt"
	"sort"
)

func main() {
	v := Constructor()
	v.AddNum(10)
	v.AddNum(13)
	fmt.Println(v.FindMedian())
	v.AddNum(12)
	fmt.Println(v.FindMedian())

}

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left

}

type MedianFinder struct {
	Nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.Nums = append(this.Nums, num)
	for i := 1; i < len(this.Nums); i++ {
		vi := this.Nums[i]
		j := i
		for ; j > 0 && this.Nums[j-1] >= vi; j-- {
			this.Nums[j] = this.Nums[j-1]
		}
		this.Nums[j] = vi
	}
	//index := BinarySearch(this.Nums, num)

	//this.Nums = append(this.Nums, 0) // 切片扩展1个空间
	//copy(this.Nums[index+1:], this.Nums[index:])
	//this.Nums[index] = num
	fmt.Println(this.Nums)

}

func (this *MedianFinder) FindMedian() float64 {
	length := len(this.Nums)
	sort.Ints(this.Nums)
	if length%2 == 0 {
		return float64((float64(this.Nums[length/2] + this.Nums[length/2-1])) / float64(2))
	}
	return float64(this.Nums[length/2])

}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
