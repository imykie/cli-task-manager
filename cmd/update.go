package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use: "update",
	Short: "Updates a particular task",
	Long: "Usage: task update {taskId} {newValue}",

	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Update called!")
	},
}

func init(){
	RootCmd.AddCommand(updateCmd)
}