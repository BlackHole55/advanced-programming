package store

import (
	"sync"

	"github.com/MeirhanSyzdykov/Assignment2/internal/model"
)

type MemoryStore struct {
	mu    sync.Mutex
	tasks map[string]*model.Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: make(map[string]*model.Task),
	}
}

func (m *MemoryStore) Save(task *model.Task) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tasks[task.ID] = task
}

func (m *MemoryStore) Get(id string) (*model.Task, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	task, ok := m.tasks[id]

	return task, ok
}

func (m *MemoryStore) All() []*model.Task {
	m.mu.Lock()
	defer m.mu.Unlock()

	result := make([]*model.Task, 0, len(m.tasks))

	for _, task := range m.tasks {
		result = append(result, task)
	}

	return result
}
