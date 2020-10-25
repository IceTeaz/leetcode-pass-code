package main

import "fmt"

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

type DoubleList struct {
	head *Node
	tail *Node
	size int
}

func (d *DoubleList) Size() int {
	return d.size
}

func (d *DoubleList) addLast(x *Node) {
	x.prev = d.tail.prev
	x.next = d.tail
	d.tail.prev.next = x
	d.tail.prev = x
	d.size++
}

func (d *DoubleList) remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	d.size--
}

func (d *DoubleList) removeFirst() *Node {
	if d.head.next == d.tail {
		return nil
	}
	first := d.head.next
	d.remove(first)
	return first
}

func NewDoubleList() *DoubleList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoubleList{
		head: head,
		tail: tail,
		size: 0,
	}
}

type LRUCache struct {
	cacheMap map[int]*Node
	cache    *DoubleList
	cap      int
}

func NeLRUCache(capacity int) LRUCache {
	return LRUCache{
		cacheMap: make(map[int]*Node),
		cache:    NewDoubleList(),
		cap:      capacity,
	}
}

func (l LRUCache) makeRecently(key int) {
	n := l.cacheMap[key]
	l.cache.remove(n)
	l.cache.addLast(n)
}

func (l LRUCache) addRecently(key int, val int) {
	n := &Node{
		key:   key,
		value: val,
	}

	l.cacheMap[key] = n
	l.cache.addLast(n)
}

func (l LRUCache) deleteKey(key int) {
	n := l.cacheMap[key]
	l.cache.remove(n)
	delete(l.cacheMap, key)
}

func (l LRUCache) removeLeastRecently() {
	first := l.cache.removeFirst()
	delete(l.cacheMap, first.key)
}

func Constructor(capacity int) LRUCache {
	return NeLRUCache(capacity)
}

func (l LRUCache) Get(key int) int {
	node, ok := l.cacheMap[key]
	if !ok {
		return -1
	}
	l.makeRecently(key)
	return node.value
}

func (l LRUCache) Put(key int, value int) {
	_, ok := l.cacheMap[key]
	if ok {
		l.deleteKey(key)
		l.addRecently(key, value)
		return
	}
	if l.cap == l.cache.size {
		l.removeLeastRecently()
	}
	l.addRecently(key, value)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(1)
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
}
