package queue

import "fmt"

// When and where to used a Queue
// 1) Any waiting line models a queue, for example a lineup at a movie theater
// 2) Can be used to efficiently keep track of the x most recently added elements
// 3) Web server request management where you want first come first serve
// 4) Bread first search (BFS) graph traversal

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empy")
	}
	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func NewQueue() *Queue {
	return &Queue{}
}
