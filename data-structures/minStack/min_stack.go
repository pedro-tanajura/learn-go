package main

import "fmt"


type MinStack struct {
	elements []int
	minStack []int
}


func Constructor() MinStack {
	return MinStack{
		elements: []int{},
		minStack: []int{},
	}
}


func (ms *MinStack) Push(val int) {
	ms.elements = append(ms.elements, val)
	if len(ms.minStack) > 0 {
		if val <= ms.minStack[len(ms.minStack)-1] {
			ms.minStack = append(ms.minStack, val)
		}
	} else {
		ms.minStack = append(ms.minStack, val)
	}
}


func (ms *MinStack) Pop() {
	if len(ms.elements) > 0 {
		if ms.elements[len(ms.elements)-1] == ms.minStack[len(ms.minStack)-1] {
			ms.minStack = ms.minStack[:len(ms.minStack)-1]
		}
		ms.elements = ms.elements[:len(ms.elements)-1]
	}
}


func (ms *MinStack) Top() int {
	if len(ms.elements) > 0 {
		return ms.elements[len(ms.elements)-1]
	}
	fmt.Println("Empty Stack")
	return -1
}


func (ms *MinStack) GetMin() int {
	if len(ms.minStack) > 0 {
		return ms.minStack[len(ms.minStack)-1]
	}
	fmt.Println("Empty Stack")
	return -1
}

func main() {
	
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("Minimum element:", minStack.GetMin())

	minStack.Pop()
	fmt.Println("Top element:", minStack.Top())
	fmt.Println("Minimum element:", minStack.GetMin())
}
