package linked_list

import "fmt"

// When and where to use a linked list
// 1) Used in many List, Queue & Stack implementation
// 2) Great for creating circular list
// 3) Can easily model real world objects such as trains
// 4) Used in separate chaining, which is present in certain HASHTABLE implementation
// to deal with hashing collisions
// 5) Often used in the implementation of adjacency list for graphs

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
}

func (l *List) InsertAtHead(data int) {
	newNode := &Node{Data: data, Next: l.Head}
	l.Head = newNode
}

func (l *List) InsertAtTail(data int) {
	newNode := &Node{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *List) DeleteByValue(data int) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == data {
		l.Head = l.Head.Next
	}

	current := l.Head
	for current.Next != nil && current.Next.Data != data {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

func (l *List) Search(data int) bool {
	current := l.Head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

func (l *List) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}
