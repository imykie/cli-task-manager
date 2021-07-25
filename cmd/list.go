package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists your tasks.",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		Must(err)

		if len(tasks) == 0 {
			fmt.Println("You have no task to complete! Add new task or take a break 😊✨✨✨")
			return
		}

		fmt.Println("You have the following tasks to complete")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
