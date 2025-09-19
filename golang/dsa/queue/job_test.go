package queue

import (
	"fmt"
	"testing"
)

func checkInvalidJobType(t *testing.T) {
	q := NewJobQueue(
		map[string]execFn{
			"valid": func(m Job) error {
				return nil
			},
		}, 0)

	job := Job{
		Header: Header{Type: "invalid"},
	}

	_, err := q.Enqueue(job)
	if err == nil {
		t.Error("expected error for invalid job type")
	}
}

func checkJobQueueMax(t *testing.T) {
	q := NewJobQueue(
		map[string]execFn{
			"valid": func(m Job) error {
				return nil
			},
		}, 1)

	job := Job{
		Header: Header{
			Type: "valid",
		},
	}

	_, err := q.Enqueue(job)
	if err != nil {
		t.Fatal("failed to enqueue job")
	}

	_, err = q.Enqueue(job)
	if err == nil {
		t.Error("expected error when queue is full")
	}
}

func checkJobExec(t *testing.T) {
	q := NewJobQueue(
		map[string]execFn{
			"valid": func(m Job) error {
				return fmt.Errorf("failed")
			},
		}, 0)

	job := Job{
		Header: Header{
			Type: "valid",
		},
	}

	_, err := q.Enqueue(job)
	if err != nil {
		t.Fatal("failed to enqueue job")
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Error("expected error for exec func")
	}
}

func checkJobEmptyDequeue(t *testing.T) {
	q := NewJobQueue(map[string]execFn{}, 0)

	_, err := q.Dequeue()
	if err == nil {
		t.Error("expected error when dequeuing from empty queue")
	}
}

func TestJobQueue(t *testing.T) {
	checkInvalidJobType(t)
	checkJobQueueMax(t)
	checkJobExec(t)
	checkJobEmptyDequeue(t)
}
