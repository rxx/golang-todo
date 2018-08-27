package main

import (
	"fmt"
)

var tasks TaskList

const fileName = "tasks.json"

func main() {
	tasks = append(tasks, NewTask("Get up early", "tomorrow morning"))
	tasks = append(tasks, NewTask("Make a cup of coffee", "tomorrow morning"))
	tasks = append(tasks, NewTask("Meet our new business partner at airport", "tomorrow morning"))

	tasks[0].Close()
	tasks[1].SetWhen("today")
	tasks[2].SetTitle("Go to work")

	fmt.Println("There are next tasks")
	fmt.Println(tasks)

	SaveJSON(tasks, fileName)
	loadedTasks := LoadJSON(fileName)

	fmt.Println("Loaded tasks")
	fmt.Println(loadedTasks)
}
