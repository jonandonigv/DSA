package main

import (
	"fmt"

	"github.com/jonandonigv/DSA/arrays"
	linked_list "github.com/jonandonigv/DSA/list"
	"github.com/jonandonigv/DSA/list/d_linked_list"
	"github.com/jonandonigv/DSA/stack"
)

func main() {
	fmt.Println("Arrays: ")
	da := arrays.DynamicArray{}
	da.NewDynamicArray(2)
	da.Append(1)
	da.Append(2)
	da.Append(3)
	da.Append(4)
	fmt.Println("Dynamic Array: ", da)
	da.Print()
	fmt.Printf("Length: %d, Capacity: %d\n", da.Len(), da.Cap())

	l := linked_list.List{}
	l.InsertAtTail(1)
	l.InsertAtTail(1)
	l.InsertAtHead(5)
	l.InsertAtTail(1)

	fmt.Println("Linked list:")
	l.Print()

	fmt.Println("Search 5:", l.Search(5))
	l.DeleteByValue(1)
	fmt.Println("After deleting 1:")
	l.Print()

	dl := d_linked_list.DoublyLinkedList{}

	fmt.Println("Doubly linked list: ")

	dl.Append(1)
	dl.Append(2)
	dl.Append(3)
	dl.Append(0)
	dl.Append(4)

	fmt.Println("Forward traversal: ")
	dl.Display()

	fmt.Println("Reverse traversal: ")
	dl.DisplayReverse()

	dl.Delete(2)
	fmt.Println("After deleting 2:")
	dl.Display()

	fmt.Println("Stack ->")
	s := stack.Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println("Stack size: ", s.Size())

	if top, err := s.Peek(); err == nil {
		fmt.Println("Top element: ", top)
	}

	for !s.IsEmpty() {
		if item, err := s.Pop(); err == nil {
			fmt.Println("Popped: ", item)
		}
	}

	if _, err := s.Pop(); err != nil {
		fmt.Println("Error: ", err)
	}
}
