package main

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	list := Constructor()

	
	list.InsertAtTail(1)
	list.InsertAtTail(2)
	list.InsertAtTail(3)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(list.PrintForward(), expected) {
		t.Errorf("InsertAtTail failed, expected %v, got %v", expected, list.PrintForward())
	}

	
	list.InsertAtHead(0)
	expected = []int{0, 1, 2, 3}
	if !reflect.DeepEqual(list.PrintForward(), expected) {
		t.Errorf("InsertAtHead failed, expected %v, got %v", expected, list.PrintForward())
	}

	
	list.Delete(0)
	expected = []int{1, 2, 3}
	if !reflect.DeepEqual(list.PrintForward(), expected) {
		t.Errorf("Delete failed for head node, expected %v, got %v", expected, list.PrintForward())
	}

	
	list.Delete(2)
	expected = []int{1, 3}
	if !reflect.DeepEqual(list.PrintForward(), expected) {
		t.Errorf("Delete failed for middle node, expected %v, got %v", expected, list.PrintForward())
	}

	
	list.Delete(3)
	expected = []int{1}
	if !reflect.DeepEqual(list.PrintForward(), expected) {
		t.Errorf("Delete failed for tail node, expected %v, got %v", expected, list.PrintForward())
	}

	
	list.Delete(1)
	expected = []int{}
	if !reflect.DeepEqual(list.PrintForward(), expected) {
		t.Errorf("Delete failed for last node, expected %v, got %v", expected, list.PrintForward())
	}

	
	list.InsertAtHead(10)
	list.InsertAtHead(20)
	expected = []int{10, 20}
	if !reflect.DeepEqual(list.PrintBackward(), expected) {
		t.Errorf("PrintBackward failed, expected %v, got %v", expected, list.PrintBackward())
	}
}
