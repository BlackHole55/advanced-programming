package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/MeirhanSyzdykov/Assignment2/internal/model"
	"github.com/MeirhanSyzdykov/Assignment2/internal/queue"
	"github.com/MeirhanSyzdykov/Assignment2/internal/store"
)

type Handler struct {
	store *store.MemoryStore
	queue *queue.TaskQueue
}

func NewHandler(store *store.MemoryStore, queue *queue.TaskQueue) *Handler {
	return &Handler{
		store: store,
		queue: queue,
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, _ *http.Request) {
	task := &model.Task{
		ID:     uuid.NewString(),
		Status: model.Pending,
	}

	h.store.Save(task)
	h.queue.Push(task)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) Stats(w http.ResponseWriter, _ *http.Request) {
	tasks := h.store.All()

	stats := struct {
		Submitted  int `json:"submitted"`
		Completed  int `json:"completed"`
		InProgress int `json:"in_progress"`
	}{
		Submitted: len(tasks),
	}

	for _, t := range tasks {
		switch t.Status {
		case model.Done:
			stats.Completed++
		case model.Running:
			stats.InProgress++
		}
	}

	json.NewEncoder(w).Encode(stats)
}
