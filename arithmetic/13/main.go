package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func romanToInt(s string) int {
	res := 0
	for _, vs := range valueSymbols {
		for strings.HasPrefix(s, vs.symbol) {
			res += vs.value
			s = s[len(vs.symbol):]
		}
		if len(s) == 0 {
			break
		}
	}
	return res
}

/*
var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
    n := len(s)
    for i := range s {
        value := symbolValues[s[i]]
        if i < n-1 && value < symbolValues[s[i+1]] {
            ans -= value
        } else {
            ans += value
        }
    }
    return
}

*/
