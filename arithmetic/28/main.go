package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "llr"))
}

/*
输入：haystack = "hello", needle = "ll"
输出：2
*/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] != needle[0] {
			continue
		} else {
			index := 0
			for k := i; k < len(haystack) && index < len(needle); {
				if haystack[k] == needle[index] {
					k++
					index++
				} else {
					break
				}
			}
			if index == len(needle) {
				return i
			}
		}
	}
	return -1

}
