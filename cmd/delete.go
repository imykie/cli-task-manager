package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes task from task list",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete called!")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
