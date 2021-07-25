package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task/db"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a particular task",
	Long:  "Example: task update {taskId} {newValue}",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 2 {
			fmt.Println("update command only takes two argument; task number and new value")
		}

		id, err := strconv.Atoi(args[0])
		Must(err)
		updatedTask := args[1]
		fmt.Println(id, updatedTask)

		tasks, err := db.AllTasks()
		Must(err)

		for i, task := range tasks {
			if id == i + 1 {
				Must(db.DeleteTask(task.Key))
				Must(db.UpdateTask(task.Key, updatedTask))
				break;
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
