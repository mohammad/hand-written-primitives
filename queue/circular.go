package main

import "fmt"

type CircularQueue struct {
	head int
	tail int
	data []int
	capacity int
	size int
}

func NewCircularQueue(capacity int) *CircularQueue {
	init := make([]int, capacity)
	return &CircularQueue{
		head: 0,
		tail: 0,
		data: init,
		capacity: capacity,
		size: 0,
	}
}

// Enqueue inserts an element at the end (tail) of the queue.
func (q *CircularQueue) Enqueue(val int) {
	if(q.size == len(q.data)) {
		fmt.Printf("Reached max capacity")
		return 
	}

	q.data[q.tail] = val
	q.tail = (q.tail + 1) % q.capacity
	q.size++
}

// Dequeue removes and returns the element at the front (head) of the queue.
func (q *CircularQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
			fmt.Println("Queue empty")
			return 0, false
		}

	val := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--

	return val, true
}

func (q *CircularQueue) IsEmpty() (bool) {
	return q.size == 0
}

func (q *CircularQueue) Size() (int) {
	return q.size
}

func main() {
	q := NewCircularQueue(10)

	q.Enqueue(5)
	q.Enqueue(10)
	q.Enqueue(15)
	dequeue, _ := q.Dequeue()

	fmt.Println(dequeue)
	
}