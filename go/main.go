package main

import (
	"fmt"

	linked_list "github.com/jonandonigv/DSA/list"
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
}
