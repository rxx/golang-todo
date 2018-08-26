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
