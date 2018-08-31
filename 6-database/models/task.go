package models

import (
	"fmt"
	"time"

	"github.com/gobuffalo/uuid"
)

// Task is main struct
type Task struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title     string    `json:"title" db:"title"`
	DueDate   string    `json:"due_date" db:"due_date"`
	Status    string    `json:"status" db:"status"`
}

type Tasks []Task

// fmt.Stringer interface for Task
func (t Task) String() string {
	return fmt.Sprintf("%s [%s] <%s>", t.Title, t.Status, t.DueDate)
}

// fmt.Stringer interface for Tasks
func (tasks Tasks) String() string {
	var result = ""
	for index, t := range tasks {
		result += fmt.Sprintf("%d) %s\n", index+1, t)
	}

	return result
}

// NewTask create a new task
func NewTask(title, dueDate string) Task {
	return Task{Title: title, DueDate: dueDate, Status: "open"}
}

// SetTitle sets title to task
func (t *Task) SetTitle(title string) {
	t.Title = title
}

// SetDueDate sets dueDate to task
func (t *Task) SetDueDate(dueDate string) {
	t.DueDate = dueDate
}

// Close closes a task
func (t *Task) Close() {
	t.Status = "close"
}
