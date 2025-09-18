package tasks

import "time"

type TaskStatus int
const  (
	StatusPending  TaskStatus = iota
	StatusInProgress
	StatusCompleted 
)

type Task struct {
	ID          int
	Status      TaskStatus
	Description string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTask(Description string) *Task {
	return &Task{
		Description: Description,
		Status:      StatusPending,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
}

func (t *Task) UpdateStatus(newStatus TaskStatus) {
	t.Status = newStatus
	t.updatedAt = time.Now()
}

func (t *Task) Update(description string) {
	t.Description = description
	t.updatedAt = time.Now()
}

