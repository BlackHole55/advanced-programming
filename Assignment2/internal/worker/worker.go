package worker

import (
	"time"

	"github.com/MeirhanSyzdykov/Assignment2/internal/model"
	"github.com/MeirhanSyzdykov/Assignment2/internal/store"
)

type WorkerPool struct {
	store *store.MemoryStore
}

func NewWorkerPool(count int, q <-chan *model.Task, store *store.MemoryStore) {
	wp := &WorkerPool{store: store}

	for range count {
		go wp.worker(q)
	}
}

func (wp *WorkerPool) worker(q <-chan *model.Task) {
	for task := range q {
		task.Status = model.Running

		// Simulate work
		time.Sleep(5 * time.Second)

		task.Status = model.Done
		task.Result = "Task completed"
		wp.store.Save(task)
	}
}
