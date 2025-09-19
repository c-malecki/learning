package queue

import (
	"fmt"
	"sync"
	"time"

	"github.com/c-malecki/learning/dsa/list"
)

type MessageQueue struct {
	list         *list.LinkedList[Message]
	messageTypes map[string]struct{}
	max          int
	lock         sync.Mutex
}

type Header struct {
	ID        int
	Type      string
	CreatedAt time.Time
}

type Message struct {
	Header  Header
	Payload interface{}
}

func NewMessageQueue(messageTypes map[string]struct{}, max int) *MessageQueue {
	m := 0
	if max != 0 {
		m = max
	}
	return &MessageQueue{
		list:         list.New[Message](),
		messageTypes: messageTypes,
		max:          m,
	}
}

func (q *MessageQueue) Enqueue(msg Message) (*list.Node[Message], error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if _, ok := q.messageTypes[msg.Header.Type]; !ok {
		return nil, fmt.Errorf("%s is not a valid message type", msg.Header.Type)
	}

	if q.max != 0 && q.list.Size() == q.max {
		// handle overflow
		return nil, fmt.Errorf("query is full")
	}

	node := q.list.AppendValue(msg)

	return node, nil
}

func (q *MessageQueue) Dequeue() (*Message, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.list.Size() == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	node := q.list.Front()
	msg := q.list.Remove(node)

	return &msg, nil
}
