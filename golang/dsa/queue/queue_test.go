package queue

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

type tStruct struct {
	Value int
}

func TestQueue(t *testing.T) {
	q := NewQueue[tStruct](1)

	item := QueueItem[tStruct]{
		Header: Header{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
	}

	_, err := q.Enqueue(item)
	if err != nil {
		t.Fatal("failed to enqueue job")
	}

	_, err = q.Enqueue(item)
	if err == nil {
		t.Error("expected error when queue is full")
	}

	_, err = q.Dequeue()
	if err != nil {
		t.Fatal("failed to dequeue item")
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Error("expected error when dequeuing from empty queue")
	}
}
