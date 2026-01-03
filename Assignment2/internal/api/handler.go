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

func (h *Handler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTask(w, r)

	case http.MethodGet:
		h.GetAllTasks(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Payload string `json:"payload"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	task := &model.Task{
		ID:      uuid.NewString(),
		Payload: req.Payload,
		Status:  model.Pending,
	}

	h.store.Save(task)
	h.queue.Push(task)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"id": task.ID,
	})
}

func (h *Handler) GetAllTasks(w http.ResponseWriter, _ *http.Request) {
	tasks := h.store.All()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
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
