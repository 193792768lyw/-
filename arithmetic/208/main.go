package main

// 208. 实现 Trie (前缀树)
func main() {

}

type Trie struct {
	isEnd bool      //该结点是否是一个串的结束
	next  [26]*Trie //字母映射表
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		isEnd: false,
		next:  [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		if node.next[char-97] == nil {
			n := Constructor()
			node.next[char-97] = &n
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
