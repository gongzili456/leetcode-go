// https://leetcode-cn.com/problems/lru-cache/solution/lru-huan-cun-ji-zhi-by-leetcode/
// LRU缓存机制
package main

import "fmt"

func main() {
	obj := Constructor(2)
	param1 := obj.Get(1)
	obj.Put(1, 1)

	fmt.Println("param1: ", param1)
}

// LinkNode is
type LinkNode struct {
	Key   int
	Value int
	Prev  *LinkNode
	Next  *LinkNode
}

// LRUCache is
type LRUCache struct {
	Capacity int
	Map      map[int]*LinkNode
	Tail     *LinkNode
	Head     *LinkNode
}

// Constructor is
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		Map:      make(map[int]*LinkNode, capacity),
		Tail:     &LinkNode{0, 0, nil, nil},
		Head:     &LinkNode{0, 0, nil, nil},
	}
}

// Get a value by key
func (l *LRUCache) Get(key int) int {
	return 0
}

// Put a key value
func (l *LRUCache) Put(key int, value int) {
	if len(l.table) < l.capacity-1 {
		l.table[key] = value
		return
	}

}
