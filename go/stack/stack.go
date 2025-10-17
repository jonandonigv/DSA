package stack

import "errors"

// When and where to use a stack
// 1) Used by undo mechanisms in text editors
// 2) Used by compiler syntax checking for matching brackets and braces
// 3) Can be used to model a pile of books
// 4) Used behind the scenes to support recursion by keeping track of previous function calls
// 5) Can be used to do Depth First Search (DFS) on a graph

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index] // Remove the top element
	return item, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}
