package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("ap"))

	trie.Insert("app")

	fmt.Println(trie.Search("app"))
}

type Trie struct {
	isEnd bool      //该结点是否是一个串的结束
	next  [26]*Trie //字母映射表
}

/** Initialize your data structure here. */
func Constructor() *Trie {
	return &Trie{
		isEnd: false,
		next:  [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		if node.next[char-97] == nil {
			node.next[char-97] = Constructor()
		}
		node = node.next[char-97]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, char := range word {
		node = node.next[char-97]
		if node == nil {
			return false
		}
	}
	return node.isEnd

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, char := range prefix {
		node = node.next[char-97]
		if node == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
