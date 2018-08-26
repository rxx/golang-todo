package main

import (
	"fmt"
)

var tasks TaskList

func init() {
	tasks = TaskList{
		Task{
			title:  "Get up early",
			when:   "tomorrow morning",
			status: "open",
		},
		Task{
			title:  "Make a cup of coffee",
			when:   "tomorrow morning",
			status: "open",
		},
		Task{
			title:  "Meet our new business partner at airport",
			when:   "tomorrow morning",
			status: "open",
		},
	}
}

func main() {
	fmt.Println(tasks)
}
