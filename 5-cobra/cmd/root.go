package cmd

import (
	"fmt"

	"github.com/rxx/golang-todo/5-cobra/task"
	"github.com/spf13/cobra"
)

var (
	err        error
	tasks      task.List
	verbose    bool
	needToSave bool
)

const fileName = "tasks.json"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Application for managing todo lists",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		tasks, err = task.LoadJSON(fileName)
		if err != nil {
			panic(err)
		}

		needToSave = true

		if verbose {
			fmt.Println("Loaded tasks from file: ", fileName)
			fmt.Println(tasks)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if !needToSave {
			if verbose {
				fmt.Println("Nothing was changed. Skip saving tasks")
			}
			return
		}

		err := task.SaveJSON(tasks, fileName)
		if err != nil {
			panic(err)
		}

		if verbose {
			fmt.Println("Saved tasks to file: ", fileName)
			fmt.Println(tasks)
		}

		fmt.Println("Completed successfully")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().
		BoolVarP(&verbose, "verbose", "v", false, "print debug information")
}
