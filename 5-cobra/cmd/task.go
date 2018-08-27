package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new open task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")
	},
}

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("close called")
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(closeCmd)
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
