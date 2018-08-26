package main

import "testing"

func TestTaskString(t *testing.T) {
	task := Task{
		title:  "Title",
		when:   "tomorrow",
		status: "open",
	}
	expectedResult := "Title [open] <tomorrow>"

	if task.String() != expectedResult {
		t.Error("task.String should return " + expectedResult + "but returns " + task.String())
	}
}

func TestTaskListString(t *testing.T) {
	tasks := TaskList{
		Task{
			title:  "Title1",
			when:   "tomorrow",
			status: "open",
		},
		Task{
			title:  "Title2",
			when:   "tomorrow",
			status: "open",
		},
	}
	expectedResult := `There are 2 tasks:

1) Title1 [open] <tomorrow>
2) Title2 [open] <tomorrow>
`

	if tasks.String() != expectedResult {
		t.Error("task.String should return " + expectedResult + "but returns " + tasks.String())
	}
}

func TestNewTask(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	expectedTask := Task{title: "Title", when: "tomorrow", status: "open"}
	if task != expectedTask {
		t.Error("NewTask should create a task " + expectedTask.String() + " but creates " + task.String())
	}
}

func TestSetTitle(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.SetTitle("New title")
	if task.title != "New title" {
		t.Error("Title should set to \"New title\" but was set to \"" + task.title + "\"")
	}
}

func TestSetWhen(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.SetWhen("today")
	if task.when != "today" {
		t.Error("When should set to \"today\" but was set to \"" + task.when + "\"")
	}
}

func TestClose(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.Close()
	if task.status != "close" {
		t.Error("Status should be \"close\" but status is \"" + task.status + "\"")
	}
}
