package main

import "fmt"

func main() {
	fmt.Println(findNumOfValidWords([]string{"apple", "pleas", "please"},
		[]string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"}))
	/*
		[0,1,3,2,0]
	*/
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	type wordInfo struct {
		w string
		n map[int32]struct{}
	}

	firstMap := map[uint8][]wordInfo{}
	puzzleMap := make([]map[int32]struct{}, len(puzzles))
	for i, puzzle := range puzzles {
		firstMap[puzzle[0]] = nil
		s := make(map[int32]struct{})
		for _, v := range puzzle {
			s[v] = struct{}{}
		}
		puzzleMap[i] = s
	}

	wordMap := make([]wordInfo, len(words))
	for i, word := range words {
		in := wordInfo{
			w: word,
			n: map[int32]struct{}{},
		}
		for _, v := range word {
			in.n[v] = struct{}{}
		}
		wordMap[i] = in
	}

	for key, _ := range firstMap {
		value := make([]wordInfo, 0)
		for _, v := range wordMap {
			if _, ok := v.n[int32(key)]; ok {
				value = append(value, v)
			}
		}
		firstMap[key] = value
	}

	res := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		words := firstMap[puzzle[0]]
		nums := 0
		for _, word := range words {
			flag := false
			for key, _ := range word.n {
				if _, ok := puzzleMap[i][key]; ok {
					continue
				} else {
					flag = true
					break
				}
			}
			if !flag {
				nums++
			}
		}
		res[i] = nums
	}
	return res
}
