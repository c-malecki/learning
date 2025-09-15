package queue

import (
	"sync"
	"time"

	"github.com/c-malecki/learning/dsa/list"
)

type MessageQueue struct {
	list *list.LinkedList[Message]
	max  int
	lock sync.Mutex
}

type Header struct {
	MessageID   int
	MessageType string
	CreatedAt   time.Time
}

type Message struct {
	Header  Header
	Payload interface{}
}

func NewMessageQueue(max int) *MessageQueue {
	m := 0
	if max != 0 {
		m = max
	}
	return &MessageQueue{
		list: list.New[Message](),
		max:  m,
	}
}

func (q *MessageQueue) Enqueue(msg Message) *list.Node[Message] {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.max != 0 && q.list.Size() == q.max {
		// handle overflow
		return nil
	}
	node := q.list.AppendValue(msg)
	return node
}

func (q *MessageQueue) Dequeue() *list.Node[Message] {
	q.lock.Lock()
	defer q.lock.Unlock()
	front := q.list.Front()
	switch front.Value.Header.MessageType {
	case "one":
	case "two":
	default:
	}
	// do stuff
	q.list.Remove(front)
	return front
}
