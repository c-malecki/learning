package queue

import (
	"fmt"
	"sync"
	"time"

	"github.com/c-malecki/learning/dsa/list"
	"github.com/google/uuid"
)

type Header struct {
	ID        uuid.UUID
	Type      string
	CreatedAt time.Time
}

type QueueItem[T any] struct {
	Header  Header
	Payload *T
}

type Queue[T any] struct {
	list *list.SinglyLinkedList[QueueItem[T]]
	max  int
	lock sync.Mutex
}

func NewQueue[T any](max int) *Queue[T] {
	return &Queue[T]{
		list: list.NewSinglyLinkedList[QueueItem[T]](),
		max:  max,
	}
}

func (q *Queue[T]) Enqueue(item QueueItem[T]) (*QueueItem[T], error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.max != 0 && q.list.Size() >= q.max {
		return nil, fmt.Errorf("queue is full")
	}

	node := q.list.AppendValue(item)
	return &node.Value, nil
}

func (q *Queue[T]) Dequeue() (*QueueItem[T], error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.list.Size() == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	item := q.list.Front()
	q.list.Remove(item)

	return &item.Value, nil
}

func (q *Queue[T]) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.list.Size()
}
