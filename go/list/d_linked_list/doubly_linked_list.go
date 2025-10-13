package d_linked_list

import "fmt"

type Node struct {
	Next  *Node
	Value int
	Prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (dl *DoublyLinkedList) Append(val int) {
	newNode := &Node{Value: val, Prev: nil, Next: nil}
	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
		return
	}
	newNode.Prev = dl.Tail
	dl.Tail.Next = newNode
	dl.Tail = newNode
}

func (dl *DoublyLinkedList) Prepend(val int) {
	newNode := &Node{Value: val, Prev: nil, Next: nil}

	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
		return
	}

	newNode.Next = dl.Head
	dl.Head.Prev = newNode
	dl.Head = newNode
}

func (dl *DoublyLinkedList) Delete(val int) {
	current := dl.Head

	for current != nil {
		if current.Value == val {
			if current == dl.Head && current == dl.Tail {
				dl.Head = nil
				dl.Tail = nil
				return
			}

			if current == dl.Head {
				dl.Head = current.Next
				dl.Head.Prev = nil
				return
			}

			if current == dl.Tail {
				dl.Tail = current.Prev
				dl.Tail.Prev = nil
				return
			}
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
			return
		}
		current = current.Next
	}
}

func (dl *DoublyLinkedList) Display() {
	current := dl.Head

	for current != nil {
		fmt.Printf("%d <-> ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

func (dl *DoublyLinkedList) DisplayReverse() {
	current := dl.Tail
	for current != nil {
		fmt.Printf("%d <-> ", current.Value)
		current = current.Prev
	}
	fmt.Println()
}
