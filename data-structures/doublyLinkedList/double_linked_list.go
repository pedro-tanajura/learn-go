package main


type Node struct {
	value int
	prev  *Node
	next  *Node
}


type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}


func Constructor() DoublyLinkedList {
	return DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}


func (list *DoublyLinkedList) InsertAtHead(value int) {
	newNode := &Node{value: value, next: list.head, prev: nil}
	if list.head != nil {
		list.head.prev = newNode
	} else {
		list.tail = newNode
	}
	list.head = newNode
	list.size++
}


func (list *DoublyLinkedList) InsertAtTail(value int) {
	newNode := &Node{value: value, next: nil, prev: list.tail}
	if list.tail != nil {
		list.tail.next = newNode
	} else {
		list.head = newNode
	}
	list.tail = newNode
	list.size++
}


func (list *DoublyLinkedList) Delete(value int) {
	current := list.head
	if current == nil {
		return
	}
	if current.value == value {
		list.head = current.next
		if list.head != nil {
			list.head.prev = nil
		} else {
			list.tail = nil
		}
		list.size--
		return
	}
	for current != nil {
		if current.value == value {
			if current.next != nil {
				current.next.prev = current.prev
			} else {
				list.tail = current.prev
			}
			if current.prev != nil {
				current.prev.next = current.next
			}
			list.size--
			return
		}
		current = current.next
	}
}


func (list *DoublyLinkedList) Search(value int) bool {
	current := list.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}


func (list *DoublyLinkedList) PrintForward() []int {
	result := []int{}
	current := list.head
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}
	return result
}


func (list *DoublyLinkedList) PrintBackward() []int {
	result := []int{}
	current := list.tail
	for current != nil {
		result = append(result, current.value)
		current = current.prev
	}
	return result
}

func main() {
	
	list := Constructor()

	list.InsertAtTail(1)
	list.InsertAtHead(10)
	list.InsertAtTail(20)
	list.InsertAtTail(30)
	list.PrintForward()
	list.Delete(20)
	list.PrintForward()
	list.PrintBackward()
}
