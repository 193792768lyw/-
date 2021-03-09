package main

import (
	"bytes"
	"fmt"
)

// 1047. 删除字符串中的所有相邻重复项
func main() {

	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(s string) string {
	stack := []byte{}
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func removeDuplicates1(S string) string {
	var s func(str string) string
	s = func(str string) string {
		var sTemp bytes.Buffer
		for i := 0; i < len(str); {
			if i == len(str)-1 || str[i] != str[i+1] {
				sTemp.WriteByte(str[i])
				i++
			} else {
				i += 2
			}
		}
		if len(sTemp.String()) == len(str) {
			return str
		}
		return s(sTemp.String())
	}

	return s(S)
}
