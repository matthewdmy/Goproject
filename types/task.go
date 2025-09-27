package types

import "time"

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusDeleted    TaskStatus = "deleted"
)

type CreateTask struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description" `
	Status      TaskStatus `json:"status" `
	UserID      uint       `json:"user_id"`
	DueDate     time.Time  `json:"due_date"`
	ReminderAt  time.Time  `json:"reminder_at"`
	CreatedAt   time.Time  `json:"created_at" `
	UpdatedAt   time.Time  `json:"updated_at"`
	Tags        []string   `json:"tags" `
}
