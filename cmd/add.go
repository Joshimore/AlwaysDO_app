/*
Copyright © 2026 Joshimore <https://github.com/Joshimore/AlwaysDO_app>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <your task>",
	Short: "A command to add a task to the list",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		/* Validate priority
		validPriorities := map[string]bool{"high": true, "medium": true, "low": true}
		if !validPriorities[priority] {
			fmt.Printf("Invalid priority '%s'. Use: high, medium, or low\n", priority)
			os.Exit(1)
		}*/

		fmt.Println("Task received: ", args[0])
	},

	Args: cobra.ExactArgs(1),
}


func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
