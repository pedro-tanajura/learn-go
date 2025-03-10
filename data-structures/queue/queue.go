package main

import "fmt"


type Queue struct {
	elements []int
}


func Constructor() Queue {
	return Queue{
		elements: []int{},
	}
}


func (q *Queue) Enqueue(val int) {
	q.elements = append(q.elements, val)
}


func (q *Queue) Dequeue() (int, bool) {
	if len(q.elements) == 0 {
		fmt.Println("Queue is empty")
		return -1, false
	}
	aux := q.elements[0]
	q.elements = q.elements[1:]
	return aux, true
}


func (q *Queue) Front() (int, bool) {
	if len(q.elements) == 0 {
		fmt.Println("Queue is empty")
		return -1, false
	}
	return q.elements[0], true
}


func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func main() {
	
	Queue := Constructor()

	Queue.Enqueue(1)
	Queue.Enqueue(2)
	Queue.Enqueue(3)
	val, ok := Queue.Front()
	fmt.Println("Front element:", val, ok)

	val, ok = Queue.Dequeue()
	fmt.Println("Front element after Dequeue:", val, ok)

	Queue.Dequeue()
	Queue.Dequeue()
	fmt.Println("Is Queue empty?", Queue.IsEmpty())
}
