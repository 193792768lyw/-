package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {

	arr := []int{1, -1, -2, -3, -4, -5, 2, 3, 4, 5}
	a := 0
	b := 1
	for a < len(arr) && b < len(arr) {
		for a < len(arr) && arr[a] < 0 {
			a += 2
		}
		for b < len(arr) && arr[b] > 0 {
			b += 2
		}
		if a >= len(arr) || b >= len(arr) {
			break
		}
		arr[a], arr[b] = arr[b], arr[a]

	}
	fmt.Println(arr)

	//A := make(chan bool,1)
	//B := make(chan bool)
	//Exit := make(chan bool)
	//go func() {
	//	s := []int{1, 2, 3, 4}
	//	for i := 0; i < len(s) ; i++  {
	//		if ok := <-A; ok {
	//			fmt.Println("A:", s[i])
	//			B <- true
	//		}
	//	}
	//}()
	//go func() {
	//	defer func() {
	//		close(Exit)
	//	}()
	//	s := []byte{'A', 'B', 'C', 'D'}
	//	for i := 0; i < len(s); i++ {
	//		if ok := <- B; ok {
	//			fmt.Printf("B: %c\n", s[i])
	//			A <- true
	//		}
	//	}
	//
	//}()
	//A <- true
	//<- Exit
}

//func main() {
//	//fmt.Println(findRepeatNumber([]int{3, 1, 2, 3}))
//	fmt.Println(getLength([][]int{{1, 3}, {2, 6}, {11, 12}, {10, 13}}))
//}
//
//func main() {
//
//	intChan := make(chan int, 20)
//	exitChan := make(chan bool, 1)
//
//	go func(intChan chan int){
//		for i:=0; i< 20; i++ {
//
//			intChan<- (i%2 + 1)
//		}
//		close(intChan)
//	}(intChan)
//
//	go func(intChan chan int, exitChan chan bool) {
//		for{
//			v, ok := <-intChan
//			if ok {
//				fmt.Println(v)
//			}else{
//				exitChan<- true
//				close(exitChan)
//				break
//			}
//		}
//	}(intChan, exitChan)
//
//	for {
//		_, ok := <- exitChan
//		if !ok {
//			break
//		}
//	}
//
//}

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

func slidingPuzzle(board [][]int) int {
	const target = "123450"

	s := make([]byte, 0, 6)
	for _, r := range board {
		for _, v := range r {
			s = append(s, '0'+byte(v))
		}
	}
	start := string(s)
	if start == target {
		return 0
	}

	// 枚举 status 通过一次交换操作得到的状态
	get := func(status string) (ret []string) {
		s := []byte(status)
		x := strings.Index(status, "0")
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			ret = append(ret, string(s))
			s[x], s[y] = s[y], s[x]
		}
		return
	}

	type pair struct {
		status string
		step   int
	}
	q := []pair{{start, 0}}
	seen := map[string]bool{start: true}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if !seen[nxt] {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}
	return -1
}

func findRepeatNumber(nums []int) int {
	for _, v := range nums {
		v = int(math.Abs(float64(v)))
		if nums[v] < 0 {
			return v
		}
		nums[v] *= -1
	}
	return 0
}

func getLength(arr [][]int) int {
	if len(arr) == 0 {
		return 0
	}

	sort.Slice(arr, func(i, j int) bool {
		fmt.Println(arr[i][0] < arr[j][0])
		return arr[i][0] < arr[j][0]
	})
	res := 0
	start := arr[0][0]
	end := arr[0][1]
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		if v[0] < end && v[1] > end {
			end = v[1]
			if i == len(arr)-1 {
				res += end - start

			}
			continue
		}
		res += end - start
		start = v[0]
		end = v[1]

	}
	return res

}
