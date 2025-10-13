package main

import (
	"fmt"

	linked_list "github.com/jonandonigv/DSA/list"
	"github.com/jonandonigv/DSA/list/d_linked_list"
)

func main() {
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

	fmt.Println("Forward traversal: ")
	dl.Display()

	fmt.Println("Reverse traversal: ")
	dl.DisplayReverse()

	dl.Delete(2)
	fmt.Println("After deleting 2:")
	dl.Display()
}
