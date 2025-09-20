package queue

import "testing"

func TestMessageQueue(t *testing.T) {
	q := NewMessageQueue[tStruct](1)

	err := q.Send("test", tStruct{Value: 1})
	if err != nil {
		t.Fatal("failed to send message")
	}

	err = q.Send("test", tStruct{Value: 1})
	if err == nil {
		t.Error("expected error when message queue is full")
	}

	_, err = q.Receive()
	if err != nil {
		t.Fatal("failed to receive message")
	}

	_, err = q.Receive()
	if err == nil {
		t.Error("expected error when receiving message from empty queue")
	}
}
