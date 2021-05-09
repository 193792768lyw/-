package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(strings.Index("hhhh","y"))
	GuessingGame()
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func minDays(bloomDay []int, m, k int) int {
	if m > len(bloomDay)/k {
		return -1
	}
	minDay, maxDay := math.MaxInt32, 0
	for _, day := range bloomDay {
		if day < minDay {
			minDay = day
		}
		if day > maxDay {
			maxDay = day
		}
	}
	return minDay + sort.Search(maxDay-minDay, func(days int) bool {
		days += minDay
		flowers, bouquets := 0, 0
		for _, d := range bloomDay {
			if d > days {
				flowers = 0
			} else {
				flowers++
				if flowers == k {
					bouquets++
					flowers = 0
				}
			}
		}
		return bouquets >= m
	})
}
