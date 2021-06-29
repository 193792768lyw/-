package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	dead := map[string]bool{}
	for _, s := range wordList {
		dead[s] = true
	}
	// endWord 不在字典中，所以无法进行转换。
	if !dead[endWord] {
		return 0
	}

	delete(dead, beginWord)

	type pair struct {
		status string
		step   int
	}

	q := []pair{{beginWord, 0}}
	seen := map[string]bool{beginWord: true}

	// 枚举 status 通过一次旋转得到的数字
	get := func(status string) (ret []string) {
		s := []byte(status)
		for i, b := range s {
			for k := 97; k <= 'z'; k++ {
				if byte(k) != b {
					s[i] = byte(k)
					if !seen[string(s)] && dead[string(s)] {
						ret = append(ret, string(s))
						seen[string(s)] = true

					}
					s[i] = b
				}
			}
			s[i] = b
		}
		return
	}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if nxt == endWord {
				return p.step + 1 + 1
			}
			q = append(q, pair{nxt, p.step + 1})
		}
	}
	return 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	distBegin := make([]int, len(wordId))
	for i := range distBegin {
		distBegin[i] = inf
	}
	distBegin[beginId] = 0
	queueBegin := []int{beginId}

	distEnd := make([]int, len(wordId))
	for i := range distEnd {
		distEnd[i] = inf
	}
	distEnd[endId] = 0
	queueEnd := []int{endId}

	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		q := queueBegin
		queueBegin = nil
		for _, v := range q {
			if distEnd[v] < inf {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, w := range graph[v] {
				if distBegin[w] == inf {
					distBegin[w] = distBegin[v] + 1
					queueBegin = append(queueBegin, w)
				}
			}
		}

		q = queueEnd
		queueEnd = nil
		for _, v := range q {
			if distBegin[v] < inf {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, w := range graph[v] {
				if distEnd[w] == inf {
					distEnd[w] = distEnd[v] + 1
					queueEnd = append(queueEnd, w)
				}
			}
		}
	}
	return 0
}
