package main

import "fmt"

func main() {

	var f = func(va interface{}) int64 {
		switch va.(type) {
		case int, int8, int16, int32, int64, uint, uint8:
			return va.(int64)
		case nil:
			return 0
		}
		return 0
	}

	var f1 interface{}

	fmt.Println(f(f1))

}
func findMaxLength(nums []int) (maxLength int) {
	mp := map[int]int{0: -1}
	counter := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if prevIndex, has := mp[counter]; has {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
