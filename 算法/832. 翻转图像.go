package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := map[string]interface{}{
		"kk": 99,
	}
	groupMap1 := map[string]interface{}{}
	for _, va := range []string{"kk", "pp"} {
		groupMap1[va] = a[va]
	}
	vv, _  :=  json.Marshal(groupMap1)
	fmt.Println(string(vv))
	//v := a["kk"].(int)
	//fmt.Println(a["kk1"] == nil)
	//fmt.Println(v)
	// delete(a,"pp")

	//fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}
func flipAndInvertImage(A [][]int) [][]int {

	for _, va := range A {
		reversal(va)
	}

	return A
}
func reversal(arr []int) {
	left, right := 0, len(arr)-1
	for i, _ := range arr {
		arr[i] ^= 1
	}
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
