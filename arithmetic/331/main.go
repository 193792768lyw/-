package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("9,#,#,1"))
}

func isValidSerialization(preorder string) bool {
	stack := make([]string, 0)
	str := strings.Split(preorder, ",")
	for i := 0; i < len(str); i++ {
		stack = append(stack, str[i])
		for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" && stack[len(stack)-3] != "#" {
			stack = stack[:len(stack)-3]
			stack = append(stack, "#")
		}

	}
	return len(stack) == 1 && stack[0] == "#"
}
