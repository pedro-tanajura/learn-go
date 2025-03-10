package main

import "fmt"


type Node struct {
	value int
	next  *Node
}


type LinkedList struct {
	head *Node
	size int
}


func Constructor() LinkedList {
	return LinkedList{
		head: nil,
		size: 0,
	}
}


func (list *LinkedList) InsertAtHead(value int) {
	newNode := &Node{value: value, next: list.head}
	list.head = newNode
	list.size++
}


func (list *LinkedList) InsertAtTail(value int) {
	newNode := &Node{value: value, next: nil}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	list.size++
}


func (list *LinkedList) Delete(value int) {
	if list.head != nil {
		current := list.head
		if current.value == value {
			list.head = current.next
		} else {
			for current.next != nil {
				if current.next.value == value {
					current.next = current.next.next
					break
				}
				current = current.next
			}
		}
		list.size--
	}
}


func (list *LinkedList) Search(value int) bool {
	current := list.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Print("\n")
}

func (list *LinkedList) InvertList() {
	var prev *Node = nil
	current := list.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	list.head = prev
}

func main() {
	
	list := Constructor()

	list.InsertAtHead(1)
	list.InsertAtHead(2)
	list.InsertAtHead(3)
	list.InsertAtHead(4)
	list.InsertAtHead(10)
	list.PrintList()
	list.InsertAtTail(20)
	list.PrintList()
	fmt.Println("Search 10:", list.Search(10))
	list.Delete(10)
	list.PrintList()
	fmt.Println("Search 10:", list.Search(10))
	fmt.Println("List size: ", list.size)
	list.PrintList()
	list.InvertList()
	list.PrintList()
}
