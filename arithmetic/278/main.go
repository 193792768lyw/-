package main

import "fmt"

func main() {
	type aa *int
	var dd aa = nil
	fmt.Println(dd)
}

//func firstBadVersion(n int) int {
//	left , right := 1 ,n
//	for left <= right {
//		mid  := (left + right) / 2
//		flag :=  isBadVersion(mid)
//		if flag == false {
//			left = mid +1
//		}else{
//			right = mid-1
//		}
//	}
//	return left
//}
