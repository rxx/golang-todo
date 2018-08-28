package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rxx/golang-todo/5-cobra/task"
	"github.com/spf13/cobra"
)

var (
	when        string
	currentTask int
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		needToSave = false
		if len(tasks) < 1 {
			fmt.Println("No tasks")
		} else {
			fmt.Println(tasks)
		}
	},
}

var addCmd = &cobra.Command{
	Use:   "add title [when]",
	Short: `Add new open task`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var title string
		title = strings.Join(args, " ")
		task := task.NewTask(title, when)
		tasks = append(tasks, task)

		if verbose {
			fmt.Println("Added a task", task)
		}
	},
}

var editTitleCmd = &cobra.Command{
	Use:   "edit_title task_position title",
	Short: "Edit task Title",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		position, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		title := strings.Join(args[1:], " ")
		tasks[position-1].SetTitle(title)

		if len(when) > 1 {
			tasks[position-1].SetWhen(when)
		}

		if verbose {
			fmt.Println("Changed task", tasks[position-1])
		}
	},
}

var editWhenCmd = &cobra.Command{
	Use:   "edit_when task_position when",
	Short: "Edit task When",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		position, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		when = strings.Join(args[1:], " ")
		tasks[position-1].SetWhen(when)

		if verbose {
			fmt.Println("Changed task", tasks[position-1])
		}
	},
}

var closeCmd = &cobra.Command{
	Use:   "close task_position",
	Short: "Close a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		position, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		tasks[position-1].Close()

		if verbose {
			fmt.Printf("Closed task at %d position\n", position)
		}
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove task_position",
	Short: "Remove task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		position, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		tasks.Remove(position - 1)

		if verbose {
			fmt.Printf("Removed task at %d position\n", position)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(editTitleCmd)
	rootCmd.AddCommand(editWhenCmd)
	rootCmd.AddCommand(closeCmd)
	rootCmd.AddCommand(removeCmd)

	addCmd.Flags().StringVar(&when, "when", "today", "Task when time")
}
