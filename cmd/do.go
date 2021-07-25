package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task/db"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Sets command to done",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
				return
			}
			ids = append(ids, id)
		}

		tasks, err := db.AllTasks()
		Must(err)

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed.\nError: %s\n", id, err)
			} else {
				fmt.Printf("Marked task \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
