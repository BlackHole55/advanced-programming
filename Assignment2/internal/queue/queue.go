package queue

type TaskQueue[T any] struct {
	ch chan T
}

func NewTaskQueue[T any](size int) *TaskQueue[T] {
	return &TaskQueue[T]{
		ch: make(chan T, size),
	}
}

func (q *TaskQueue[T]) Push(task T) {
	q.ch <- task
}

func (q *TaskQueue[T]) Channel() <-chan T {
	return q.ch
}

func (q *TaskQueue[T]) Close() {
	close(q.ch)
}
