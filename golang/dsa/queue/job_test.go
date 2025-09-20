package queue

import (
	"fmt"
	"testing"
)

func TestJobQueue(t *testing.T) {
	q := NewJobQueue(
		map[string]jobExecFn[QueueItem[tStruct]]{
			"valid": func(item *QueueItem[tStruct]) error {
				return fmt.Errorf("failed")
			},
		}, 1)

	err := q.Submit("invalid", tStruct{Value: 1})
	if err == nil {
		t.Error("expected error for invalid job type")
	}

	err = q.Submit("valid", tStruct{Value: 1})
	if err != nil {
		t.Fatal("failed to submit job")
	}

	err = q.Submit("valid", tStruct{Value: 1})
	if err == nil {
		t.Error("expected error when job queue is full")
	}

	_, err = q.ProcessNext()
	if err == nil {
		t.Error("expected error for exec func")
	}

	_, err = q.ProcessNext()
	if err == nil {
		t.Error("expected error when processing next job from empty queue")
	}
}
