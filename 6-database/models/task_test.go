package models

import "testing"

func TestTaskString(t *testing.T) {
	task := Task{
		Title:  "Title",
		DueDate:   "tomorrow",
		Status: "open",
	}
	expectedResult := "Title [open] <tomorrow>"

	if task.String() != expectedResult {
		t.Errorf("task.String should return %q but returns %q", expectedResult, task)
	}
}

func TestTaskTasksString(t *testing.T) {
	tasks := Tasks{
		Task{
			Title:  "Title1",
			DueDate:   "tomorrow",
			Status: "open",
		},
		Task{
			Title:  "Title2",
			DueDate:   "tomorrow",
			Status: "open",
		},
	}
	expectedResult := `1) Title1 [open] <tomorrow>
2) Title2 [open] <tomorrow>
`

	if tasks.String() != expectedResult {
		t.Errorf("task.String should return %q but returns %q", expectedResult, tasks)
	}
}

func TestNewTask(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	expectedTask := Task{Title: "Title", DueDate: "tomorrow", Status: "open"}
	if task != expectedTask {
		t.Errorf("NewTask should create a task %q but creates %q ", expectedTask, task)
	}
}

func TestSetTitle(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.SetTitle("New Title")
	if task.Title != "New Title" {
		t.Errorf(`Title should set to "New Title" but was set to %q`, task.Title)
	}
}

func TestSetWhen(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.SetDueDate("today")
	if task.DueDate != "today" {
		t.Errorf(`DueDate should set to "today" but was set to %q`, task.DueDate)
	}
}

func TestClose(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.Close()
	if task.Status != "close" {
		t.Errorf(`Status should be "close" but status is %q`, task.Status)
	}
}
