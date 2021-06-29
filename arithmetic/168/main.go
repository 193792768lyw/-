package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//fmt.Println(convertToTitle(28))
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(columnTitle string) int {
	res := 0
	index := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += int(columnTitle[i]-64) * int(math.Pow(float64(26), float64(index)))
		index++
	}
	return res
}

func convertToTitle(columnNumber int) string {

	build := strings.Builder{}

	for columnNumber != 0 {
		columnNumber--
		d := columnNumber % 26
		build.WriteString(string(byte(65 + d)))
		columnNumber = columnNumber / 26
	}

	s := []byte(build.String())
	i := 0
	j := len(s) - 1
	for ; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}

	return string(s)
}
