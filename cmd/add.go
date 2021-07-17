package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a task to your task list.",

	Run: func(cmd *cobra.Command, args []string){
		if len(args) == 0 {
			fmt.Printf("Task name is required. \nExample: task add \"Do Something\"\n")
			return
		}
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}
