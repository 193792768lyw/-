package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]string{"a", "aa", "aaa"}, 2))
}

func topKFrequent(words []string, k int) []string {
	ma := map[string]int{}
	for i := range words {
		ma[words[i]] += 1
	}
	s := make([]*struct {
		key   string
		count int
	}, 0)
	for k, v := range ma {
		s = append(s, &struct {
			key   string
			count int
		}{k, v})
	}

	sort.Slice(s, func(i, j int) bool {
		if s[i].count > s[j].count {
			return true
		} else if s[i].count == s[j].count {
			return s[i].key < s[j].key
		}
		return false
	})
	res := make([]string, 0)
	for j := 0; j < k; j++ {
		res = append(res, s[j].key)
	}
	return res
}

func topKFrequent1(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	uniqueWords := make([]string, 0, len(cnt))
	for w := range cnt {
		uniqueWords = append(uniqueWords, w)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
	})
	return uniqueWords[:k]
}
