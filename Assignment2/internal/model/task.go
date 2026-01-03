package model

type TaskStatus string

const (
	Pending TaskStatus = "PENDING"
	Running TaskStatus = "RUNNING"
	Done    TaskStatus = "DONE"
)

type Task struct {
	ID      string     `json:"id"`
	Status  TaskStatus `json:"status"`
	Payload string     `json:"payload"`
	Result  string     `json:"result,omitempty"`
}
