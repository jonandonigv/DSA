package listq

import (
	"container/list"
	"fmt"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

func (q *Queue) Enqueue(item interface{}) {
	q.list.PushBack(item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.list.Len() == 0 {
		return nil, fmt.Errorf("queue is empy")
	}

	e := q.list.Front()
	q.list.Remove(e)
	return e.Value, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.list.Len() == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	return q.list.Front().Value, nil
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}

func (q *Queue) Size() int {
	return q.list.Len()
}
