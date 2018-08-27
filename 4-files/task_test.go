package main

import "testing"

func TestTaskString(t *testing.T) {
	task := Task{
		Title:  "Title",
		When:   "tomorrow",
		Status: "open",
	}
	expectedResult := "Title [open] <tomorrow>"

	if task.String() != expectedResult {
		t.Error("task.String should return " + expectedResult + "but returns " + task.String())
	}
}

func TestTaskListString(t *testing.T) {
	tasks := TaskList{
		Task{
			Title:  "Title1",
			When:   "tomorrow",
			Status: "open",
		},
		Task{
			Title:  "Title2",
			When:   "tomorrow",
			Status: "open",
		},
	}
	expectedResult := `1) Title1 [open] <tomorrow>
2) Title2 [open] <tomorrow>
`

	if tasks.String() != expectedResult {
		t.Error("task.String should return " + expectedResult + "but returns " + tasks.String())
	}
}

func TestNewTask(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	expectedTask := Task{Title: "Title", When: "tomorrow", Status: "open"}
	if task != expectedTask {
		t.Error("NewTask should create a task " + expectedTask.String() + " but creates " + task.String())
	}
}

func TestSetTitle(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.SetTitle("New Title")
	if task.Title != "New Title" {
		t.Error("Title should set to \"New Title\" but was set to \"" + task.Title + "\"")
	}
}

func TestSetWhen(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.SetWhen("today")
	if task.When != "today" {
		t.Error("When should set to \"today\" but was set to \"" + task.When + "\"")
	}
}

func TestClose(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.Close()
	if task.Status != "close" {
		t.Error("Status should be \"close\" but Status is \"" + task.Status + "\"")
	}
}
