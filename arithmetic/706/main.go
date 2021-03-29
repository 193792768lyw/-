package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	param_2 := obj.Get(1)
	fmt.Println(param_2)
	param_2 = obj.Get(3)
	fmt.Println(param_2)
	obj.Put(2, 1)
	param_2 = obj.Get(2)
	fmt.Println(param_2)
	obj.Remove(2)
	param_2 = obj.Get(2)
	fmt.Println(param_2)
}

const base = 769

type MyHashMap struct {
	data []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

type Value struct {
	key   int
	value int
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	if !this.Contains(key) {
		h := this.hash(key)
		this.data[h].PushBack(&Value{
			key:   key,
			value: value,
		})

	} else {
		h := this.hash(key)
		for e := this.data[h].Front(); e != nil; e = e.Next() {
			if e.Value.(*Value).key == key {
				e.Value.(*Value).value = value
			}
		}
	}

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(*Value).key == key {
			return e.Value.(*Value).value
		}
	}
	return -1
}

func (s *MyHashMap) hash(key int) int {
	return key % base
}

func (s *MyHashMap) Remove(key int) {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(*Value).key == key {
			s.data[h].Remove(e)
		}
	}
}

func (s *MyHashMap) Contains(key int) bool {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(*Value).key == key {
			return true
		}
	}
	return false
}
