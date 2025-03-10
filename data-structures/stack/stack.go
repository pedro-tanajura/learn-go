package main

import "fmt"


type Stack struct {
	elements []int
}


func Constructor() Stack {
	return Stack{
		elements: []int{},
	}
}


func (s *Stack) Push(val int) {
	s.elements = append(s.elements, val)
}


func (s *Stack) Pop() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	s.elements = s.elements[:len(s.elements)-1]
}


func (s *Stack) Top() int {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.elements[len(s.elements)-1]
}


func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	
	stack := Constructor()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Top element:", stack.Top())

	stack.Pop()
	fmt.Println("Top element after Pop:", stack.Top())

	stack.Pop()
	stack.Pop()
	fmt.Println("Is stack empty?", stack.IsEmpty())
}
