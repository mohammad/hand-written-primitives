package main

import "fmt"

type Node struct {
    prev *Node
    next *Node
    key int
    val int
}

type DoublyLinkedList struct {
    head *Node
    tail *Node
    size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &DoublyLinkedList{
		head: head,
		tail: tail,
        size: 0,
	}
}

func (dl *DoublyLinkedList) PushFront(n *Node) {
    n.prev = dl.head
    n.next = dl.head.next
    dl.head.next.prev = n
    dl.head.next = n
    dl.size += 1
}

func (dl *DoublyLinkedList) RemoveBack() *Node {
    if(dl.size == 0) {
        return nil
    }

    
    tail := dl.tail.prev
    tail.prev.next = tail.next
    tail.next.prev = tail.prev
    dl.size -= 1

    return tail 
}

func (dl *DoublyLinkedList) Remove(n *Node) {
    n.prev.next = n.next
    n.next.prev = n.prev
    dl.size -= 1
}

type LRUCache struct {
    hashmap map[int]*Node
    dl *DoublyLinkedList
    capacity int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        hashmap: make(map[int]*Node, capacity),
        dl: NewDoublyLinkedList(),
        capacity: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.hashmap[key] 
    
    if !ok {
        return -1
    }
    
    this.dl.Remove(node)
    this.dl.PushFront(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    newNode := &Node{
            key: key,
            val: value,
    }
    node, ok := this.hashmap[key]
    if ok {
        // first remove the node
        this.dl.Remove(node)       
        this.dl.PushFront(newNode)
        this.hashmap[key] = newNode
    } else {
        this.hashmap[key] = newNode
        this.dl.PushFront(newNode)
    }

    if this.dl.size > this.capacity {
        // remove back
        removedNode := this.dl.RemoveBack()
        // remove from hashamp
        delete(this.hashmap, removedNode.key)
    }
}

func main() {

	cache := Constructor(2)

	cache.Put(1, 10)
	cache.Put(2, 20)

	fmt.Println(cache.Get(1)) // 10

	cache.Put(3, 30)

	fmt.Println(cache.Get(2)) // -1 (evicted)
	fmt.Println(cache.Get(3)) // 30
}
