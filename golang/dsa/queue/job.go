package queue

import (
	"fmt"
	"sync"

	"github.com/c-malecki/learning/dsa/list"
)

type execFn func(j Job) error

type JobQueue struct {
	list     *list.LinkedList[Job]
	jobTypes map[string]execFn
	max      int
	lock     sync.Mutex
}

type Job struct {
	Header  Header
	Payload interface{}
}

func NewJobQueue(jobTypes map[string]execFn, max int) *JobQueue {
	m := 0
	if max != 0 {
		m = max
	}
	return &JobQueue{
		list:     list.New[Job](),
		jobTypes: jobTypes,
		max:      m,
	}
}

func (q *JobQueue) Enqueue(job Job) (*list.Node[Job], error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if _, ok := q.jobTypes[job.Header.Type]; !ok {
		return nil, fmt.Errorf("%s is not a valid job type", job.Header.Type)
	}

	if q.max != 0 && q.list.Size() == q.max {
		// handle overflow
		return nil, fmt.Errorf("query is full")
	}

	node := q.list.AppendValue(job)

	return node, nil
}

func (q *JobQueue) Dequeue() (*Job, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.list.Size() == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	node := q.list.Front()
	exec := q.jobTypes[node.Value.Header.Type]

	if err := exec(node.Value); err != nil {
		return nil, err
	}

	job := q.list.Remove(node)

	return &job, nil
}
