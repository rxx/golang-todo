package task

import "testing"

func TestTaskString(t *testing.T) {
	task := Task{
		Title:  "Title",
		When:   "tomorrow",
		Status: "open",
	}
	expectedResult := "Title [open] <tomorrow>"

	if task.String() != expectedResult {
		t.Errorf("task.String should return %q but returns %q", expectedResult, task)
	}
}

func TestTaskListRemove(t *testing.T) {
	tasks := List{
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
		Task{
			Title:  "Title3",
			When:   "tomorrow",
			Status: "open",
		},
	}
	tasks.Remove(1)

	expectedResult := `1) Title1 [open] <tomorrow>
2) Title3 [open] <tomorrow>
`

	if tasks.String() != expectedResult {
		t.Errorf("task.String should return %q but returns %q", expectedResult, tasks)
	}
}

func TestTaskListString(t *testing.T) {
	tasks := List{
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
		t.Errorf("task.String should return %q but returns %q", expectedResult, tasks)
	}
}

func TestNewTask(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	expectedTask := Task{Title: "Title", When: "tomorrow", Status: "open"}
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
	task.SetWhen("today")
	if task.When != "today" {
		t.Errorf(`When should set to "today" but was set to %q`, task.When)
	}
}

func TestClose(t *testing.T) {
	task := NewTask("Title", "tomorrow")
	task.Close()
	if task.Status != "close" {
		t.Errorf(`Status should be "close" but status is %q`, task.Status)
	}
}
