package main

import "strconv"

type Task struct {
	title  string
	status string
	when   string
}

// slice of Task
type TaskList []Task

// fmt.Stringer interface for Task
func (t Task) String() string {
	return t.title + " [" + t.status + "]" + " <" + t.when + ">"
}

// fmt.Stringer interface for TaskList
func (list TaskList) String() string {
	var result string = "There are " + strconv.Itoa(len(list)) + " tasks:\n\n"
	for index, t := range list {
		result = result + strconv.Itoa(index+1) + ") " + t.String() + "\n"
	}

	return result
}
