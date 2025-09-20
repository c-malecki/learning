package queue

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type jobExecFn[T any] func(item *T) error

type JobQueue[T any] struct {
	queue *Queue[T]
	types map[string]jobExecFn[QueueItem[T]]
}

func NewJobQueue[T any](types map[string]jobExecFn[QueueItem[T]], max int) *JobQueue[T] {
	return &JobQueue[T]{
		queue: NewQueue[T](max),
		types: types,
	}
}

func (q *JobQueue[T]) Submit(jobType string, payload T) error {
	if _, ok := q.types[jobType]; !ok {
		return fmt.Errorf("%s is not a valid job type", jobType)
	}

	item := QueueItem[T]{
		Header: Header{
			ID:        uuid.New(),
			Type:      jobType,
			CreatedAt: time.Now(),
		},
		Payload: &payload,
	}

	_, err := q.queue.Enqueue(item)
	return err
}

func (q *JobQueue[T]) ProcessNext() (*T, error) {
	item, err := q.queue.Dequeue()
	if err != nil {
		return nil, err
	}

	exec := q.types[item.Header.Type]
	if err := exec(item); err != nil {
		return nil, fmt.Errorf("exec for job %s failed: %w", item.Header.ID, err)
	}

	return item.Payload, nil
}
