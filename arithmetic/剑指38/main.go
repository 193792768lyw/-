package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permutation("aab"))
}

func permutation1(s string) []string {

	res := map[string]struct{}{}
	s1 := []byte(s)
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(s1) {
			res[string(s1)] = struct{}{}
			return
		}

		for i := index; i < len(s1); i++ {
			s1[i], s1[index] = s1[index], s1[i]
			dfs(index + 1)
			s1[i], s1[index] = s1[index], s1[i]
		}
	}
	dfs(0)
	result := make([]string, 0)
	for k := range res {
		result = append(result, k)
	}

	return result
}

func permutation(s string) (ans []string) {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	n := len(t)
	perm := make([]byte, 0, n)
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			ans = append(ans, string(perm))
			return
		}
		for j, b := range vis {
			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
				continue
			}
			vis[j] = true
			perm = append(perm, t[j])
			backtrack(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}
	backtrack(0)
	return
}
