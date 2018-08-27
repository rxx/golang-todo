package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

type Task struct {
	Title  string `json:"title"`
	Status string `json:"status"`
	When   string `json:"when"`
}

// slice of Task
type TaskList []Task

// fmt.Stringer interface for Task
func (t Task) String() string {
	return t.Title + " [" + t.Status + "]" + " <" + t.When + ">"
}

// fmt.Stringer interface for TaskList
func (list TaskList) String() string {
	var result = ""
	for index, t := range list {
		result = result + strconv.Itoa(index+1) + ") " + t.String() + "\n"
	}

	return result
}

func NewTask(title, when string) Task {
	return Task{Title: title, When: when, Status: "open"}
}

func (t *Task) SetTitle(title string) {
	t.Title = title
}

func (t *Task) SetWhen(when string) {
	t.When = when
}

func (t *Task) Close() {
	t.Status = "close"
}

func SaveJSON(tasks TaskList, fileName string) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create(fileName)
	defer jsonFile.Close()
	if err != nil {
		panic(err)
	}

	jsonFile.Write(data)
}

func LoadJSON(fileName string) TaskList {
	var tasks TaskList

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}

	return tasks
}
