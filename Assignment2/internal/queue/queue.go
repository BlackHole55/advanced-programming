package queue

import "github.com/MeirhanSyzdykov/Assignment2/internal/model"

type TaskQueue struct {
	ch chan *model.Task
}

func New(size int) *TaskQueue {
	return &TaskQueue{
		ch: make(chan *model.Task, size),
	}
}

func (q *TaskQueue) Push(task *model.Task) {
	q.ch <- task
}

func (q *TaskQueue) Channel() <-chan *model.Task {
	return q.ch
}

func (q *TaskQueue) Close() {
	close(q.ch)
}
