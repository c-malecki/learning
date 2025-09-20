package queue

import (
	"time"

	"github.com/google/uuid"
)

type MessageQueue[T any] struct {
	queue *Queue[T]
}

func NewMessageQueue[T any](max int) *MessageQueue[T] {
	return &MessageQueue[T]{
		queue: NewQueue[T](max),
	}
}

func (q *MessageQueue[T]) Send(msgType string, payload T) error {
	item := QueueItem[T]{
		Header: Header{
			ID:        uuid.New(),
			Type:      msgType,
			CreatedAt: time.Now(),
		},
		Payload: &payload,
	}

	_, err := q.queue.Enqueue(item)
	return err
}

func (q *MessageQueue[T]) Receive() (*QueueItem[T], error) {
	return q.queue.Dequeue()
}
