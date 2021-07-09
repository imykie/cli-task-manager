package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a task to your task list.",

	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Add called!")
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}
