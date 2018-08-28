package task

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

// Task is main struct
type Task struct {
	Title  string `json:"title"`
	Status string `json:"status"`
	When   string `json:"when"`
}

// List is a slice of Task
type List []Task

// fmt.Stringer interface for Task
func (t Task) String() string {
	return t.Title + " [" + t.Status + "]" + " <" + t.When + ">"
}

// fmt.Stringer interface for List
func (list List) String() string {
	var result = ""
	for index, t := range list {
		result = result + strconv.Itoa(index+1) + ") " + t.String() + "\n"
	}

	return result
}

// Remove removes task from list
func (list *List) Remove(pos int) {
	slice := []Task(*list)
	*list = append(slice[:pos], slice[pos+1:]...)
}

// NewTask create a new task
func NewTask(title, when string) Task {
	return Task{Title: title, When: when, Status: "open"}
}

// SetTitle sets title to task
func (t *Task) SetTitle(title string) {
	t.Title = title
}

// SetWhen sets when to task
func (t *Task) SetWhen(when string) {
	t.When = when
}

// Close closes a task
func (t *Task) Close() {
	t.Status = "close"
}

//SaveJSON saves tasks to file
func SaveJSON(tasks List, fileName string) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(fileName)
	defer jsonFile.Close()
	if err != nil {
		return err
	}

	jsonFile.Write(data)
	return nil
}

// LoadJSON loads tasks from file
func LoadJSON(fileName string) (List, error) {
	var tasks List

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return tasks, nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
