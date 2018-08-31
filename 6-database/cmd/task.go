package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gobuffalo/pop"
	"github.com/rxx/golang-todo/6-database/models"
	"github.com/spf13/cobra"
)

var (
	dueDate     string
	currentTask int
	tasks       models.Tasks
	tx          *pop.Connection
	err         error
)

func init() {
	tx, err = pop.Connect("development")
	if err != nil {
		panic(err)
	}

	err = tx.All(&tasks)
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(editTitleCmd)
	rootCmd.AddCommand(editWhenCmd)
	rootCmd.AddCommand(closeCmd)
	rootCmd.AddCommand(removeCmd)

	addCmd.Flags().StringVar(&dueDate, "due_date", "today", "Task due date")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(tasks) < 1 {
			fmt.Println("No tasks")
		} else {
			fmt.Println(tasks)
		}
	},
}

var addCmd = &cobra.Command{
	Use:   "add title [due_date]",
	Short: `Add new open task`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var title string
		title = strings.Join(args, " ")
		task := models.NewTask(title, dueDate)

		err = tx.Create(&task)
		if err != nil {
			panic(err)
		}

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

		if len(dueDate) > 1 {
			tasks[position-1].SetDueDate(dueDate)
		}

		err = tx.Update(&tasks[position-1])
		if err != nil {
			panic(err)
		}

		if verbose {
			fmt.Println("Changed task", tasks[position-1])
		}
	},
}

var editWhenCmd = &cobra.Command{
	Use:   "edit_due_date task_position dueDate",
	Short: "Edit task DueDate",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		position, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		dueDate = strings.Join(args[1:], " ")
		tasks[position-1].SetDueDate(dueDate)
		err = tx.Update(&tasks[position-1])
		if err != nil {
			panic(err)
		}

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
		err = tx.Update(&tasks[position-1])
		if err != nil {
			panic(err)
		}

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

		err = tx.Destroy(&tasks[position-1])
		if err != nil {
			panic(err)
		}

		if verbose {
			fmt.Printf("Removed task at %d position\n", position)
		}
	},
}
