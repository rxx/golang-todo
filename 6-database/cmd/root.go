package cmd

import (
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

// rootCmd represents the base command dueDate called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Application for managing todo lists",
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
